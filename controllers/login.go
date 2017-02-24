package controllers

import (
	"github.com/echo-contrib/sessions"
	"github.com/h-tko/task-list/models"
	"github.com/labstack/echo"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login(c echo.Context) error {
	session := sessions.Default(c)

	mailAddress := c.FormValue("mail_address")
	password := c.FormValue("password")

	memberModel := models.NewMember()
	memberModel.One(mailAddress, password)

	if memberModel.ID < 1 {
		this.SetResponse("err", "アカウントが見つかりません。")
	} else {
		this.SetResponse("MemberID", memberModel.ID)
		this.SetResponse("MemberName", memberModel.Name)

		session.Set("MemberID", memberModel.ID)
		session.Set("MemberName", memberModel.Name)
		session.Save()
	}

	return this.JSON(c, http.StatusOK)
}

func (this *LoginController) Logout(c echo.Context) error {
	session := sessions.Default(c)

	session.Delete("MemberID")
	session.Delete("MemberName")
	session.Save()

	return this.JSON(c, http.StatusOK)
}
