package repository

import (
	"belajar-go-echo/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUser(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username"}).
				AddRow(1, "fahmy").AddRow(2, "abida"))
	// result query GORM nya seperti apa
	listUser, err := tu.IFaceUserRepo.GetAllUser()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(listUser)
}

func TestGetUserByUsername(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username", "password", "name", "role"}).
				AddRow(1, "fahmy", "1234", "fahmy", "admin"))
	// result query GORM nya seperti apa
	user, err := tu.IFaceUserRepo.GetUserByUsername("fahmy")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestInsertUser(t *testing.T) {
	tu := NewTestUnit()
	user := model.User{
		Id:       1,
		Username: "abida@mail.com",
		Password: "5678",
		Name:     "abida",
		Role:     "user",
	}
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectBegin()
	tu.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`username`,`password`,`name`,`role`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs(user.Username, user.Password, user.Name, user.Role, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	tu.Mock.ExpectCommit()
	// result query GORM nya seperti apa
	err := tu.IFaceUserRepo.InsertUser(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success insert")
}

func TestGetUserById(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username", "password", "name", "role"}).
				AddRow(1, "fahmy", "1234", "fahmy", "admin"))
	// result query GORM nya seperti apa
	user, err := tu.IFaceUserRepo.GetUserById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}
