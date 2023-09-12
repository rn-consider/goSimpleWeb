package my_jwt

import "github.com/golang-jwt/jwt/v5"

// CustomClaims 自定义jwt的声明字段信息+标准字段
type CustomClaims struct {
	UserId    int64  `json:"user_id"`
	Name      string `json:"user_name"`
	Phone     string `json:"phone"`
	ExpiresAt int64  `json:"expires_at"`
	jwt.Claims
}
