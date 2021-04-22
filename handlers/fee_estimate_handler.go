package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func FeeEstimateHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "FEEESTIMATES",
    "msg": "Fee Estimates",
  })
  
}