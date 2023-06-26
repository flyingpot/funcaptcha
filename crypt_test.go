package funcaptcha

import "testing"

func TestEncryptDecrypt1(t *testing.T) {
	enc := Encrypt("test-data", "test-key")
	println(enc)
}
