package funcaptcha

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	enc := encrypt("test-data", "test-key")
	println(enc)
}
