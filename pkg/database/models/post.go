package models

type Posts struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    Users  `gorm:"foreignKey:UserID" json:"author"`
}
