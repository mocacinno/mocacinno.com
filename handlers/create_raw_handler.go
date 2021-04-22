package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func CreateRawHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "CREATERAW",
    "msg": "Create raw transactions",
  })
  
}