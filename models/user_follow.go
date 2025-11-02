package models

type UserFollow struct {
	Model
	FollowerID uint `gorm:"not null;uniqueIndex:idx_follower_followed" json:"follower_id"` // 粉丝ID
	FollowedID uint `gorm:"not null;uniqueIndex:idx_follower_followed" json:"followed_id"` // 被关注者ID
}
