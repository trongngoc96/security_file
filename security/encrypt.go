package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"security_file/file"
)

func Encrypt(plainstring []string, keystring string) {
	for i := 0; i < len(plainstring); i++ {
		plaintext, err := ioutil.ReadFile(plainstring[i])
		if err != nil {
			panic(err.Error())
		}
		// Key
		key := []byte(keystring)
		// Create the AES cipher
		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}
		// Empty array of 16 + plaintext length
		// Include the IV at the beginning
		ciphertext := make([]byte, aes.BlockSize+len(plaintext))
		// Slice of first 16 bytes
		iv := ciphertext[:aes.BlockSize]
		// Write 16 rand bytes to fill iv
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			panic(err)
		}
		// Return an encrypted stream
		stream := cipher.NewCFBEncrypter(block, iv)
		// Encrypt bytes from plaintext to ciphertext
		stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
		file.WriteToFile(string(ciphertext), plainstring[i])
	}

}
