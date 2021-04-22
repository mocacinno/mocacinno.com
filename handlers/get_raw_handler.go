package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func GetRawHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "GETRAW",
    "msg": "Check those raw transactions",
  })
  
}