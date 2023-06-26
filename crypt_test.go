package funcaptcha

import "testing"

func TestEncryptDecrypt1(t *testing.T) {
	enc := encrypt("test-data", "test-key")
	println(enc)
}
