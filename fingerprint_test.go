package funcaptcha

import (
	"testing"
)

func TestGetFingerprint(t *testing.T) {
	fp := getFingerprint()
	println(fp)
}
