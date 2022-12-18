package tools

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Id       int64
	Username string
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userId int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 2)
	claims := Claims{
		Id:       userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cloudRestaurant",
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString("golang")
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
