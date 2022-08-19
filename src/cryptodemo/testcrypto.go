package main

import (
	"cryptodemo/mylib"
	"encoding/base64"
	"os"
)

func main() {
	testMsg := os.Args[1]
	var keyFile string
	if len(os.Args) > 2 {
		keyFile = os.Args[2]
	} else {
		//keyFile = privKeyPath
		keyFile = mylib.PubKeyPath
		//keyFile = aesKey
	}
	if _, error := os.Stat(keyFile); error != nil {
		keyString := mylib.AESKeyGen()
		key, _ := base64.StdEncoding.DecodeString(keyString)
		mylib.SaveAESKey(key, keyFile)
	}
	//mylib.AsymCrypto(keyFile, testMsg)
	keyString, _ := mylib.LoadAESKey(keyFile)
	mylib.SymCrypto(keyString, testMsg)
}
