package model

type Favour struct {
	ArticleID int64   `gorm:"column:article_id" json:"article_id`
	Sum       int64   `gorm: "column:sum" json:"sum"`
	UserID    []int64 `gorm:"column:user_id" json:"user_id"`
}

func (Favour) TableName() string {
	return "favour"
}
