package api

type Api struct {
	LoginApi      LoginApi
	RegisterApi   RegisterApi
	EmailApi      EmailApi
	UserApi       UserApi
	ArticleApi    ArticleApi
	ArticleTagApi ArticleTagApi
}

var App = Api{}
