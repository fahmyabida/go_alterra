package repository

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAllUser() ([]model.User, error)
	InsertUser(user model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db}
}

func (r UserRepo) GetAllUser() ([]model.User, error) {
	users := []model.User{}
	err := r.db.Find(&users).Error // SELECT * FROM user
	if err != nil {
		fmt.Println("error while GetAllUser", err)
	}
	return users, err
}

func (repo UserRepo) InsertUser(user model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Println("error while InsertUser", err)
	}
	return err
}
