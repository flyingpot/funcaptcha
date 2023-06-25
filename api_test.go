package funcaptcha

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	options := &GetTokenOptions{
		PKey: "35536E1E-65B4-4D96-9D97-6ADB7EFF8147",
		SURL: "https://tcr9i.chat.openai.com",
	}
	res, _ := GetToken(options)
	println(res.Token)
}
