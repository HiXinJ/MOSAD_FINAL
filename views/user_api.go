package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
	"github.com/hixinj/MOSAD_FINAL/dal/model"
)

func UserLogin(c *gin.Context) {
	var user model.User
	// user.UserName, _ = c.GetPostForm("user_name")
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": "param error",
		})
	}

	// 验证用户密码
	check := mydb.GetUser(user.UserName)
	if check.UserName != user.UserName || user.UserName == "" {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": "用户不存在",
		})
		return
	}
	if check.Password != user.Password || user.Password == "" {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": "密码错误",
		})
		return
	}

	// 签发token
	tokenString, err := SignToken(user.UserName)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": "sign token error",
		})
	}

	c.JSON(200, gin.H{
		"message":       "success",
		"token":         tokenString,
		"error_message": "",
	})
}

func UserRegister(c *gin.Context) {
	var user model.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	// 用户密码合法性检查
	ok, _ := regexp.MatchString(`^[\w]{3,18}$`, user.UserName)
	if !ok {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "用户名必须为3-18个字母，数字或者下划线",
		})
		return
	}
	ok, _ = regexp.MatchString(`^[\w]{6,18}$`, user.Password)
	if !ok {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "密码必须为3-18个字母，数字或者下划线",
		})
		return
	}

	// if user.UserName == "" || user.Password == "" {
	// 	c.JSON(200, gin.H{
	// 		"message":       "",
	// 		"error_message": "参数错误",
	// 	})
	// 	return
	// }
	// 验证UserName
	check := mydb.GetUser(user.UserName)
	if check.UserName != "" {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": "用户已存在",
		})
		return
	}
	// 存储User
	err := mydb.PutUsers([]model.User{user})
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "",
			"error_message": err.Error(),
		})
		return
	}

	tokenString, err := SignToken(user.UserName)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "error",
			"error_message": "signing token error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":       "success",
		"token":         tokenString,
		"error_message": "",
	})
}

func DaKa(c *gin.Context) {
	userName := ValidateToken(c.Writer, c.Request)
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "daka authentication fail",
		})
		return
	}
	now := time.Now()
	today := model.Date{int64(now.Year()), int64(now.Month()), int64(now.Day())}
	user := mydb.GetUser(userName)
	if i := len(user.DaKa); i == 0 || !(user.DaKa[i-1].Year == today.Year && user.DaKa[i-1].Month == today.Month && user.DaKa[i-1].Day == today.Day) {
		user.DaKa = append(user.DaKa, today)
	}
	mydb.PutUsers([]model.User{user})
	ndays := len(user.DaKa)
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	res["ndays"] = ndays
	res["date"] = user.DaKa
	c.JSON(200, res)
}

func GetHead(c *gin.Context) {
	userName := ValidateToken(c.Writer, c.Request)
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "authentication fail",
		})
		return
	}

	user := mydb.GetUser(userName)
	c.Data(200, "image/png", user.Head)
	// f, _ := os.Open("1.png")
	// img, _ := ioutil.ReadAll(f)
	// c.Data(200, "image/png", img)
}

func PostHead(c *gin.Context) {
	userName := ValidateToken(c.Writer, c.Request)

	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "authentication fail",
		})
		return
	}

	user := mydb.GetUser(userName)
	img, err := ioutil.ReadAll(c.Request.Body)
	ioutil.WriteFile("1.png", img, 0644)

	if err != nil {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": err.Error(),
		})
		return
	}

	user.Head = make([]byte, len(img))
	copy(user.Head, img)
	mydb.PutUsers([]model.User{user})
	t := mydb.GetUser(user.UserName)
	fmt.Println(t.UserName)
	c.JSON(200, gin.H{
		"message":       "success",
		"error_message": "",
	})
}

//******************** token ********************

// SecretKey ..
const SecretKey = "123qwe"

func ValidateToken(w http.ResponseWriter, r *http.Request) string {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		return ""
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "token is invalid")
		return ""
	}

	if v, ok := token.Claims.(jwt.MapClaims); ok {
		name, _ := v["name"].(string)
		return name
	}

	return ""
}

func SignToken(userName string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(10)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["name"] = userName
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}
