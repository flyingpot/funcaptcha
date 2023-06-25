package funcaptcha

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

type EncryptionData struct {
	Ct string `json:"ct"`
	Iv string `json:"iv"`
	S  string `json:"s"`
}

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func encrypt(data string, key string) string {
	salt := generateSalt(8)
	paddedData := pkcs7Pad([]byte(data))

	md5Hash := md5.New()
	salted := ""
	var dx []byte

	for i := 0; i < 3; i++ {
		md5Hash.Write(dx)
		md5Hash.Write([]byte(key))
		md5Hash.Write([]byte(salt))

		dx = md5Hash.Sum(nil)
		md5Hash.Reset()

		salted += hex.EncodeToString(dx)
	}

	cipherBlock, err := aes.NewCipher([]byte(salted[:32]))
	if err != nil {
		panic(err)
	}

	iv := []byte(salted[32 : 32+16])

	mode := cipher.NewCBCEncrypter(cipherBlock, iv)
	mode.CryptBlocks(paddedData, paddedData)

	encData := EncryptionData{
		Ct: base64.StdEncoding.EncodeToString(paddedData),
		Iv: salted[64 : 64+32],
		S:  hex.EncodeToString([]byte(salt)),
	}

	encDataJson, err := json.Marshal(encData)
	if err != nil {
		panic(err)
	}

	return string(encDataJson)
}

func hexDecode(s string) []byte {
	data, _ := hex.DecodeString(s)
	return data
}

func generateSalt(length int) string {
	bs := make([]byte, length)
	rand.Read(bs)

	for i, b := range bs {
		bs[i] = alphabet[b%byte(len(alphabet))]
	}
	return string(bs)
}

func pkcs7Pad(data []byte) []byte {
	blockSize := aes.BlockSize
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
