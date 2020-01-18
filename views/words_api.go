package views

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
)

func getfanyi(wordsList []string) []mydb.SimpleTranslation {
	baseurl := "http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q="

	size := len(wordsList)
	cnt := 0
	fanyi := make([]mydb.SimpleTranslation, 0)
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
		resBody := new(mydb.Translation)
		simpleTrans := mydb.SimpleTranslation{}
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

	wordsList, _ := user.UpdateNewWords(size)
	res := gin.H{
		"message":       "success",
		"error_message": "",
	}
	// res["words"] = wordsList
	translation := getfanyi(wordsList)
	res["fanyi"] = translation
	user.NewWords = make(map[string]int64, len(translation))
	for _, t := range translation {
		user.NewWords[t.Query] = 1
	}
	mydb.PutUsers([]mydb.User{user})
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
	// user.LearnedWords[word] = 1
	user.AddLearnedWord(word)
	// 更新
	mydb.PutUsers([]mydb.User{user})

	c.JSON(200, gin.H{
		"message":       "success",
		"error_message": "",
	})
}

// GetReviews
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
	reviewList := user.GetReviews(size)
	c.JSON(200, gin.H{
		"message":       "success",
		"error_message": "",
		"translation":   getfanyi(reviewList),
	})
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

	fanyi := make([]mydb.SimpleTranslation, 0)

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
	resBody := new(mydb.Translation)
	simpleTrans := mydb.SimpleTranslation{}
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
