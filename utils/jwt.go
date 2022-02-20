package utils

import (
	"errors"
	"github.com/mix-go/dotenv"
	"github.com/pascaldekloe/jwt"
	"time"
)

func getJWTSecret() string {
	// 获取JWT签名密钥
	secret := dotenv.Getenv("JWT_SECRET").String("")
	if secret == "" {
		panic("JWT_SECRET must not be empty, stopped.")
	}
	return secret
}

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
	// 设置过期时间为 1小时
	claims.Expires = jwt.NewNumericTime(time.Now().Add(time.Hour * 1).Round(time.Second))
	// 设置JWT内容
	claims.Set = *context

	// 获取JWT签名密钥
	secret := getJWTSecret()

	// 签名并生成base64
	token, err := claims.HMACSign("HS384", []byte(secret))
	if err != nil {
		return ""
	}
	return string(token)
}

// CheckJWT godoc
// @Description 检查JWT，如果正确会返回一个claim
func CheckJWT(token string) (*jwt.Claims, error) {
	// 获取JWT签名密钥
	secret := getJWTSecret()

	// 检查JWT签名
	claims, err := jwt.HMACCheck([]byte(token), []byte(secret))
	if err != nil {
		// 有内鬼，中止交易
		return nil, err
	}

	// 检查JWT是否到期
	if !claims.Valid(time.Now()) {
		// 保质期过了，扔掉
		return nil, errors.New("the token has expired")
	}

	return claims, nil
}
