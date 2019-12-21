package views

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
)

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
	res["fanyi"] = fanyi
	c.JSON(200, res)
	// fmt.Println(size)
}
