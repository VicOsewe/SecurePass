package auth

type AuthService interface {
}

// AuthServiceImpl implements the AuthService interface.
type AuthServiceImpl struct {
}

func NewAuthService() *AuthServiceImpl {
	return &AuthServiceImpl{}
}
