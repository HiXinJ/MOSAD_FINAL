package model

type User struct {
	UserID       int64            `json:"user_id"`
	UserName     string           `json:"user_name"`
	Password     string           `json:"password"`
	LearnedWords map[string]int64 `json:"learned_words"`
	DaKa         []Date           `json:"daka"`
	Head         []byte           `json:"head`
}

func TableName() string {
	return "user"
}
