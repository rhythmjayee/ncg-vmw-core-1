package auth

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
