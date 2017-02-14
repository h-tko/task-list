package controllers

import (
	"fmt"
	"github.com/h-tko/task-list/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TaskController struct {
	BaseController
}

func (this *TaskController) Regist(c echo.Context) error {
	this.BeforeFilter(c)

	title := c.FormValue("title")
	body := c.FormValue("body")

	taskModel := models.NewTask()
	taskModel.Title = title
	taskModel.Body = body
	taskModel.RegistMemberID = 1

	taskModel.Regist()

	this.SetResponse("status", "ok")

	return this.JSON(c, http.StatusOK)
}

func (this *TaskController) Detail(c echo.Context) error {
	this.BeforeFilter(c)

	taskID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	taskModel := models.NewTask()
	taskModel.One(taskID)

	this.SetResponse("Task", taskModel)

	return this.JSON(c, http.StatusOK)
}
