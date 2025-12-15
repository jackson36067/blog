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
	SessionApi  SessionApi
}

var App = Api{}
