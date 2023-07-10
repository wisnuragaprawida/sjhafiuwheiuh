package crashy

const (
	// ErrCodeUnexpected code for generic error for unrecognized cause
	ErrCodeUnexpected = "ERR_UNEXPECTED"

	// ErrCodeNetBuild code  for resource connection build issue
	ErrCodeNetBuild = "ERR_NET_BUILD"

	// ErrCodeNetConnect code for resource connection issue
	ErrCodeNetConnect = "ERR_NET_CONNECT"

	//ErrCodeValidation code for validation error
	ErrCodeValidation = "ERR_VALIDATION"

	//ErrCodeFormatting code for formatting error
	ErrCodeFormatting = "ERR_FORMATTING"

	//ErrCodeDataRead code for data read error
	ErrCodeDataRead = "ERR_DATA_READ"

	//ErrCodeDataWrite code for data write error

	ErrCodeDataWrite = "ERR_DATA_WRITE"

	//ErrCodeNoResult code for no result error
	ErrCodeNoResult = "ERR_NO_RESULT"

	//ErrCodeUnauthorized code for unauthorized error
	ErrCodeUnauthorized = "ERR_UNAUTHORIZED"
	ErrCodeExpired      = "ERR_EXPIRED"
	ErrCodeForbidden    = "ERR_FORBIDDEN"

	//ErrCodeToManyRequests code for too many requests error
	ErrCodeToManyRequests = "ERR_TOO_MANY_REQUESTS"
	//ErrCodeDataIncomplete code for incomplete data error
	ErrCodeDataIncomplete = "ERR_DATA_INCOMPLETE"

	//ErrCodeEncryptData code for data encryption error
	ErrCodeEncryptData = "ERR_ENCRYPT_DATA"

	//Err code when failed get activity data
	ErrCodeGetActivity = "ERR_GET_ACTIVITY"

	//Err code when failed get template data
	ErrCodeGetTemplate = "ERR_GET_TEMPLATE"

	//ErrCodeTimeResend code for time resend error
	ErrCodeTimeResend = "ERR_TIME_RESEND"

	ErrParsingData = "ERR_PARSING_DATA"
)

var mapper = map[ErrCode]string{
	ErrCodeUnexpected:     "unexpected error",
	ErrCodeNetBuild:       "error building connection",
	ErrCodeNetConnect:     "error connecting to resource",
	ErrCodeValidation:     "validation error",
	ErrCodeFormatting:     "formatting error",
	ErrCodeDataRead:       "error reading data",
	ErrCodeDataWrite:      "error writing data",
	ErrCodeNoResult:       "no result found",
	ErrCodeUnauthorized:   "unauthorized access",
	ErrCodeExpired:        "expired access",
	ErrCodeForbidden:      "forbidden access",
	ErrCodeToManyRequests: "too many requests",
	ErrCodeDataIncomplete: "incomplete data",
	ErrCodeEncryptData:    "error encrypting data",
	ErrCodeGetActivity:    "error getting activity data",
	ErrCodeGetTemplate:    "error getting template data",
	ErrCodeTimeResend:     "error time resend",
	ErrParsingData:        "error parsing data",
}

func Message(code ErrCode) string {
	if s, ok := mapper[code]; ok {
		return s
	}
	return mapper[ErrCodeUnexpected]
}

func Messages() map[ErrCode]string {
	return mapper
}
