package mydb

import (

	// "code.byted.org/gopkg/logs"

	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"time"

	"os"
	"path"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hixinj/MOSAD_FINAL/dal/model"
	"github.com/jinzhu/gorm"
)

func GetArticles(offset, pageSize int) ([]*model.Article, error) {
	// 连接数据库，该数据库为讲师的阿里云数据库，仅限于教学、实验使用
	// 此处也没有使用连接池，可自己再优化一下
	db, err := gorm.Open("mysql", "visitor:Visitor123@tcp(rm-wz9evf905zc36q4e7fo.mysql.rds.aliyuncs.com:3306)/news")
	if err != nil {
		// to process error
		// logs.Error("connect to database failed, err=%v", err)
		return nil, errors.New("connect to database failed.")
	}
	defer db.Close()

	stream := make([]*model.Stream, 0)
	// 使用gorm api访问数据库，获取文章信息流
	db.Select("article_id, rank_1_score").Order("rank_1_score asc").Offset(offset).Limit(pageSize).Find(&stream)

	// 可能获取到的数据不足10条
	if len(stream) < pageSize {
		pageSize = len(stream)
	}

	// 组合文章id列表，准备获取文章信息
	articleIdList := make([]int64, pageSize)
	for _, item := range stream {
		articleIdList = append(articleIdList, int64(item.ArticleId))
	}

	m := make(map[int64]*model.Article)
	articleList := make([]*model.Article, 0)
	// 获取文章信息
	db.Where("article_id in (?)", articleIdList).Find(&articleList)
	for _, item := range articleList {
		// 为了后面将文章信息填充进信息流
		m[item.ArticleId] = item
	}
	return articleList, nil
}

func GetClick(articleID int64, num int64) (model.Favour, error) {

	return model.Favour{int64(1), 312093, []int64{81, 123, 9541}}, nil
}

//************************************************************************

func GetDBDIR() string {
	// ostype := runtime.GOOS
	// if ostype == "windows" {
	// 	pt, _ := os.Getwd()
	// 	return pt + "\\dal\\db\\Blog.db"
	// }
	return path.Join(os.Getenv("GOPATH"), "src", "github.com", "hixinj", "MOSAD_FINAL", "dal", "db")
}
func GetDBPATH() string {
	return path.Join(GetDBDIR(), "data", "final.db")
}

func PutUsers(users []model.User) error {
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b != nil {
			for i := 0; i < len(users); i++ {
				username := users[i].Username
				data, _ := json.Marshal(users[i])
				b.Put([]byte(username), data)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func GetUser(username string) model.User {
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := model.User{
		Username: "",
		Password: "",
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b != nil {
			data := b.Get([]byte(username))
			if data != nil {
				err := json.Unmarshal(data, &user)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func GetWords(size int64) []string {
	wordList := make([]string, 0, size)
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("words1"))
		if b != nil {
			cnt := 0
			hasList := make([]int64, 5000)
			rand.Seed(time.Now().Unix())

			for {
				if int64(cnt) == size {
					break
				}
				i := rand.Intn(401)
				if hasList[i] == 0 || true {
					hasList[i] = 1
					cnt++
					// key := make([]byte, 8)
					// binary.LittleEndian.PutUint64(key, uint64(i))
					word := string(b.Get([]byte(string(i))))

					// word2 := b.Get([]byte{1, 0, 0, 0, 0, 0, 0, 0})
					wordList = append(wordList, word)
					// fmt.Print(string(word2))
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return wordList
}
