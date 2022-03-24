package controllers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dryairship/techmeet22-nhid-service/config"
)

var (
	RSAPublicKey       *rsa.PublicKey
	RSAPrivateKey      *rsa.PrivateKey
	RSAPublicKeyString string
)

func init() {
	privateKeyBytes, err := ioutil.ReadFile(config.RSAPrivateKeyFile)
	if err != nil {
		log.Fatal("[ERROR] Could not read private key file: ", err.Error())
	}

	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	RSAPrivateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Fatal("[ERROR] Could not parse private key: ", err.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile(config.RSAPublicKeyFile)
	if err != nil {
		log.Fatal("[ERROR] Could not read public key file: ", err.Error())
	}

	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	RSAPublicKey, err = x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Fatal("[ERROR] Could not parse public key: ", err.Error())
	}

	RSAPublicKeyString = string(publicKeyBytes)
}

func GetPublicKey(c *gin.Context) {
	c.String(http.StatusOK, RSAPublicKeyString)
}
