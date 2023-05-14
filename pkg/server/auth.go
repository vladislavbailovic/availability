package server

import "net/http"

const AuthHeader string = "x-avbl-auth"

func GetAuthHeader(secret string) http.Header {
	hdr := http.Header{}
	if secret != "" {
		hdr.Add(AuthHeader, secret)
	}
	return hdr
}
