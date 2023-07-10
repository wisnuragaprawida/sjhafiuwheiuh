package auth

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/wisnuragaprawida/project/generated/users"
	"github.com/wisnuragaprawida/project/internal/api/response"
	"github.com/wisnuragaprawida/project/pkg/crashy"
	"github.com/wisnuragaprawida/project/pkg/one-go/utils"
)

type AuthHandler struct {
	db       *sqlx.DB
	userRepo *users.Queries
}

func NewAuthHandler(db *sqlx.DB) *AuthHandler {
	return &AuthHandler{
		db:       db,
		userRepo: users.New(db),
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	//register user with bcrypt as hash pasword

	var (
		req RegisterRequest
	)

	if err := render.Bind(r, &req); err != nil {

		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode(err.Error())), http.StatusBadRequest)
		// response.Nay(w, r, crashy.Wrapc(err, crashy.ErrParsingData), http.StatusBadRequest)
		return
	}

	//check if email already exist
	_, err := ah.userRepo.FindUserByEmail(r.Context(), req.Email)
	if err == nil {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode("Email Alredy Exist!")), http.StatusBadRequest)
		return
	}

	//hash pasword
	PasswordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode("Hash Password Failed")), http.StatusBadRequest)
		return
	}

	//register user use bycrypt to encrypt password
	_, err = ah.userRepo.RegisterUser(r.Context(), users.RegisterUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: PasswordHash,
	})

	if err != nil {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode("Register Failed")), http.StatusBadRequest)
		return
	}

	response.Yay(w, r, "Register Success", http.StatusOK)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//login user with bcrypt as hash pasword

	var (
		req LoginRequest
	)

	if err := render.Bind(r, &req); err != nil {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode(err.Error())), http.StatusBadRequest)
		return
	}

	//check if email already exist
	user, err := ah.userRepo.FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode("Invalid Credential")), http.StatusBadRequest)
		return
	}

	//check if password match
	ok := utils.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		response.Nay(w, r, crashy.Wrapc(err, crashy.ErrCode("Invalid Credential")), http.StatusBadRequest)
		return
	}
	//generate token
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,

		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	var jwtKey = []byte("my_secret_key")

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself

	response.Yay(w, r, TokenResponse{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}, http.StatusOK)
}
