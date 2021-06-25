package xjtoken

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

//var jwtSecret=[]byte(conf.AppSetting.JwtSecret)

type Claims struct {
	Userid int    `json:"userid"`
	Ext    string `json:"ext"`
	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(jwtSecret string, jwtExpire, userid int, ext string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtExpire) * time.Minute)

	claims := Claims{
		userid,
		ext,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "xjgo",
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

// 验证token的函数
func ParseToken(jwtSecret string, token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//
	return nil, err
}
