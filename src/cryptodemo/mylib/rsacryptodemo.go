package mylib

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const PrivKeyPath = "key/private-key.pem"
const PubKeyPath = "key/public-key.pem"
const OutputFile = "output/ciphertext"

func test() {
	testMsg := os.Args[1]
	var keyFile string
	if len(os.Args) > 2 {
		keyFile = os.Args[2]
	} else {
		//keyFile = PrivKeyPath
		keyFile = PubKeyPath
		//keyFile = aesKey
	}

	AsymCrypto(keyFile, testMsg)
	keyString, _ := LoadAESKey(AesKeyFile)
	SymCrypto(keyString, testMsg)
}

func AsymCrypto(keyFile string, testMsg string) {
	if _, err := os.Stat(PrivKeyPath); errors.Is(err, os.ErrNotExist) {
		keyGen(2048)
	}
	privateKey, publicKey, _ := loadKey(keyFile)

	ciphertext, _ := encrypt([]byte(testMsg), *publicKey)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %s\n", encoded)
	fileDump(ciphertext, OutputFile)

	if privateKey != nil {
		recoveredMsg := string(decrypt(privateKey, ciphertext))

		// We get back the original information in the form of bytes, which we
		// the cast to a string and print
		fmt.Println("decrypted message: ", recoveredMsg)
	}
}

func fileDump(ciphertext []byte, filePath string) {
	fileDir := filepath.Dir(filePath)
	if _, err := os.Stat(fileDir); err != nil {
		os.MkdirAll(fileDir, os.ModePerm)
	}
	// Open a new file for writing only
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write bytes to file
	bytesWritten, err := file.Write(ciphertext)
	if err != nil {
		panic(err)
	}
	log.Printf("Wrote %d bytes to %s.\n", bytesWritten, filePath)
}

func decrypt(privateKey *rsa.PrivateKey, ciphertext []byte) []byte {
	// The first argument is an optional random data generator (the rand.Reader we used before)
	// we can set this value as nil
	// The OAEPOptions in the end signify that we encrypted the data using OAEP, and that we used
	// SHA256 to hash the input.
	decryptedBytes, err := rsa.DecryptPKCS1v15(nil, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}
	return decryptedBytes
}

func encrypt(secretMessage []byte, publicKey rsa.PublicKey) ([]byte, error) {
	// cryptodemo/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	ciphertext, err := rsa.EncryptPKCS1v15(rng, &publicKey, secretMessage)
	if err != nil {
		panic(fmt.Sprintf("Error from encryption: %s\n", err))
	}
	return ciphertext, nil
}

func loadKey(keyFile string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKeyFile, err := os.Open(keyFile)
	if err != nil {
		panic(err)
	}
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}
	data, _ := pem.Decode(pembytes)
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err == nil {
		return privateKeyImported, &privateKeyImported.PublicKey, nil
	}
	publicKey, err2 := x509.ParsePKCS1PublicKey(data.Bytes)
	if err2 == nil {
		return nil, publicKey, nil
	}
	panic(fmt.Sprintf("loading key failed for both private key: \"%s\" and public: \"%s\"", err,
		err2))
}

func keyGen(size int) {
	if stat, _ := os.Stat(PrivKeyPath); stat != nil {
		panic("private key file already exists")
	}

	keyDir := filepath.Dir(PrivKeyPath)
	if _, err := os.Stat(keyDir); err != nil {
		os.MkdirAll(keyDir, os.ModePerm)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		panic(err)
	}
	pemPrivateFile, err := os.Create(PrivKeyPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemPrivateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPrivateFile.Close()
}
