package database

import (
	"alterra/go_alterra/crud-basic/config"
	"alterra/go_alterra/crud-basic/models"
	"fmt"
)

func GetAllUser() ([]models.User, error) {
	var listUser []models.User
	// SELECT * FROM users;
	err := config.DB.Find(&listUser).Error
	if err != nil {
		fmt.Println(err)
	}
	return listUser, err
}

func GetSingleUserById(id string) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

func DeleteUserById(id string) error {
	err := config.DB.Delete(&models.User{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewUser(user models.User) error {
	err := config.DB.Save(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateUserById(id string, user models.User) error {
	// UPDATE table_name SET column1 = value1, column2 = value2, ...
	// WHERE id = $id AND email = $email;
	err := config.DB.Model(&user).Where("id = ?", id).Updates(user).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
