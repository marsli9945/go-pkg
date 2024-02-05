package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Auth struct {
	Username string `bson:"username"json:"username"`
	Role     string `bson:"role"json:"role"`
}

type Claims struct {
	Auth Auth
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(auth Auth, jwtExpireTime time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(jwtExpireTime)

	claims := Claims{
		auth,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
