package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func SigCheckHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "SIGCHECK",
    "msg": "Check those signatures",
  })
  
}