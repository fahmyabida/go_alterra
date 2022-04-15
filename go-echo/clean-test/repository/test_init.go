package repository

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUnit struct {
	Mock          sqlmock.Sqlmock
	IFaceUserRepo IUserRepo
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
	tu.Mock = mock
	iFaceUserRepo := NewUserRepo(dbGorm)
	tu.IFaceUserRepo = iFaceUserRepo
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username"}).
				AddRow(1, "fahmy").AddRow(2, "abida"))
	return tu
}
