package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// 私钥
const secretkey = "lijialin"

// 生成token
func GenerateToken(uid int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 60 * time.Minute)
	myClaims := MyClaims{
		ID:       uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ljl",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := tokenClaims.SignedString([]byte(secretkey))
	return token, err
}

// 解析token
func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
