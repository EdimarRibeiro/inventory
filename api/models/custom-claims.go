package models

import "github.com/dgrijalva/jwt-go"

var SecretKey = []byte("!QAZxsw2#EDCvfr1pltobtçtyoPIKD8695rmk4j40g98v307kf8&*")

type CustomClaims struct {
	Username   string `json:"username"`
	ExternalId uint64 `json:"externalId"`
	jwt.StandardClaims
}
