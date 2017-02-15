package main

import (
	"fmt"
	"github.com/echo-contrib/sessions"
	"github.com/h-tko/task-list/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pelletier/go-toml"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	models.InitDatabase()
	defer models.CloseDatabase()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config, err := toml.LoadFile("./config/app.toml")

	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	cacheHost := config.Get("cache.host").(string)
	cachePort := config.Get("cache.port").(string)
	cacheSecret := config.Get("cache.secret").(string)

	store, err := sessions.NewRedisStore(16, "tcp", fmt.Sprintf("%s:%s", cacheHost, cachePort), "", []byte(cacheSecret))

	if err != nil {
		panic(err)
	}

	e.Use(sessions.Sessions("echosession", store))

	e.Pre(middleware.AddTrailingSlash())

	e.Static("/static", "assets")

	t := &Template{
		templates: template.Must(template.New("").Funcs(TemplateHelpers).ParseGlob("views/*.html")),
	}

	e.Renderer = t

	routes(e)

	port := config.Get("application.port").(string)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
