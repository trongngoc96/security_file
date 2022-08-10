package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"security_file/file"
	"strings"
)

type DataFile struct {
	KEY   string `json:"key"`
	PATH  string `json:"path"`
	VALUE string `json:"value"`
}

func Decrypt(arrFile []string, keystring string, findkey string) []DataFile {
	var arrDataValue []DataFile
	var dataFile DataFile
	for i := 0; i < len(arrFile); i++ {
		if ciphertext, err := file.ReadFromFile(arrFile[i]); err != nil {
			fmt.Println("File is not found")
		} else {
			ciphertext := []byte(string(ciphertext))

			// Key
			key := []byte(keystring)

			// Create the AES cipher
			block, err := aes.NewCipher(key)
			if err != nil {
				panic(err)
			}

			// Before even testing the decryption,
			// if the text is too small, then it is incorrect
			if len(ciphertext) < aes.BlockSize {
				panic("Text is too short")
			}

			// Get the 16 byte IV
			iv := ciphertext[:aes.BlockSize]

			// Remove the IV from the ciphertext
			ciphertext = ciphertext[aes.BlockSize:]

			// Return a decrypted stream
			stream := cipher.NewCFBDecrypter(block, iv)

			// Decrypt bytes from ciphertext
			stream.XORKeyStream(ciphertext, ciphertext)

			removeDistance := strings.ReplaceAll(string(ciphertext), " ", "")
			removeLine := strings.ReplaceAll(removeDistance, "\n", "")
			indexFindKey := strings.Index(removeLine, findkey)
			removeBeforeFindKey := removeLine[indexFindKey+len(findkey)+1 : len(removeLine)-1]
			indexValueByFindKey := strings.Index(removeBeforeFindKey, ",")
			valueByFindKey := removeBeforeFindKey[:indexValueByFindKey]
			dataFile.PATH = arrFile[i]
			dataFile.KEY = findkey
			dataFile.VALUE = valueByFindKey
			arrDataValue = append(arrDataValue, dataFile)
			//fmt.Println(valueByFindKey)
		}
	}
	// Byte array of the string
	return arrDataValue
}
