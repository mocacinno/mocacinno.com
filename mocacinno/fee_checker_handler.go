package mocacinno

import (
  "net/http"
  "github.com/labstack/echo"
  //"html/template"
)

func FeeCheckerHandler(c echo.Context) error { 
  txid := c.FormValue("txid")
  return c.String(http.StatusOK, txid) 
  
}
