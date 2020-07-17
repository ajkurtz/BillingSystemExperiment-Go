package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

const key = "12345678901234567890123456789012"

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Encrypt(plaintext string) string {
	passphrase := []byte(key)

	block, err := aes.NewCipher(passphrase)
	CheckError(err)

	text := []byte(plaintext)

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	CheckError(err)

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(ciphertext string) string {
	passphrase := []byte(key)

	block, err := aes.NewCipher(passphrase)
	CheckError(err)

	text, err := base64.StdEncoding.DecodeString(ciphertext)
	CheckError(err)

	if len(text) < aes.BlockSize {
		log.Fatal("Ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	plaintext, err := base64.StdEncoding.DecodeString(string(text))
	CheckError(err)

	return string(plaintext)
}
