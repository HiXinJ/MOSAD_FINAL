package model

type User struct {
	UserID   int64  `gorm:"user_id", json:"user_id"`
	Username string `gorm:"user_name",json:"user_name"`
	Password string `gorm:"password",json:"password"`
}

func TableName() string {
	return "user"
}
