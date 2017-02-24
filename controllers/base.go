package controllers

import (
	"github.com/echo-contrib/sessions"
	"github.com/labstack/echo"
)

type BaseController struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaH1          string
	MetaRobots      string

	response map[string]interface{}
}

func (this *BaseController) BeforeFilter(c echo.Context) {
	session := sessions.Default(c)
	memberID := session.Get("MemberID")
	memberName := session.Get("MemberName")

	this.clearResponse()

	if memberID != nil {
		this.SetResponse("MemberID", memberID)
		this.SetResponse("MemberName", memberName)
	}
}

func (this *BaseController) SetResponse(key string, val interface{}) {

	if this.response == nil {
		this.response = make(map[string]interface{})
	}

	this.response[key] = val
}

func (this *BaseController) Render(c echo.Context, status int, oFile string) error {
	if this.response == nil {
		this.response = make(map[string]interface{})
	}

	this.setMeta()

	return c.Render(status, oFile, this.response)
}

func (this *BaseController) JSON(c echo.Context, status int) error {
	if this.response == nil {
		this.response = make(map[string]interface{})
	}

	this.setMeta()

	return c.JSON(status, this.response)
}

func (this *BaseController) setMeta() {
	this.response["mt"] = this.MetaTitle
	this.response["md"] = this.MetaDescription
	this.response["mk"] = this.MetaKeywords
	this.response["mh1"] = this.MetaH1
	this.response["mr"] = this.MetaRobots
}

func (this *BaseController) clearResponse() {
	for key := range this.response {
		delete(this.response, key)
	}
}
