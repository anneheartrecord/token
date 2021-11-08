package common

import (
	"github.com/dgrijalva/jwt-go"
	"golangstudy/jike/gingormpractise/model"
	"time"
)
var jwtKey=[] byte("a_secret_crect")
type Claims struct {  //定义claim
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string,error) {
	expirationTime:=time.Now().Add(7*24*time.Hour) //存活时间
	claims:=&Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:  expirationTime.Unix(), //发放时间
			IssuedAt: time.Now().Unix(),
			Issuer: "111",
			Subject: "user token",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtKey)
	if err!=nil {
		return "",err
	}
	return tokenString,nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims,error) {
	claims:=&Claims{}
	token,err:=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{},err error) {
		return jwtKey,nil
	})
	return token,claims,err

}