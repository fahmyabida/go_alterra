package controller

import "github.com/labstack/echo/v4"

type IController interface {
	GetAllData(c echo.Context) error
	Create(c echo.Context) error
}

type ProductController struct {
}

func NewProductController() IController {
	return &ProductController{}
}

func (pc ProductController) GetAllData(c echo.Context) error {
	panic("implement me!")
}

func (ctrl ProductController) Create(c echo.Context) error {
	panic("implement me!")
}
