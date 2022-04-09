package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	iUserRepo repository.IUserRepo
}

func NewUserController(iUserRepo repository.IUserRepo) UserController {
	return UserController{iUserRepo: iUserRepo}
}

func (ctrl UserController) GetAllData(c echo.Context) error {
	user, err := utils.ParsingJWT(c)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	} else if user.Role != "admin" {
		return c.JSON(200, echo.Map{
			"error": "restricted (*only for admin)",
		})
	}
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
	userPayload, err := utils.ParsingJWT(c)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	} else if userPayload.Role != "admin" {
		return c.JSON(200, echo.Map{
			"error": "restricted (*only for admin)",
		})
	}
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err = ctrl.iUserRepo.InsertUser(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (ctrl UserController) GetSingleData(c echo.Context) error {
	payloadUser, err := utils.ParsingJWT(c)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err,
		})
	}
	stringId := c.QueryParam("id")
	if stringId != "" {
		idRequest, err := strconv.Atoi(stringId)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err,
			})
		}
		/*
			admin : bisa lihat dirinya atau lainnya,
			user  : hanya bisa lihat dirinya
		*/
		if payloadUser.Role != "admin" && payloadUser.Role != "user" {
			return c.JSON(500, map[string]interface{}{
				"error": "role is invalid",
			})
		} else if payloadUser.Role == "user" && payloadUser.Id != idRequest {
			return c.JSON(500, map[string]interface{}{
				"error": "user anda memiliki role user, tidak diizinkan untuk melihat user lain",
			})
		}
		user, err := ctrl.iUserRepo.GetUserById(idRequest)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(200, map[string]interface{}{
			"data": user,
		})
	} else if stringId == "" {
		user, err := ctrl.iUserRepo.GetUserById(payloadUser.Id)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(200, map[string]interface{}{
			"data": user,
		})
	}
	return nil
}
