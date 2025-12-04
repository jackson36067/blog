package api

type Api struct {
	LoginApi    LoginApi
	RegisterApi RegisterApi
	EmailApi    EmailApi
	UserApi     UserApi
	ArticleApi  ArticleApi
	FavoriteApi FavoriteApi
	CommentApi  CommentApi
	MessageApi  MessageApi
}

var App = Api{}
