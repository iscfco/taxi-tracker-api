package jwttasks

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type jwtTasks struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var jwtTasksInstance *jwtTasks = nil

func NewJwtTasks() (error, *jwtTasks) {
	err, privKey := getPrivateKey()
	if err != nil {
		return err, nil
	}

	err, pubKey := getPublicKey()
	if err != nil {
		return err, nil
	}
	if jwtTasksInstance == nil {
		jwtTasksInstance = &jwtTasks{
			privateKey: privKey,
			PublicKey:  pubKey,
		}
		fmt.Println(" ----- jwtTasks instance created ----- ")
	}

	return nil, jwtTasksInstance
}

func (jwtTasks *jwtTasks) GenerateAccessToken(userId *string) (tokenStr string, err error){
	return jwtTasks.generateToken(userId, TypeAccessToken, accessTokenDuration)
}

func (jwtTasks *jwtTasks) GenerateRefreshToken(userId *string) (tokenStr string, err error){
	return jwtTasks.generateToken(userId, TypeRefreshToken, refreshTokenDuration)
}

func (jwtTasks *jwtTasks) generateToken(usrId *string, useType string, durationInMin int) (tknStr string, err error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(durationInMin)).Unix(),
		"iat": time.Now().Unix(),
		"sub": usrId,
		"use": useType,
	}
	tknStr, err = token.SignedString(jwtTasks.privateKey)
	if err != nil {
		return tknStr, err
	}
	return tknStr, err
}

func getPrivateKey() (error, *rsa.PrivateKey) {
	privateKeyFile, err := os.Open(getJwtPrivateKeyPath())
	if err != nil {
		return err, nil
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	return nil, privateKeyImported
}

func getPublicKey() (error, *rsa.PublicKey) {
	publicKeyFile, err := os.Open(getJwtPublicKeyPath())
	if err != nil {
		return err, nil
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	publicKeyFile.Close()
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}

	return nil, rsaPub
}
