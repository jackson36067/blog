package flags

import (
	"blog/global"
	"blog/models"

	"github.com/sirupsen/logrus"
)

func MigrateDB() {
	err := global.MysqlDB.AutoMigrate(
		&models.User{},
		&models.UserConfig{},
		&models.Article{},
		&models.ArticleCategory{},
		&models.ArticleLike{},
		&models.Favorite{},
		&models.UserArticleCollect{},
		&models.UserTopArticle{},
		&models.Image{},
		&models.UserArticleBrowseHistory{},
		&models.Comment{},
		&models.Log{},
		&models.Banner{},
		&models.UserLogin{},
		&models.GlobalNotification{},
		&models.ArticleTag{},
		&models.UserFollow{},
	)
	if err != nil {
		logrus.Fatalf("数据库迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功")
}
