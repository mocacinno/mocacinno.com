package mocacinno

import (
	//"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"log"
	"github.com/btcsuite/btcd/wire"
	"fmt"
	"gopkg.in/ini.v1"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"bytes"
	"encoding/hex"
)

func Client(cfg *ini.File) *rpcclient.Client {
	connCfg := &rpcclient.ConnConfig{
		Host:         cfg.Section("bitcoind").Key("host").String() + ":" + cfg.Section("bitcoind").Key("port").String(),
		User:         cfg.Section("bitcoind").Key("user").String() ,
		Pass:         cfg.Section("bitcoind").Key("password").String() ,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client

}
/*
func EmptyTx() (*wire.MsgTx, error) {
	return wire.NewMsgTx(wire.TxVersion), nil
 }
*/
func Blockcount(client *rpcclient.Client) int64 {
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	return blockCount
}

func Rawtx(client *rpcclient.Client, txid string) (*btcutil.Tx, error) {
	txid_hash, err := chainhash.NewHashFromStr(txid)
	/*
	fmt.Printf("txid_hash (functions) : %+v\n\n", txid_hash)
	fmt.Printf("txid_hash type (functions) : %T\n\n", txid_hash)
	*/
	rawtransaction, err := client.GetRawTransaction(txid_hash)
	/*
	fmt.Printf("rawtransaction (functions) : %+v\n\n", rawtransaction)
	fmt.Printf("rawtransaction type (functions) : %T\n\n", rawtransaction)
	*/
	if err != nil {
		fmt.Printf("error : %v\n\n", err)
		return nil, err
	}
	return rawtransaction, nil
}

func TxToHex(tx *wire.MsgTx) string {
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	tx.Serialize(buf)
	return hex.EncodeToString(buf.Bytes())
}
