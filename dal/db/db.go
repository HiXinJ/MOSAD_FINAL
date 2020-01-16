package mydb

import (

	// "code.byted.org/gopkg/logs"

	"encoding/json"
	"log"
	"math/rand"
	"runtime"
	"time"

	"os"
	"path"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hixinj/MOSAD_FINAL/dal/model"
)

//************************************************************************

func GetDBDIR() string {
	ostype := runtime.GOOS
	pt, _ := os.Getwd()
	if ostype == "windows" {
		return pt + "\\dal\\db"
	}

	return path.Join(os.Getenv("GOPATH"), "src", "github.com", "hixinj", "MOSAD_FINAL", "dal", "db")
}
func GetDBPATH() string {
	ostype := runtime.GOOS
	if ostype == "windows" {
		return GetDBDIR() + "\\data\\final.db"
	}
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
				username := users[i].UserName
				if users[i].LearnedWords == nil {
					users[i].LearnedWords = make(map[string]int64)
				}
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
		UserName: "",
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

func FilterWords(size int64, filter func(word string) bool) []string {
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
			// cursor := b.Cursor()
			var hasList [5000]int64
			rand.Seed(time.Now().Unix())

			for /*k, v := cursor.First(); k != nil; k, v = cursor.Next()*/ {
				if int64(cnt) == size {
					break
				}
				i := rand.Intn(401)
				word := string(b.Get([]byte(string(i))))
				if filter(word) && hasList[i] == 0 {
					hasList[i] = 1
					cnt++
					wordList = append(wordList, word)
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
