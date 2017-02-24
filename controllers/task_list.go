package controllers

import (
	"github.com/h-tko/task-list/models"
	"github.com/labstack/echo"
	"net/http"
)

type TaskListController struct {
	BaseController
}

func (this *TaskListController) Index(c echo.Context) error {
	this.BeforeFilter(c)

	taskModel := models.NewTask()
	tasks := taskModel.All()

	this.SetResponse("tasks", tasks)

	return this.Render(c, http.StatusOK, "index.html")
}

func (this *TaskListController) Search(c echo.Context) error {
	this.BeforeFilter(c)

	freeword := c.FormValue("freeword")

	taskModel := models.NewTask()
	var tasks []*models.Task

	if len(freeword) > 0 {
		tasks = taskModel.Search(freeword)
	} else {
		tasks = taskModel.All()
	}

	this.SetResponse("tasks", tasks)

	return this.JSON(c, http.StatusOK)
}

func (this *TaskListController) Tasks(c echo.Context) error {
	this.BeforeFilter(c)

	taskModel := models.NewTask()
	tasks := taskModel.All()

	this.SetResponse("tasks", tasks)

	return this.JSON(c, http.StatusOK)
}
