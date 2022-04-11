package model

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UserCostum struct { // table user_custom
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// TableName overrides the table name used by User to `profiles`
func (UserCostum) TableName() string {
	return "users"
}
