package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type AdminPath struct {
	Path   string
	Method string
}

type Claims struct {
	jwt.RegisteredClaims
	CustomAttributes map[string]interface{} `json:"custom:attributes"`
}
