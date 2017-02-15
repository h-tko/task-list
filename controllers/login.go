package controllers

import (
	"github.com/h-tko/task-list/models"
	"github.com/labstack/echo"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login(c echo.Context) error {
	this.BeforeFilter(c)

	mailAddress := c.FormValue("mail_address")
	password := c.FormValue("password")

	memberModel := models.NewMember()
	memberModel.One(mailAddress, password)

	if memberModel.ID < 1 {
		this.SetResponse("err", "アカウントが見つかりません。")
	} else {
		this.SetResponse("Member", memberModel)
	}

	return this.JSON(c, http.StatusOK)
}
