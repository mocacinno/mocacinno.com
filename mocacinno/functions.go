package mocacinno

import (
	//"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	//"github.com/btcsuite/btcutil"
	"log"
	//"fmt"
	"gopkg.in/ini.v1"
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

func Blockcount(client *rpcclient.Client) int64 {
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	return blockCount
}
