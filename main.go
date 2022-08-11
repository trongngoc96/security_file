package main

import (
	"flag"
	"fmt"
	"security_file/blockchains"
	"security_file/models"
	"security_file/security"
	"strings"
)

func main() {
	var data models.User
	pathFile := flag.String("file", "", "")
	key := flag.String("key", "", "")
	feature := flag.String("feature", "", "")
	findKey := flag.String("findkey", "", "")
	flag.Parse()
	arrFile := strings.Split(string(*pathFile), ",")

	switch *feature {
	case "encrypt":
		security.Encrypt(arrFile, string(*key))
	case "decrypt":
		dataFile := security.Decrypt(arrFile, string(*key), string(*findKey))
		fmt.Println(dataFile)
		data.PASSWORD_WALLET = dataFile[0].VALUE
		blockchains.CreateKs(&data)
	}
}
