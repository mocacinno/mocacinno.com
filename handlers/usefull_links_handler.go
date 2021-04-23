package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func UsefullLinksHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "USEFULLINKS",
    "msg": "usefull links",
  })
  
}