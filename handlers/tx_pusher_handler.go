package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func TxPusherHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "TXPUSH",
    "msg": "Push those transactions",
  })
  
}