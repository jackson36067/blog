package api

type Api struct {
	LoginApi    LoginApi
	RegisterApi RegisterApi
	EmailApi    EmailApi
}

var App = Api{}
