package mocacinno

import (
  "net/http"
  "github.com/labstack/echo"
  "gopkg.in/ini.v1"
  "fmt"
  "os"
  "github.com/btcsuite/btcd/blockchain"
  "bytes"
  //"html/template"
)

func FeeCheckerHandler(c echo.Context) error { 
  cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read ini file: %v", err)
        os.Exit(1)
    }
  txid := c.FormValue("txid")
  client := Client(cfg)
  rawtx, err := Rawtx(client, txid)
  if err !=nil {
    return c.String(http.StatusOK, txid + " not found") 
  }
  
  fmt.Printf("client (handler) : %+v\n\n", client)
  fmt.Printf("client type (handler) : %T\n\n", client)

  fmt.Printf("rawtx (handler) : %+v\n\n", rawtx)
  fmt.Printf("rawtx type (handler) : %T\n\n", rawtx)
  
  msgtx := rawtx.MsgTx()
  
  fmt.Printf("msgtx (handler) : %+v\n\n", msgtx)
  fmt.Printf("msgtx type (handler) : %T\n\n", msgtx)
  

  txhex := TxToHex(msgtx)
  
  fmt.Printf("txhex (handler) : %+v\n\n", txhex)
  fmt.Printf("txhex type (handler) : %T\n\n", txhex)
  

  weight := blockchain.GetTransactionWeight(rawtx)
  
  fmt.Printf("weight (handler) : %+v\n\n", weight)
  fmt.Printf("weight type (handler) : %T\n\n", weight)
  
	blockcount := Blockcount(client)
	fmt.Printf("blockcount (handler) : %v\n", blockcount) 
  fmt.Printf("blockcount type(handler) : %T\n", blockcount) 

  baseSize := msgtx.SerializeSizeStripped()
  fmt.Printf("baseSize (handler) : %v\n", baseSize) 
  fmt.Printf("baseSize type(handler) : %T\n", baseSize) 
  totalSize := msgtx.SerializeSize()
  fmt.Printf("totalSize (handler) : %v\n", totalSize) 
  fmt.Printf("totalSize type(handler) : %T\n", totalSize) 
  altweight := int64((baseSize * (4 - 1)) + totalSize)
  fmt.Printf("altweight (handler) : %v\n", altweight) 
  fmt.Printf("altweight type(handler) : %T\n", altweight) 


  fee_1, err := client.EstimateSmartFee(1, nil)
  fee_1_result := fee_1.FeeRate
  fmt.Printf("fee_1_result (handler) : %v\n", *fee_1_result) 
  fmt.Printf("fee_1_result type(handler) : %T\n", *fee_1_result)

  
  //buf := make([]byte, 0, msgtx.Serialize())
	buf := bytes.NewBuffer(make([]byte, 0, msgtx.SerializeSize()))
	msgtx.Serialize(buf)

  decoderawtx, err := client.DecodeRawTransaction([]byte(buf.Bytes()))
  fmt.Printf("decoderawtx (handler) : %v\n", decoderawtx) 
  fmt.Printf("decoderawtx type(handler) : %T\n", decoderawtx)

  txsize := decoderawtx.Size 
  fmt.Printf("txsize (handler) : %v\n", txsize) 
  fmt.Printf("txsize type(handler) : %T\n", txsize)
  /*
    baseSize := msgTx.SerializeSizeStripped()
    totalSize := msgTx.SerializeSize()
  
    // (baseSize * 3) + totalSize
    return int64((baseSize * (WitnessScaleFactor - 1)) + totalSize)

  */


  return c.String(http.StatusOK, txid) 
  
}
