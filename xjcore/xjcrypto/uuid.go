package xjcrypto

import "github.com/google/uuid"

func Uuid() string {
	return uuid.NewString()
}
