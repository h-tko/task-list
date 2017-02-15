package main

import (
	. "github.com/h-tko/task-list/controllers"
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/", new(TaskListController).Index)
	e.GET("/tasks/", new(TaskListController).Tasks)
	e.POST("/search/", new(TaskListController).Search)
	e.POST("/regist_task/", new(TaskController).Regist)
	e.GET("/detail/:id/", new(TaskController).Detail)
	e.POST("/login/", new(LoginController).Login)
}
