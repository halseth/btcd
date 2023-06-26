package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type config struct {
	Leaves string `short:"l" long:"leaves" description:"space separated string of hex values to commit to (must be power of 2)"`
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
	leaves := strings.Split(cfg.Leaves, " ")

	fmt.Println("input", cfg.Leaves)

	var level [][32]byte
	for _, l := range leaves {
		h, err := hex.DecodeString(l)
		if err != nil {
			return err
		}

		hash := sha256.Sum256(h)
		level = append(level, hash)
	}

	for len(level) > 1 {
		if len(level)%2 != 0 {
			return fmt.Errorf("invalid number of leaves")
		}

		s := "level"
		for _, l := range level {
			s += fmt.Sprintf(" %x", l)
		}

		fmt.Println(s)

		var nextLevel [][32]byte
		for i := 1; i < len(level); i += 2 {
			item := make([]byte, 64)
			copy(item[:32], level[i-1][:])
			copy(item[32:], level[i][:])
			hash := sha256.Sum256(item)

			nextLevel = append(nextLevel, hash)
		}

		level = nextLevel

	}

	fmt.Printf("Merkle root: %x\n", level[0])

	return nil
}
