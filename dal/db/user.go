package mydb

import (
	"time"
)

type User struct {
	UserID        uint32           `json:"user_id"`
	UserName      string           `json:"user_name"`
	Password      string           `json:"password"`
	LearnedWords  map[string]int64 `json:"learned_words"`
	PendingReview map[string]int64 `json:"pending_review"`
	NewWords      map[string]int64 `json:"new_words"`
	DaKa          []Date           `json:"daka"`
	Head          string           `json:"head`
	LastUpdate    Date             `json:"lastUpdate"`
}

func TableName() string {
	return "user"
}

func (user *User) LearnWord(word string) {
	user.LearnedWords[word] = 1
	delete(user.NewWords, word)
}

func (user *User) UpdateNewWords(size int64) ([]string, bool) {
	// 每天产生一次新词
	if user.LastUpdate.Equals(new(Date).Reset(time.Now())) {
		wordList := make([]string, 0)
		for item, n := range user.NewWords {
			if n != 0 {
				wordList = append(wordList, item)
			}
		}
		return wordList, false
	}

	var today Date
	today.Reset(time.Now())
	user.LastUpdate = today

	// 删除以前的新词，重新生成新词
	user.NewWords = make(map[string]int64)
	wordsList := FilterWords(size, func(word string) bool {
		_, ok := user.LearnedWords[word]
		_, ok2 := user.PendingReview[word]
		if ok || ok2 {
			return false
		}
		return true
	})
	for _, word := range wordsList {
		user.NewWords[word] = 1
	}
	// PutUsers([]User{*user})
	return wordsList, true
}

func (user *User) AddLearnedWord(word string) {
	if user.NewWords[word] == 1 {
		// user.NewWords[word] = 0
		delete(user.NewWords, word)
		user.PendingReview[word] = 1
	} else if user.PendingReview[word] == 1 {
		// user.PendingReview[word] = 0
		delete(user.PendingReview, word)
		user.LearnedWords[word] = 1
	}
}

// GetReviews  从user.PendingReview 选取复习单词
func (user *User) GetReviews(size int64) []string {
	if size > int64(len(user.PendingReview)) {
		size = int64(len(user.PendingReview))
	}
	reviewList := make([]string, size)
	i := 0
	for k, v := range user.PendingReview {
		if v == 1 {
			reviewList[i] = k
			i++
		}
		if i == int(size) {
			break
		}
	}
	return reviewList
}
