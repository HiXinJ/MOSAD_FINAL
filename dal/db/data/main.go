package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	mydb "github.com/hixinj/MOSAD_FINAL_Group05/dal/db"
	"github.com/hixinj/MOSAD_FINAL_Group05/views"
)

func main() {
	db, err := bolt.Open(mydb.GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Batch(func(tx *bolt.Tx) error {
		fmt.Println("batch writing ...")
		b, _ := tx.CreateBucket([]byte("words_translation"))

		fd, err := os.Open(mydb.GetDBDIR() + "/data/data.txt")
		defer fd.Close()
		if err != nil {
			log.Fatal(err)
		}
		buff := bufio.NewReader(fd)
		var cnt uint64
		cnt = 0
		for {
			line := make([]byte, 20)
			line, _, err := buff.ReadLine()
			if line == nil {
				fmt.Println("nil line")
				continue
			}
			if trans, ok := getfanyi(string(line)); ok && trans.ExampleCH != "" {
				cnt++
				fmt.Println(trans.Query, trans.UKP, trans.ExampleCH)
				data, _ := json.Marshal(trans)
				b.Put(line, data)
			}
			if err == io.EOF || cnt == 400 {
				break
			}

		}
		return nil
	})
}

func getfanyi(word string) (mydb.SimpleTranslation, bool) {
	url := "http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q=" + word

	//生成client 参数为默认
	client := &http.Client{}

	//提交请求
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(reqest)
	resBody := new(mydb.Translation)
	json.NewDecoder(response.Body).Decode(&resBody)
	simpleTrans := mydb.SimpleTranslation{}
	examples := views.GetExample(word)
	if examples != nil {
		simpleTrans.ExampleEN = examples[0]
		simpleTrans.ExampleCH = examples[1]
		simpleTrans.Query = word
		simpleTrans.Explains = resBody.Basic.Explains
		simpleTrans.UKP = resBody.Basic.UKP
		simpleTrans.USP = resBody.Basic.USP
	} else {
		return simpleTrans, false
	}
	return simpleTrans, true
}
