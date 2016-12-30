// Package diffiehellman implements a library for the Diffie-Hellman key exchange.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random private key.
func PrivateKey(p *big.Int) *big.Int {
	privateKey, _ := rand.Int(rand.Reader, big.NewInt(0).Sub(p, big.NewInt(2)))

	return privateKey.Add(privateKey, big.NewInt(2))
}

// PublicKey generates a public key for the given private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair generates a key pair with a random private key.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)

	return private, public
}

// SecretKey generates a secret key for the given key pair.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
