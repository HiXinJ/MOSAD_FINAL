package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/hixinj/iOSServer/dal/model"
)

func UserLogin(c *gin.Context) {

}

func UserRegister(c *gin.Context) {
	var user model.User
	user.UserName = c.Param("user_name")
	user.Password = c.Param("password")

	// TODO: 验证UserName

	// TODO: 存储User

	tokenString, err := SignToken(user.UserName)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "error",
			"error_message": "signing token error",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

const SecretKey = "123qwe"

func ValidateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, bool) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		return token, false
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "token is invalid")
		return token, false
	}

	return token, true
}

func SignToken(userName string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["name"] = userName
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}
