package funcaptcha

import (
	"testing"
)

func TestGetMurmur128String(t *testing.T) {
	if "ff55565a476832ed3409c64597508ca4" != GetMurmur128String("test", 31) {
		t.Fatal("murmur error!")
	}
}
