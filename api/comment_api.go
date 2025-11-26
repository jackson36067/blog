package api

import (
	"blog/consts"
	"blog/dto/request"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"blog/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentApi struct {
}

// PublishCommentView 发布评论
func (CommentApi) PublishCommentView(c *gin.Context) {
	// 获取评论用户id
	userIdAny, ok := c.Get(consts.UserId)
	if !ok {
		res.Fail(c, 401, "未登录")
		return
	}
	userId := userIdAny.(uint)

	// 绑定参数
	var params request.PublishCommentRequestParams
	if err := c.ShouldBindJSON(&params); err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
		return
	}

	// 开启事务
	db := global.MysqlDB
	err := db.Transaction(func(tx *gorm.DB) error {

		// 插入评论
		result := tx.Create(&models.Comment{
			Content:       params.Content,
			UserID:        userId,
			ArticleID:     params.CommentArticleId,
			ParentID:      params.ParentCommentId,
			RootParentID:  params.RootCommentId,
			LikeCount:     0,
			ReplyToUserID: params.ReplyToUserId,
		})

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("插入评论失败：RowsAffected = 0")
		}

		// 更新文章评论数
		if err := tx.Model(&models.Article{}).
			Where("id = ?", params.CommentArticleId).
			Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		res.Fail(c, 500, consts.AffairCommitError)
		return
	}
}

// LikeCommentView 点赞评论
func (CommentApi) LikeCommentView(c *gin.Context) {
	// 获取评论用户id
	userIdAny, ok := c.Get(consts.UserId)
	if !ok {
		res.Fail(c, 401, "未登录")
		return
	}
	userId := userIdAny.(uint)
	commentIdStr := c.Param("id")
	commentId, _ := utils.StringToUint(commentIdStr)
	var params request.LikeCommentRequestParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
		return
	}
	db := global.MysqlDB
	if !params.IsLike {
		err = db.Transaction(func(tx *gorm.DB) error {
			result := tx.Create(&models.CommentLike{
				CommentID: commentId,
				UserID:    userId,
			})

			if result.Error != nil {
				return result.Error
			}

			if result.RowsAffected == 0 {
				return errors.New("点赞评论失败：RowsAffected = 0")
			}

			if err := tx.Model(&models.Comment{}).
				Where("id = ?", commentId).
				Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			res.Fail(c, 500, consts.AffairCommitError)
			return
		}
	} else {
		err = db.Transaction(func(tx *gorm.DB) error {
			result := tx.Where("comment_id = ? AND user_id = ?", commentId, userId).Delete(&models.CommentLike{})

			if result.Error != nil {
				return result.Error
			}

			if result.RowsAffected == 0 {
				return errors.New("取消点赞评论失败：RowsAffected = 0")
			}

			if err := tx.Model(&models.Comment{}).
				Where("id = ?", commentId).
				Update("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			res.Fail(c, 500, consts.AffairCommitError)
			return
		}
	}
}

// DeleteCommentView 删除评论
func (CommentApi) DeleteCommentView(c *gin.Context) {
	commentIdStr := c.Param("id")
	commentId, _ := utils.StringToUint(commentIdStr)
	var params request.DeleteCommentRequestParams
	if err := c.ShouldBindJSON(&params); err != nil {
		res.Fail(c, 500, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 为根评论(需要删除该评论以及根评论为该评论的评论)
	var err error
	if params.IsRoot {
		err = db.Transaction(func(tx *gorm.DB) error {
			result := tx.Where("id = ?", commentId).Delete(&models.Comment{})
			if result.Error != nil {
				return result.Error
			}
			result = tx.Where("root_parent_id = ?", commentId).Delete(&models.Comment{})
			if result.Error != nil {
				return result.Error
			}
			return nil
		})
	} else {
		// 不为根评论(删除该评论以及父评论为该评论的评论)
		err = db.Transaction(func(tx *gorm.DB) error {
			// 获取该评论所有子评论(递归获取)
			ids, err := service.GetAllChildIDs(tx, commentId)
			if err != nil {
				return err
			}
			// 一次性删除所有子树节点
			if err := tx.Where("id IN ?", ids).Delete(&models.Comment{}).Error; err != nil {
				return err
			}
			return nil
		})
	}
	if err != nil {
		res.Fail(c, 500, consts.AffairCommitError)
	}
}
