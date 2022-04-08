package model

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
