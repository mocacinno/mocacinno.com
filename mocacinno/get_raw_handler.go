package mocacinno

import (
  "net/http"
  "github.com/labstack/echo"
  "gopkg.in/ini.v1"
  "fmt"
  "os"
  "html/template"
  //"github.com/btcsuite/btcutil"
  //"github.com/davecgh/go-spew/spew"
)

func GetRawHandler(c echo.Context) error { 
  cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read ini file: %v", err)
        os.Exit(1)
    }
  txid := c.FormValue("txid")
  output := fmt.Sprintf("<ul><li>your tx id is %v</li>", txid)
  client := Client(cfg)
  rawtransaction, err := Rawtx(client, txid)
  /*
  spew.Dump(rawtransaction)
  */
  if err != nil {
	
    return c.String(http.StatusOK, txid + " not found on my node!!! (" + err.Error()  + ")") 
  }
  rawtxhex := TxToHex(rawtransaction.MsgTx())
  output += fmt.Sprintf("<li>your raw transaction is %v</li>",rawtxhex)
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "Get raw transaction",
    "description": "here is raw transaction "+txid+"!!!",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc, get, raw transaction",
    "content": template.HTML(output),
  })

  
}
