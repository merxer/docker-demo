package main

import (
  "os"
  "io"
  "net/http"
  "html/template"

  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
)

type (
  Template struct {
    templates *template.Template
  }

  Info struct {
    Hostname string
    Version string
  }
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func GetHostname() string{
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }
  return name
}

func main() {
  e := echo.New()
  e.Use(middleware.Static("assets"))
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  info := Info{
    Hostname: GetHostname(),
    Version: "0.0.1",
  }

  t := &Template{
      templates: template.Must(template.ParseGlob("*.html")),
  }
  e.SetRenderer(t)

  // Route => handler
  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index", info)
  })

  // Start server
  e.Run(standard.New(":8888"))
}
