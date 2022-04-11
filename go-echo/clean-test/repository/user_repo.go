package repository

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAllUser() ([]model.UserCostum, error)
	GetUserByUsername(username string) (model.User, error)
	InsertUser(user model.User) error
	GetUserById(id int) (model.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db}
}

func (r UserRepo) GetAllUser() ([]model.UserCostum, error) {
	users := []model.UserCostum{}
	err := r.db.Find(&users).Error // SELECT * FROM users;
	if err != nil {
		fmt.Println("error while GetAllUser", err)
	}
	return users, err
}

func (repo UserRepo) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := repo.db.Where("username = ?", username).First(&user).Error // SELECT * FROM users WHERE username = ? LIMIT 1;
	if err != nil {
		fmt.Println("error while GetUserByUsernameAndPassword", err)
	}
	return user, err
}

func (repo UserRepo) InsertUser(user model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Println("error while InsertUser", err)
	}
	return err
}

func (r UserRepo) GetUserById(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println("error while GetUserById", err)
	}
	return user, err
}
