package utils

import (
	"blog/consts"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义密钥（建议放在配置文件中）
var jwtSecret = []byte("jackson")

// Claims 结构体，自定义你的 Token 载荷
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间，例如 2 小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(consts.ExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 在该时间之前的token无效
			Issuer:    consts.Issuer,                  // 签发者
			Subject:   consts.Subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析并验证 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 判断 token 是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New(consts.InValidToken)
}

// IsTokenExpired 检查 Token 是否过期
func IsTokenExpired(tokenString string) bool {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return true
	}
	return claims.ExpiresAt.Time.Before(time.Now())
}
