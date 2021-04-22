package handler

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "html/template"
)

func FeeCheckerHandler(c echo.Context) error { 
  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "(in)Sufficient Fee checker",
    "description": "a basic tool that will let you know if you added sufficient fees to your transaction, especially usefull when you see a transaction being unconfirmed for hours/days",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc, fee, checker",
    "content": template.HTML("<form action=\"..\\.\\page\\feechecker\" method=\"post\">insert your tx id here: <input type=\"text\" name=\"txid\"><br><input type=\"submit\" name=\"submit\" value=\"check my fee\"></form>"),
  })
  
}