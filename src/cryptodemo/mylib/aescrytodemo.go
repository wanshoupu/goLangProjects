package mylib

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

const AesKeyFile = "key/aes-key.pem"
const CipherTextFile = "output/ciphertext"

func main() {

	message := strings.Repeat("Hello Encrypt", 2)
	fmt.Printf("msg to be encrypted %s\n", message)
	if _, err := os.Stat(AesKeyFile); err != nil {
		keyString := AESKeyGen()
		key, _ := base64.StdEncoding.DecodeString(keyString)
		SaveAESKey(key, AesKeyFile)
	}

	keyString, _ := LoadAESKey(AesKeyFile)
	SymCrypto(keyString, message)
	//os.Remove(AesKeyFile)
	fmt.Printf("key to encrypt/DecryptAES : %s\n", keyString)
}

func SymCrypto(keyBase64 string, message string) {
	if keyBase64 == "" {
		keyBase64 = AESKeyGen()
	}
	key, _ := base64.StdEncoding.DecodeString(keyBase64)

	ciphertext := EncryptAES(key, []byte(message))
	b64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println(b64)
	os.WriteFile(CipherTextFile, []byte(b64), 0644)
	recoveredMsg := DecryptAES(key, ciphertext)
	if string(recoveredMsg) != message {
		panic("Recovered message != original message")
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

func SaveAESKey(key []byte, keyFile string) error {
	pemPrivateBlock := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}

	pemFile, err := os.Create(keyFile)
	defer pemFile.Close()

	if err != nil {
		panic(err)
	}

	err = pem.Encode(pemFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return err
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
	aesGCM, err := CreateGCMCipher(key)

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case,
	//we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, message, nil)
	return ciphertext
}

func CreateGCMCipher(key []byte) (cipher.AEAD, error) {
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
	return aesGCM, err
}

func DecryptAES(key []byte, ciphertext []byte) []byte {
	aesGCM, err := CreateGCMCipher(key)

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
