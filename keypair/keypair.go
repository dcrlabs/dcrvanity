// Copyright (c) 2015, 2018 The Decred Developers

package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/dcrlabs/dcrvanity/wif"
	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/chaincfg/chainec"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/decred/dcrd/dcrutil"
)

// KeyPairAddress generates a random private key for the specified network using
// the curve provided by wif.Curve(). Return values are the ecdsa PrivateKey,
// the secp256k1.PublicKey, the P2PKH address, and a WIF structure for importing
// into a wallet.
func KeyPairAddress(params *chaincfg.Params) (*ecdsa.PrivateKey, *secp256k1.PublicKey, dcrutil.Address, *wif.WIF, error) {
	curve := wif.Curve()

	// Generate public-private key pair
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// ecdsa.PublicKey with serialization functions
	pub := secp256k1.PublicKey{
		Curve: curve,
		X:     key.PublicKey.X,
		Y:     key.PublicKey.Y,
	}

	// PubKeyHashAddrID (Ds) followed by ripemd160 hash of secp256k1 pubkey
	addr, err := dcrutil.NewAddressPubKeyHash(wif.Hash160(pub.SerializeCompressed()),
		params, chainec.ECTypeSecp256k1)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	priv := secp256k1.PrivateKey(*key)
	privWif := wif.NewWIF(priv, params)

	return key, &pub, addr, privWif, nil
}

// PrivKey creates a private key for the specified elliptic curve and scalar.
func PrivKey(c elliptic.Curve, d *big.Int) *ecdsa.PrivateKey {
	X, Y := c.ScalarBaseMult(d.Bytes())
	return &ecdsa.PrivateKey{
		D: d,
		PublicKey: ecdsa.PublicKey{
			Curve: c,
			X:     X,
			Y:     Y,
		},
	}
}

func PrintPrivateKey(p *ecdsa.PrivateKey) string {
	return fmt.Sprintf("Private key (curve = secp256k1.S256):\n\tD = %v,\n\tX = %v,\n\tY = %v.\n", p.D, p.X.String(), p.Y)
}
