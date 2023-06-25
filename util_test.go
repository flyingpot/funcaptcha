package funcaptcha

import "testing"

func TestGetBda(t *testing.T) {
	res := GetBda("test-useragent", "test-referer", "test-location")
	println(res)
}
