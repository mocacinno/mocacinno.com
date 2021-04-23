package mocacinno

import (
  "net/http"
  "github.com/labstack/echo"
  //"fmt"
  //"html/template"
)

func FeeCheckerHandler(c echo.Context) error { 
  txid := c.FormValue("txid")
  /*
  client := Client()
	blockcount := Blockcount(client)
	fmt.Printf("%v\n", blockcount) 
  */
  return c.String(http.StatusOK, txid) 
  
}
