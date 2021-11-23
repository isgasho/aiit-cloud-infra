package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
)

type Keys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (k *Keys) createPrivateKey() (err error) {
	k.PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	return nil
}

func (k *Keys) createPublicKey() error {
	var ok bool
	k.PublicKey, ok = k.PrivateKey.Public().(*rsa.PublicKey)
	if !ok {
		return errors.New("create public key error")
	}
	return nil
}

func NewKeys() (*Keys, error) {
	keys := &Keys{}
	if err := keys.createPrivateKey(); err != nil {
		return nil, err
	}
	if err := keys.createPublicKey(); err != nil {
		return nil, err
	}
	return keys, nil
}

func (k *Keys) CreatePrivateKeyFile(id int) (string, error) {
	filePath := fmt.Sprintf("/tmp/instance-%v-private.pem", id)
	derRsaPrivateKey := x509.MarshalPKCS1PrivateKey(k.PrivateKey)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("CreatePrivateKeyFile close error")
		}
	}(f)

	if err := pem.Encode(f, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: derRsaPrivateKey}); err != nil {
		return "", err
	}
	return filePath, nil
}

func (k *Keys) CreatePublicKeyFile(id int) (string, string, error) {
	filePath := fmt.Sprintf("/tmp/instance-%v-public.pem", id)
	derRsaPublicKey := x509.MarshalPKCS1PublicKey(k.PublicKey)
	f, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("CreatePublicKeyFile create close error")
		}
	}(f)

	if err := pem.Encode(f, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: derRsaPublicKey}); err != nil {
		return "", "", err
	}

	pemf, err := os.Open(filePath)
	if err != nil {
		return filePath, "", err
	}
	defer func(pemf *os.File) {
		err := pemf.Close()
		if err != nil {
			fmt.Println("CreatePublicKeyFile open close error")
		}
	}(pemf)

	bytes, err := io.ReadAll(pemf)
	if err != nil {
		fmt.Println(err)
		return filePath, "", err
	}

	return filePath, string(bytes), nil
}
