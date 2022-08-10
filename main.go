package main

import (
	"flag"
	"fmt"
	"security_file/security"
	"strings"
)

func main() {
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
	}
}
