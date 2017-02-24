package controllers

import (
	"fmt"
	"github.com/echo-contrib/sessions"
	"github.com/h-tko/task-list/models"
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
	session := sessions.Default(c)

	member := models.NewMember()
	member.Name = c.FormValue("name")
	member.MailAddress = c.FormValue("mail_address")
	member.Password = c.FormValue("password")

	if err := c.Bind(member); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if errs := c.Validate(member); errs != nil {
		resErrors := ValidationErrors(errs)
		fmt.Printf("%v\n", resErrors)
		this.SetResponse("err", resErrors)

		return this.JSON(c, http.StatusOK)
	}

	if member.Exists(member.MailAddress) {
		this.SetResponse("err", map[string]string{"key": "MailAddress", "message": "そのメールアドレスのアカウントはすでに存在しています。"})

		return this.JSON(c, http.StatusOK)
	}

	member.Regist()
	member.One(member.MailAddress, member.Password)

	this.SetResponse("MemberID", member.ID)
	this.SetResponse("MemberName", member.Name)
	session.Set("MemberID", member.ID)
	session.Set("MemberName", member.Name)
	session.Save()

	return this.JSON(c, http.StatusOK)
}
