package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	message := strings.Repeat("Hello Encrypt", 1000)
	keyString := "BLCI27HrCTIWc5/QxeRK5n59ygE2bSbulofJ1DP0fGw="
	keyString, _ = LoadAESKey("key/aes-key.pem")
	symCrypto(keyString, message)
}

func symCrypto(keyBase64 string, message string) {
	if keyBase64 == "" {
		keyBase64 = AESKeyGen()
		fmt.Printf("key to encrypt/DecryptAES : %s\n", keyBase64)

	}
	key, _ := base64.StdEncoding.DecodeString(keyBase64)

	fmt.Printf("msg to be encrypted %s\n", message)

	ciphertext := EncryptAES(key, []byte(message))

	fmt.Printf("ciphertext %x\n", ciphertext)

	recoveredMsg := DecryptAES(key, ciphertext)
	if string(recoveredMsg) != message {
		panic("Recovered message != original message")
	} else {
		fmt.Println("Successfully completed encryption/decryption")
	}
}

func AESKeyGen() string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	//encode key in bytes to string and keep as secret, put in a vault
	return base64.StdEncoding.EncodeToString(bytes)
}

func LoadAESKey(keyFile string) (string, error) {
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
		return base64.StdEncoding.EncodeToString(data.Bytes), nil
	}
	panic(fmt.Sprintf("loading key failed: \"%s\"", err))
}

func EncryptAES(key []byte, message []byte) []byte {
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, message, nil)
	return ciphertext
}

func DecryptAES(key []byte, ciphertext []byte) []byte {

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}
