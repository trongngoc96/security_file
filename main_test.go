package main

import (
	"flag"
	"security_file/blockchains"
	"security_file/models"
	"security_file/security"
	"strings"
	"testing"
)

var pathFile = flag.String("file", "", "")
var key = flag.String("key", "", "")
var findKey = flag.String("findkey", "", "")

func BenchmarkSieveOfErastosthenes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var data models.User
		flag.Parse()
		arrFile := strings.Split(string(*pathFile), ",")
		dataFile := security.Decrypt(arrFile, string(*key), string(*findKey))
		data.PASSWORD_WALLET = dataFile[0].VALUE
		blockchains.CreateKs(&data)

	}
}
