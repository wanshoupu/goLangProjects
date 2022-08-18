package main

import (
	"bufio"
	"crypto/aes"
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

const privKeyPath = "key/private-key.pem"
const pubKeyPath = "key/public-key.pem"
const aesKey = "key/public-key.pem"
const outputFile = "output/ciphertext"

func main() {
	testMsg := os.Args[1]
	var keyFile string
	if len(os.Args) > 2 {
		keyFile = os.Args[2]
	} else {
		//keyFile := privKeyPath
		//keyFile = pubKeyPath
		keyFile = aesKey
	}

	//asymCrypto(keyFile, testMsg)
	ciphertext := symCrypto(keyFile, testMsg)
	fmt.Println(ciphertext)
}

func symCrypto(keyFile string, msg string) string {
	if _, err := os.Stat(keyFile); errors.Is(err, os.ErrNotExist) {
		AESKeyGen(keyFile, 32)
	}
	key, err := loadAESKey(keyFile)
	if err != nil {
		panic("Missing Key")
	}
	return EncryptAES(key, msg)
}

func AESKeyGen(keyFile string, length int) []byte {
	if stat, _ := os.Stat(keyFile); stat != nil {
		panic("private key file already exists")
	}

	keyDir := filepath.Dir(keyFile)
	if _, err := os.Stat(keyDir); err != nil {
		os.MkdirAll(keyDir, os.ModePerm)
	}
	pemFile, err := os.Create(keyFile)
	if err != nil {
		panic(err)
	}

	key := make([]byte, length)

	// Reads {length} cryptographically secure pseudorandom numbers from rand.Reader and writes them to key
	_, err = rand.Read(key)

	pemPrivateBlock := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}
	err = pem.Encode(pemFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemFile.Close()
	return key
}

func loadAESKey(keyFile string) ([]byte, error) {
	kf, err := os.Open(keyFile)
	if err != nil {
		panic(err)
	}
	pemfileinfo, _ := kf.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(kf)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}
	data, _ := pem.Decode(pembytes)
	kf.Close()
	if err == nil {
		return data.Bytes, nil
	}
	panic(fmt.Sprintf("loading key failed: \"%s\"", err))
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return base64.StdEncoding.EncodeToString(out)
}

func asymCrypto(keyFile string, testMsg string) {
	if _, err := os.Stat(privKeyPath); errors.Is(err, os.ErrNotExist) {
		keyGen(2048)
	}
	privateKey, publicKey, _ := loadKey(keyFile)

	ciphertext, _ := encrypt([]byte(testMsg), *publicKey)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %s\n", encoded)
	fileDump(ciphertext, outputFile)

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
	// crypto/rand.Reader is a good source of entropy for randomizing the
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
	if stat, _ := os.Stat(privKeyPath); stat != nil {
		panic("private key file already exists")
	}

	keyDir := filepath.Dir(privKeyPath)
	if _, err := os.Stat(keyDir); err != nil {
		os.MkdirAll(keyDir, os.ModePerm)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		panic(err)
	}
	pemPrivateFile, err := os.Create(privKeyPath)
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
