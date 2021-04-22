package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func ValidateSigHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "VALIDATESIG",
    "msg": "Validate that signature",
  })
  
}