package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

func RawDecodeHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "name": "RAWDECODER",
    "msg": "Decode those raw txs",
  })
  
}