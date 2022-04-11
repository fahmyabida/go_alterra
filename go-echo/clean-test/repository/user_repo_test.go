package repository

import (
	"belajar-go-echo/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUnit struct {
	mock          sqlmock.Sqlmock
	iFaceUserRepo IUserRepo
}

func NewTestUnit() TestUnit {
	tu := TestUnit{}
	// bersifat inisialisasi
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql-mock",
		ServerVersion:             "1.0.0",
		DSN:                       "mysql-mock",
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tu.mock = mock
	iFaceUserRepo := NewUserRepo(dbGorm)
	tu.iFaceUserRepo = iFaceUserRepo
	return tu
}

func TestGetAllUser(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username"}).
				AddRow(1, "fahmy").AddRow(2, "abida"))
	// result query GORM nya seperti apa
	listUser, err := tu.iFaceUserRepo.GetAllUser()
	if err != nil {
		t.Error(err)
	}
	t.Log(listUser)
}

func TestGetUserByUsername(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username", "password", "name", "role"}).
				AddRow(1, "fahmy", "1234", "fahmy", "admin"))
	// result query GORM nya seperti apa
	user, err := tu.iFaceUserRepo.GetUserByUsername("fahmy")
	if err != nil {
		t.Error(err)
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
	tu.mock.ExpectBegin()
	tu.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`username`,`password`,`name`,`role`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs(user.Username, user.Password, user.Name, user.Role, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	tu.mock.ExpectCommit()
	// result query GORM nya seperti apa
	err := tu.iFaceUserRepo.InsertUser(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success insert")
}
