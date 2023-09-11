package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your-secret-key") // 替换成您自己的密钥

// GenerateJWT 生成JWT令牌
func GenerateJWT(username string, permission int) (string, error) {
	// 创建一个新的JWT令牌
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置令牌的声明（Payload）
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["permission"] = permission
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置令牌过期时间，此处为一天

	// 使用密钥进行签名
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析JWT令牌
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// 解析JWT令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌签名
	if !token.Valid {
		return nil, errors.New("无效的JWT令牌")
	}

	// 获取声明（Payload）
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("无效的JWT令牌声明")
	}

	return claims, nil
}
