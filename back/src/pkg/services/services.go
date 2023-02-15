package services

var registerService RegisterService
var loginService LoginService
var jwtService JWTService

func InitializeServices() {
	registerService = RegisterService(&RegisterServiceStruct{})
	loginService = LoginService(&LoginServiceStruct{})
	jwtService = JWTService(&JWTServiceStruct{})
}

func GetRegisterService() RegisterService {
	return registerService
}

func GetLoginService() LoginService {
	return loginService
}

func GetJWTService() JWTService {
	return jwtService
}
