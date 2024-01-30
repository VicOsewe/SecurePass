package registration

type AuthService interface {
}

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}

}
