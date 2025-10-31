package api

type Api struct {
	LoginApi    LoginApi
	RegisterApi RegisterApi
	EmailApi    EmailApi
	UserApi     UserApi
	ArticleApi  ArticleApi
}

var App = Api{}
