package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/boltdb/bolt"
	mydb "github.com/hixinj/MOSAD_FINAL_Group05/dal/db"
)

func main() {
	db, err := bolt.Open(mydb.GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("words1"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("words1"))

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("create bucket words1")
		}

		b = tx.Bucket([]byte("words2"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("words2"))

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("create bucket words2")
		}

		b = tx.Bucket([]byte("user"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("user"))
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("create bucket user")
		return nil
	})

	err = db.Batch(func(tx *bolt.Tx) error {
		fmt.Println("batch writing ...")
		b := tx.Bucket([]byte("words1"))
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
			cnt++

			// key := make([]byte, 8)
			key := []byte(string(cnt))
			// binary.LittleEndian.PutUint64(key, cnt)
			b.Put(key, line)
			if err == io.EOF || cnt == 400 {
				break
			}
			if cnt > 40 && cnt < 470 {
				continue
			}
			fmt.Println(cnt, string(line))
			word := string(b.Get(key))

			word = string(b.Get([]byte(string(2))))
			fmt.Println(string(word))
			word = string(b.Get([]byte(string(3))))
			fmt.Println(string(word))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
