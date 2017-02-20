package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Index(c echo.Context) error {
	return this.JSON(c, http.StatusOK)
}

func (this *AccountController) New(c echo.Context) error {
	return this.JSON(c, http.StatusOK)
}
