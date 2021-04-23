package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func FeeEstimateHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "FEEESTIMATES",
    "msg": "Fee Estimates",
  })
  
}