// main.go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	b := x509.MarshalPKCS1PrivateKey(key)
	priv := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: b,
	})
	if err := os.WriteFile("key.pem", priv, 0600); err != nil {
		panic(err)
	}
	fmt.Printf("PVT KEy:%v", string(priv))

	b, _ = x509.MarshalPKIXPublicKey(key.Public())
	pub := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	})
	if err := os.WriteFile("pub.pem", pub, 0600); err != nil {
		panic(err)
	}
}
