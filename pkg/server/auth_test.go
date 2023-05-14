package server

import "testing"

func Test_AuthHeader_NoSecret(t *testing.T) {
	hdr := GetAuthHeader("")
	if len(hdr) > 0 {
		t.Errorf("invalid auth header: %v", hdr)
	}
}

func Test_AuthHeader_WithSecret(t *testing.T) {
	hdr := GetAuthHeader("test")
	if len(hdr) <= 0 {
		t.Errorf("invalid auth header: %v", hdr)
	}

	if hdr.Get(AuthHeader) != "test" {
		t.Errorf("invalid auth header: %v", hdr)
	}
}
