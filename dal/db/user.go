package mydb

import "time"

type User struct {
	UserID       int64            `json:"user_id"`
	UserName     string           `json:"user_name"`
	Password     string           `json:"password"`
	LearnedWords map[string]int64 `json:"learned_words"`
	DaKa         []Date           `json:"daka"`
	Head         []byte           `json:"head`
	lastUpdate   Date
	NewWords     map[string]int64
}

func TableName() string {
	return "user"
}

func (user *User) LearnWord(word string) {
	user.LearnedWords[word] = 1
}

func (user *User) UpdateNewWords(size int64) []string {
	// 每天产生一次新词
	if user.lastUpdate.Equals(new(Date).Reset(time.Now())) {
		return nil
	}

	var today Date
	today.Reset(time.Now())
	user.lastUpdate = today

	// 删除以前的新词，重新生成新词
	user.NewWords = make(map[string]int64)
	WordsList := FilterWords(size, func(word string) bool {
		if _, ok := user.LearnedWords[word]; ok {
			return false
		}
		return true
	})
	for _, word := range WordsList {
		user.NewWords[word] = 1
	}
	return WordsList
}
