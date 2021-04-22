package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func TimestampedMessageHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "TIMESTAMPEDMESSGEA",
    "msg": "create a timestamped message",
  })
  
}