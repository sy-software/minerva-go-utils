package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// StrToPrivateKey parses private and public key string into a *rsa.PrivateKey instance
func StrToPrivateKey(private string, public string) (*rsa.PrivateKey, error) {
	privPem, _ := pem.Decode([]byte(private))

	if privPem.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("RSA private key is of the wrong type")
	}

	privPemBytes := privPem.Bytes

	privateKey, err := x509.ParsePKCS1PrivateKey(privPemBytes)
	if err != nil {
		return nil, err
	}

	pubPem, _ := pem.Decode([]byte(public))

	if pubPem.Type != "PUBLIC KEY" {
		return nil, errors.New("public key is of the wrong type")
	}

	pubPemBytes := pubPem.Bytes

	parsedKey, err := x509.ParsePKIXPublicKey(pubPemBytes)
	if err != nil {
		return nil, err
	}

	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key")
	}

	privateKey.PublicKey = *pubKey

	return privateKey, nil
}
