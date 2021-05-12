package handler

import (
  "net/http"
  "github.com/labstack/echo"
  "html/template"
)

func GetRawHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "Get raw transaction",
    "description": "get your raw transaction from my mempool (if it exists in my mempool ofcourse)",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc, raw, transaction, get",
    "content": template.HTML("<form action=\"..\\.\\page\\getraw\" method=\"post\">insert your tx id here: <input type=\"text\" name=\"txid\"><br><input type=\"submit\" name=\"submit\" value=\"get my raw transaction\"></form>"),
  })
  
}