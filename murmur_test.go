package funcaptcha

import (
	"testing"
)

func TestGetMurmur128String(t *testing.T) {
	if GetMurmur128String("test", 31) != "ff55565a476832ed3409c64597508ca4" {
		t.Fatal("murmur error!")
	}
}
