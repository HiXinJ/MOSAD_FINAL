package mydb

import "time"

type User struct {
	UserID       int64            `json:"user_id"`
	UserName     string           `json:"user_name"`
	Password     string           `json:"password"`
	LearnedWords map[string]int64 `json:"learned_words"`
	DaKa         []Date           `json:"daka"`
	Head         string           `json:"head`
	lastUpdate   Date             `json:"lastUpdate"`
	NewWords     map[string]int64 `json:"new_words"`
}

func TableName() string {
	return "user"
}

func (user *User) LearnWord(word string) {
	user.LearnedWords[word] = 1
	delete(user.NewWords, word)
}

func (user *User) UpdateNewWords(size int64) []string {
	// 每天产生一次新词
	if user.lastUpdate.Equals(new(Date).Reset(time.Now())) {
		wordList := make([]string, 0)
		for item, n := range user.NewWords {
			if n != 0 {
				wordList = append(wordList, item)
			}
		}
		return wordList
	}

	var today Date
	today.Reset(time.Now())
	user.lastUpdate = today

	// 删除以前的新词，重新生成新词
	user.NewWords = make(map[string]int64)
	wordsList := FilterWords(size, func(word string) bool {
		if _, ok := user.LearnedWords[word]; ok {
			return false
		}
		return true
	})
	for _, word := range wordsList {
		user.NewWords[word] = 1
	}
	return wordsList
}
