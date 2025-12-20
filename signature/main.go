package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func LoadKeys(privateKeyPEM, publicKeyPEM []byte) (*RSAKeyPair, error) {
	privateBlock, _ := pem.Decode(privateKeyPEM)
	if privateBlock == nil {
		return nil, errors.New("failed to parse private key PEM")
	}
	privateKeyAny, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := privateKeyAny.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not a RSA private key")
	}

	publicBlock, _ := pem.Decode(publicKeyPEM)
	if publicBlock == nil {
		return nil, errors.New("failed to parse public key PEM")
	}
	publicKeyAny, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyAny.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not a RSA public key")
	}

	return &RSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func (k *RSAKeyPair) Sign(message string) (string, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, k.PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (k *RSAKeyPair) Verify(message, signatureB64 string) error {
	signature, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(message))
	return rsa.VerifyPKCS1v15(k.PublicKey, crypto.SHA256, hashed[:], signature)
}

func main() {
	privateKeyPEM, err := os.ReadFile("signature.pem")
	if err != nil {
		panic(err)
	}
	publicKeyPEM, err := os.ReadFile("signature.pub.pem")
	if err != nil {
		panic(err)
	}

	keys, err := LoadKeys(privateKeyPEM, publicKeyPEM)
	if err != nil {
		panic(err)
	}

	message := "Order 123456"
	signature, err := keys.Sign(message)
	if err != nil {
		panic(err)
	}

	err = keys.Verify(message, signature)
	if err != nil {
		fmt.Println("Invalid Signature!")
	} else {
		fmt.Println("Trusted Message.")
	}
}
