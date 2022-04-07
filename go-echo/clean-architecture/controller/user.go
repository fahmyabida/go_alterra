package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	iUserRepo repository.IUserRepo
}

func NewUserController(iUserRepo repository.IUserRepo) UserController {
	return UserController{iUserRepo: iUserRepo}
}

func (ctrl UserController) GetAllData(c echo.Context) error {
	users, err := ctrl.iUserRepo.GetAllUser()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (ctrl UserController) Create(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err := ctrl.iUserRepo.InsertUser(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}
