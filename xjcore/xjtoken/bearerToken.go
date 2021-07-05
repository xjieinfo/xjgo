package xjtoken

import (
	"errors"
	"github.com/xjieinfo/xjgo/xjhttp"
)

//BearerToken格式为:bearer + token,取真正的token
func GetBearerToken(ctx *xjhttp.Context) (string, error) {
	Authorization := ctx.Request.Header.Get("Authorization")
	if len(Authorization) < 8 {
		err := errors.New("Authorization fail.")
		return "", err
	}
	AccessToke := Authorization[7:]
	return AccessToke, nil
}

func GetUserFromToken(token, AccessSecret string) (user string, err error) {
	tokenClaims, err1 := ParseToken(AccessSecret, token)
	if err1 != nil {
		err = err1
		return
	}
	user = tokenClaims.User
	return
}

func GetExtFromToken(token, AccessSecret string) (ext string, err error) {
	tokenClaims, err1 := ParseToken(AccessSecret, token)
	if err1 != nil {
		err = err1
		return
	}
	ext = tokenClaims.Ext
	return
}

func GetAllFromToken(token, AccessSecret string) (user, ext string, err error) {
	tokenClaims, err1 := ParseToken(AccessSecret, token)
	if err1 != nil {
		err = err1
		return
	}
	user = tokenClaims.User
	ext = tokenClaims.Ext
	return
}

func GetUserFromBearerToken(ctx *xjhttp.Context, AccessSecret string) (user string, err error) {
	token, err := GetBearerToken(ctx)
	if err != nil {
		return "0", err
	}
	user, err = GetUserFromToken(token, AccessSecret)
	return
}

func GetAllFromBearerToken(ctx *xjhttp.Context, AccessSecret string) (user, ext string, err error) {
	token, err := GetBearerToken(ctx)
	if err != nil {
		return "0", "", err
	}
	user, ext, err = GetAllFromToken(token, AccessSecret)
	return
}
