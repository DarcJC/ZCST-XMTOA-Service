package utils

import (
	"github.com/mix-go/dotenv"
	"github.com/pascaldekloe/jwt"
	"time"
)

// NewJWT godoc
// @Description 创建一个新的JWT凭据
func NewJWT(subject string, context *AnyStruct) string {
	var claims jwt.Claims
	// 设置JWT的主题
	claims.Subject = subject
	// 设置签发者
	claims.Issuer = "OA-system"
	// 设置签发时间为 当前时间
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	// 设置JWT内容
	claims.Set = *context

	// 获取JWT签名密钥
	secret := dotenv.Getenv("JWT_SECRET").String("")
	if secret == "" {
		panic("JWT_SECRET must not be empty, stopped.")
	}

	// 签名并生成base64
	token, err := claims.HMACSign("HS384", []byte(secret))
	if err != nil {
		return ""
	}
	return string(token)
}
