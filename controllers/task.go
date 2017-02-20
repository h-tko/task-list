package controllers

import (
	"fmt"
	"github.com/echo-contrib/sessions"
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

	taskCommentModel := models.NewTaskComment()
	taskComments := taskCommentModel.ListByTaskID(taskID)

	memberModel := models.NewMember()
	memberModel.FromID(this.response["MemberID"].(uint))

	this.SetResponse("Task", taskModel)
	this.SetResponse("TaskComments", taskComments)
	this.SetResponse("Member", memberModel)

	return this.JSON(c, http.StatusOK)
}

func (this *TaskController) SendComment(c echo.Context) error {
	this.BeforeFilter(c)

	session := sessions.Default(c)
	memberID := session.Get("MemberID").(uint)

	id, err := strconv.Atoi(c.FormValue("id"))

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	comment := c.FormValue("comment")

	taskComment := models.NewTaskComment()
	taskComment.TaskID = uint(id)
	taskComment.Comment = comment
	taskComment.MemberID = memberID

	if err := c.Bind(taskComment); err != nil {
		fmt.Printf("%v\n", err)

		return err
	}

	//    if errs := c.Validate(taskComment); errs != nil {
	//
	//        resErrors := ValidationErrors(errs)
	//
	//        fmt.Printf("%v", resErrors)
	//
	//        this.SetResponse("err", resErrors)
	//
	//        return this.JSON(c, http.StatusOK)
	//    }

	taskComment.Regist()

	return this.JSON(c, http.StatusOK)
}
