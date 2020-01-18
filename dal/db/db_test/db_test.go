package mydb_test

import (
	"testing"

	mydb "github.com/hixinj/MOSAD_FINAL/dal/db"
)

func TestGetWords(t *testing.T) {
	mydb.GetWords(1)
}
