package auth

import (
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt/v5"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr RegisterRequest) Bind(r *http.Request) error {
	if err := rr.Validate(); err != nil {
		return err
	}
	return nil
}
func (rr RegisterRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Email, validation.Required),
		validation.Field(&rr.Password, validation.Required),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr LoginRequest) Bind(r *http.Request) error {
	if err := lr.Validate(); err != nil {
		return err
	}
	return nil
}
func (lr LoginRequest) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required),
		validation.Field(&lr.Password, validation.Required),
	)
}

type Claims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
