package services

var iRegisterService IRegisterService
var iLoginService ILoginService
var iJwtService IJWTService

func InitializeServices() {
	iRegisterService = NewRegisterService()
	iLoginService = NewLoginService()
	iJwtService = NewJwtService()
}

func RegisterService() IRegisterService {
	return iRegisterService
}

func LoginService() ILoginService {
	return iLoginService
}

func JwtService() IJWTService {
	return iJwtService
}
