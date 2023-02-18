package services

import (
	"back/src/pkg/services/security"
	"back/src/pkg/services/user"
)

var iRegisterService user.IRegisterService
var iLoginService user.ILoginService
var iJwtService security.IJWTService

func InitializeServices() {
	iRegisterService = user.NewRegisterService()
	iLoginService = user.NewLoginService()
	iJwtService = security.NewJwtService()
}

func RegisterService() user.IRegisterService {
	return iRegisterService
}

func LoginService() user.ILoginService {
	return iLoginService
}

func JwtService() security.IJWTService {
	return iJwtService
}
