package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func FeeCheckerHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "FEECHECKER",
    "msg": "Check those fees",
  })
  
}