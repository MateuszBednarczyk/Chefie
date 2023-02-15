package services

var iRegisterService IRegisterService
var iLoginService ILoginService
var iJwtService IJWTService

func InitializeServices() {
	iRegisterService = NewRegisterService()
	iLoginService = NewLoginService()
	iJwtService = NewJwtService()
}

func GetRegisterService() IRegisterService {
	return iRegisterService
}

func GetLoginService() ILoginService {
	return iLoginService
}

func GetJWTService() IJWTService {
	return iJwtService
}
