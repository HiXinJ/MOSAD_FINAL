package main

import (
	"fmt"

	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
)

func main() {
	fmt.Println(mydb.GetWords(10))
}
