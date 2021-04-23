package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func BlockHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "BLOCK",
    "msg": "Blocksssss",
  })
  
}