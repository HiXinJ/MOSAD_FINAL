package views

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
	"github.com/hixinj/MOSAD_FINAL/dal/model"
)

func getfanyi(wordsList []string) []model.SimpleTranslation {
	baseurl := "http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q="

	size := len(wordsList)
	cnt := 0
	fanyi := make([]model.SimpleTranslation, 0)
	for i := 0; i < len(wordsList); i++ {
		item := wordsList[i]
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
		// resBody := new(interface{})
		resBody := new(model.Translation)
		simpleTrans := model.SimpleTranslation{}
		json.NewDecoder(response.Body).Decode(&resBody)
		if err != nil {
			log.Fatal(err)
		}
		examples := GetExample(item)
		if examples == nil {
			wordsList = append(wordsList, mydb.GetWords(1)[0])
			continue
		}
		simpleTrans.ExampleEN = examples[0]
		simpleTrans.ExampleCH = examples[1]
		simpleTrans.Query = item
		simpleTrans.Explains = resBody.Basic.Explains
		simpleTrans.UKP = resBody.Basic.UKP
		simpleTrans.USP = resBody.Basic.USP
		fanyi = append(fanyi, simpleTrans)
		cnt++
		if cnt == size {
			break
		}
	}

	return fanyi
}

// 获取今日新词
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
			return false
		}
		return true
	})
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	// res["words"] = wordsList
	res["fanyi"] = getfanyi(wordsList)
	c.JSON(200, res)
}

// 随机获取单词
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
		return
	}
	userName := ValidateToken(c.Writer, c.Request)
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "authetication failed",
		})
		return
	}
	user := mydb.GetUser(userName)
	user.LearnedWords[word] = 1
	// 更新
	mydb.PutUsers([]model.User{user})

	c.JSON(200, gin.H{
		"message":       "success",
		"error_message": "",
	})
}

func GetReviews(c *gin.Context) {
	userName := ValidateToken(c.Writer, c.Request)
	if userName == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "authetication failed",
		})
		return
	}

	sizeStr := c.DefaultQuery("size", "5")
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	user := mydb.GetUser(userName)
	wordsList := make([]string, len(user.LearnedWords))
	i := 0
	for k := range user.LearnedWords {
		wordsList[i] = k
		i++
	}
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	if int(size) > len(user.LearnedWords) {
		res["translation"] = getfanyi(wordsList)
	} else { // 从用户已学单词中随机获取size个单词
		visited := make([]int, 1000)
		reviewList := make([]string, 0)
		rand.Seed(time.Now().Unix())
		for i = 0; i < (int)(size); {
			t := rand.Intn(len(user.LearnedWords))
			if visited[t] == 0 {
				reviewList = append(reviewList, wordsList[t])
				i++
			}
			visited[t] = 1
		}
		res["translation"] = getfanyi(reviewList)
	}

	c.JSON(200, res)
}

func GetTranslation(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		c.JSON(200, gin.H{
			"message":       "failed",
			"error_message": "empty param",
		})
		return
	}

	baseurl := "http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q="

	fanyi := make([]model.SimpleTranslation, 0)

	item := word
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
	// resBody := new(interface{})
	resBody := new(model.Translation)
	simpleTrans := model.SimpleTranslation{}
	json.NewDecoder(response.Body).Decode(&resBody)
	if err != nil {
		log.Fatal(err)
	}
	if resBody.Basic.Explains == nil {
		c.JSON(200, gin.H{
			"message":        "failed",
			"error_message:": "单词未找到",
		})
		return
	}
	examples := GetExample(item)
	if examples != nil {
		simpleTrans.ExampleEN = examples[0]
		simpleTrans.ExampleCH = examples[1]
	}
	simpleTrans.Query = item
	simpleTrans.Explains = resBody.Basic.Explains
	simpleTrans.UKP = resBody.Basic.UKP
	simpleTrans.USP = resBody.Basic.USP
	fanyi = append(fanyi, simpleTrans)

	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	res["translation"] = fanyi
	c.JSON(200, res)
}
