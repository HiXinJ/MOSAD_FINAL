package views

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
	"github.com/hixinj/MOSAD_FINAL/dal/model"
)

func getfanyi(wordsList []string) []interface{} {
	baseurl := "http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q="

	fanyi := make([]interface{}, 0)
	for _, item := range wordsList {
		url := baseurl + item
		//生成client 参数为默认
		client := &http.Client{}

		//提交请求
		reqest, err := http.NewRequest("GET", url, nil)

		if err != nil {
			log.Fatal(err)
		}

		//处理返回结果
		response, _ := client.Do(reqest)
		// resBody, err := ioutil.ReadAll(response.Body)
		resBody := new(interface{})
		json.NewDecoder(response.Body).Decode(&resBody)
		if err != nil {
			log.Fatal(err)
		}
		fanyi = append(fanyi, resBody)
	}
	return fanyi
}

func GetNewWords(c *gin.Context) {
	userName := c.Query("user_name")
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "user_name is empty.",
		})
		return
	}
	sizeStr := c.DefaultQuery("size", "5")
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": err.Error(),
		})
		return
	}

	user := mydb.GetUser(userName)
	if user.UserName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "user doesn't exist",
		})
		return
	}

	wordsList := mydb.FilterWords(size, func(word string) bool {
		if _, ok := user.LearnedWords[word]; ok {
			return true
		}
		return false
	})
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	res["words"] = wordsList
	res["fanyi"] = getfanyi(wordsList)
	c.JSON(200, res)
}

func GetWords(c *gin.Context) {
	sizeStr := c.DefaultQuery("size", "5")
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"message":       "error",
			"error_message": err.Error(),
		})
	}
	wordsList := mydb.GetWords(size)
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	res["words"] = wordsList

	res["fanyi"] = getfanyi(wordsList)
	c.JSON(200, res)
	// fmt.Println(size)
}

func AddLearnedWrod(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "parameter is empty.",
		})
	}
	userName := ValidateToken(c.Writer, c.Request)
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "authetication failed",
		})
	}
	user := mydb.GetUser(userName)
	user.LearnedWords[word] = 1
	mydb.PutUsers([]model.User{user})
}

func getExamples(word string) []string {

	return nil
}
