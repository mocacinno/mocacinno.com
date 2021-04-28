package mocacinno

import (
  "net/http"
  "github.com/labstack/echo"
  "gopkg.in/ini.v1"
  "fmt"
  "os"
  "html/template"
)

func FeeCheckerHandler(c echo.Context) error { 
  cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read ini file: %v", err)
        os.Exit(1)
    }
  txid := c.FormValue("txid")
  output := fmt.Sprintf("<ul><li>your tx id is %v</li>", txid)
  client := Client(cfg)
  

  fee_1, err := client.EstimateSmartFee(6, nil)
  if err != nil {
    return c.String(http.StatusOK, txid + " my node could not determine the smart fee (6 blocks)... Retry in a couple of hours, or contact me directly!!!") 
  }
  fee_1_result := fee_1.FeeRate
  fee_1_per_vbyte := *fee_1_result * float64(100000)
  output += fmt.Sprintf("<li>the optimal fee per vbyte to be ~rather~ sure your tx ends up in the next ~6 blocks is: %v sat/vbyte</li>", fee_1_per_vbyte)
  mempoolentry, err := client.GetMempoolEntry(txid)
  if err != nil {
	erroutput := fmt.Sprintf("Oops, it seems like your tx (%v) could not be found in my nodes mempool... This means it was either to old, the fee was so low my node rejected your transaction, your transaction was nonstandard, your transaction was doublespending the same unspent output as an other (unconfirmed) transaction or the amount of transactions being broadcasted to me was so high my node started pruning due to memory constraints... It's also possible my node was restarted recently, and my mempool was emptied in the process...", txid) 
    return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "(in)Sufficient Fee checker",
    "description": "txid "+txid+" Not found in my mempool!!!",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc, fee, checker",
    "content": template.HTML(erroutput),
  })
  }
  vsize := mempoolentry.VSize
  output += fmt.Sprintf("<li>your vsize (in vbytes) is %v</li>", vsize)
  payedfee := mempoolentry.Fee * float64(100000000)
  output += fmt.Sprintf("<li>your payed a fee of  %v sats</li>", payedfee)
  payedsatpervbyte := int32(payedfee) / vsize
  output += fmt.Sprintf("<li>which boils down to %v sats/vbyte</li>", payedsatpervbyte)
  if int32(payedsatpervbyte) > int32(fee_1_per_vbyte) {
    output += fmt.Sprintf("<li>your fee is high enough... yeey!!!</li></ul>")
  } else {
    output += fmt.Sprintf("<li>Oooooooops.... that's not enough i'm afraid... I guess you'll have to wait untill the mempool clears a little bit OR if you opted in RBF you can bump your fee, or you can do a CPFP, or pay a big mining pool to include your tx in the block they're trying to solve right now</li></ul>")
  }

  return c.Render(http.StatusOK, "layout.html", map[string]interface{}{
    "title": "(in)Sufficient Fee checker",
    "description": "here, you'll see the output of my feechecker for tx "+txid+"!!!",
    "keywords": "mocacinno, usefull, bitcoin, tools, btc, fee, checker",
    "content": template.HTML(output),
  })

  
}
