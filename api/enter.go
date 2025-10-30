package api

type Api struct {
	LoginApi    LoginApi
	RegisterApi RegisterApi
	EmailApi    EmailApi
	UserApi     UserApi
}

var App = Api{}
