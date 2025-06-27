package models

type User struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Updated int64  `gorm:"autoUpdateTime:nano" json:"updated"`
}
