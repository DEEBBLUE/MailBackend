package jwtAuth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccess(payload string,cred string) (string,error){
	claims := jwt.MapClaims{
		"email": payload,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(cred))

}

func ValidateToken(token string,cred string) (*jwt.Token,error){
	return jwt.Parse(token,func(t *jwt.Token) (interface{},error) {
		return []byte(cred),nil
	}) 
}

func CreateRefresh(data string,cred string) (string,error) {
	claims := jwt.MapClaims{
		"hash_password": data,
		"exp": time.Now().Add(time.Hour * 360).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(cred))
}
