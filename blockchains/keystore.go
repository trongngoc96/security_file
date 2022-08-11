package blockchains

import (
	"log"
	"security_file/models"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func CreateKs(user *models.User) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := user.PASSWORD_WALLET
	key, err := ks.NewAccount(password)
	//fmt.Println("in ra key tuc la account truyen trong enscryp ne:", key)
	if err != nil {
		log.Fatal(err)
	}
	user.ADDRESS = key.Address.Hex()
	byteArray, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		log.Fatal(err)
	}
	strKeyStore := string(byteArray[:])
	// data, err := keystore.DecryptKey(byteArray, password)
	// fmt.Println("in ra key tuc la account truyen trong enscryp ne:", hex.EncodeToString(data.PrivateKey.D.Bytes()))
	byteArrayAdmin, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		log.Fatal(err)
	}
	strKeyStoreAdmin := string(byteArrayAdmin[:])
	user.KEYSTORE = strKeyStore
	user.KEYSTORE_ADMIN = strKeyStoreAdmin
}
