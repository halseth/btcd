package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/txscript"
	flags "github.com/jessevdk/go-flags"
)

type config struct {
	Key    string `short:"k" long:"key" description:"key to use"`
	Tweak  string `short:"t" long:"tweak" description:"tweak to set"`
	Tweak2 string `short:"u" long:"tweak2" description:"second tweak to set"`
}

var cfg = config{}

func main() {
	fmt.Println("hello")

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

	var keyBytes []byte
	if cfg.Key == "" {
		privKey, err := btcec.NewPrivateKey()
		if err != nil {
			return err
		}
		pubKey := privKey.PubKey()
		keyBytes = schnorr.SerializePubKey(pubKey)
	} else {
		var err error
		keyBytes, err = hex.DecodeString(cfg.Key)
		if err != nil {
			return err
		}

	}
	fmt.Println("key:", hex.EncodeToString(keyBytes))
	pubKey, err := schnorr.ParsePubKey(keyBytes)
	if err != nil {
		return err
	}

	data, err := hex.DecodeString(cfg.Tweak)
	if err != nil {
		return err
	}
	fmt.Println("tweak:", hex.EncodeToString(data))

	// Tweak pubkey with data.
	tweaked := txscript.ComputeTaprootOutputKey(pubKey, data)
	tweakedBytes := schnorr.SerializePubKey(tweaked)
	fmt.Println("tweaked key:", hex.EncodeToString(tweakedBytes))

	// Optionally tweak with second tweak.
	if cfg.Tweak2 != "" {
		data, err := hex.DecodeString(cfg.Tweak2)
		if err != nil {
			return err
		}
		fmt.Println("tweak2:", hex.EncodeToString(data))

		tweaked2 := txscript.ComputeTaprootOutputKey(tweaked, data)
		tweakedBytes2 := schnorr.SerializePubKey(tweaked2)
		fmt.Println("tweaked key2:", hex.EncodeToString(tweakedBytes2))
	}

	return nil
}
