package main

import (
	"html/template"
	"io"
	"./handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates *template.Template
  }

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
  }

  func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}, time=${time_unix}, remote_ip=${remote_ip}, host=${host}, path=${path}, protocol=${protocol}, user_agent=${user_agent}, error=${error}, bytes_in=${bytes_in}, bytes_out=${bytes_out}, header=${header}, query=${query}, form=${form}\n\n",
    }))

	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("views/*.html")),
	  }
	
	e.Static("/static", "staticfiles")
	e.GET("/", handler.HomeHandler)
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
