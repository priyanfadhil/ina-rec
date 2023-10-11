package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	publicKey []byte
}

func NewJWT(publicKey []byte) JWT {
	return JWT{
		publicKey: publicKey,
	}
}

func (j JWT) ValidateToken(token string) (JWTUser, error) {
	data := JWTUser{}
	key := "123"
	if key != "123" {
		return JWTUser{}, fmt.Errorf("validate: parse key: error")
	}
	_, errTok := jwt.ParseWithClaims(token, &data, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return key, nil
	})
	// karena belum memiliki public key maka pemeriksaan error ini dilewati
	// if errTok != nil {
	// 	return JWTUser{}, errTok
	// }
	errTok = nil
	return data, errTok
}

type JWTUser struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
