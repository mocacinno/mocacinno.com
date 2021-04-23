package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func GetRawHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "GETRAW",
    "msg": "Check those raw transactions",
  })
  
}