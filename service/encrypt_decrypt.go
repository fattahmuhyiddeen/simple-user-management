package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	b64 "encoding/base64"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"
)

func Encrypt(message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher([]byte(config.AESKey))
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func Decrypt(securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(config.AESKey))
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}

func URLFriendlyEncrypt(data string) string {
	data, _ = Encrypt(data)
	return b64.StdEncoding.EncodeToString([]byte(data))
}

func URLFriendlyDecrypt(data string) (decryptedString string, err error) {
	decodedByte, err := b64.StdEncoding.DecodeString(data)

	if err == nil {
		decryptedString, err = Decrypt(string(decodedByte))
	}
	return
}
