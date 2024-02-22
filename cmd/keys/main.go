package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	flags "github.com/jessevdk/go-flags"
)

type config struct {
	Num int `short:"n" long:"num" description:"number of keys to generate"`
}

var cfg = config{}

func main() {
	if _, err := flags.Parse(&cfg); err != nil {
		fmt.Println(err)
		return
	}

	err := run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func run() error {
	// Random key.

	for i := 0; i < cfg.Num; i++ {
		privKey, err := btcec.NewPrivateKey()
		if err != nil {
			return err
		}
		privKeyBytes := privKey.Serialize()
		pubKey := privKey.PubKey()
		pubKeyBytes := schnorr.SerializePubKey(pubKey)
		fmt.Printf("privkey: %x pubkey: %x\n", privKeyBytes, pubKeyBytes)
	}
	return nil
}
