package handlers

import (
	"github.com/VicOsewe/secure-pass/pkg/usecase/auth"
)

type Presentation interface {
}

type RestFulAPIs struct {
	authService auth.AuthService
}

func NewRestFulAPIs() *RestFulAPIs {
	return &RestFulAPIs{}
}
