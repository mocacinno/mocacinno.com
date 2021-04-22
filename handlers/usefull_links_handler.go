package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func UsefullLinksHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "USEFULLINKS",
    "msg": "usefull links",
  })
  
}