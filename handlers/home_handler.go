package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func HomeHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "Mocacinno's usefull bitcoin tools",
    "description": "The 2021 version of mocacinno's usefull bitcoin tools",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc",
    "content": "This website contains the new version of mocacinno's usefull bitcoin tools",
  })
  
}