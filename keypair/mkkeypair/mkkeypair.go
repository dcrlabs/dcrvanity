package main

import (
	"fmt"

	"github.com/dcrlabs/dcrvanity/keypair"
	"github.com/dcrlabs/dcrvanity/wif"
	"github.com/decred/dcrd/chaincfg"
)

func main() {
	params := &chaincfg.MainNetParams
	priv, pub, addr, WIF, _ := keypair.KeyPairAddress(params)
	fmt.Printf("pubkey: %x\n", pub.SerializeCompressed())
	fmt.Printf("pubkey hash: %x\n", wif.Hash160(pub.SerializeCompressed()))
	fmt.Printf("pubkey hash address (P2PKH): %s\n", addr)
	fmt.Print(keypair.PrintPrivateKey(priv))
	fmt.Printf("Wallet import format (WIF): %v\n", WIF)

	// priv2 := keypair.PrivKey(wif.Curve(), priv.D)
	// fmt.Print(keypair.PrintPrivateKey(priv2))
	// if keypair.PrintPrivateKey(priv2) != keypair.PrintPrivateKey(priv) {
	// 	panic("priv != priv2")
	// }
}
