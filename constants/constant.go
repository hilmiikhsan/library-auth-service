package constants

var (
	ErrUsernameAlreadyRegistered     = "username already registered"
	ErrUsernameOrPasswordIsIncorrect = "username or password is incorrect"
	ErrAuthorizationIsEmpty          = "authorization is empty"
	ErrFailedGenerateToken           = "failed to generate token"
	ErrFailedGenerateRefreshToken    = "failed to generate refresh token"
)

const (
	ErrFailedBadRequest = "failed to parse request"
)

const (
	TokenTypeAccess    = "token"
	RefreshTokenAccess = "refresh_token"
)
