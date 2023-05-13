package server

import (
	"errors"
	"fmt"
	"net/http"
)

func WithExpectedMethod(method string, f Handler) Handler {
	return func(w *Response, r *http.Request) error {
		if r == nil {
			return errors.New("invalid request")
		}
		if r.Method != method {
			return fmt.Errorf("unsupported request type: %q, expected %q",
				r.Method, method)
		}
		return f(w, r)
	}
}

func WithExpectedHeaders(hdr http.Header, f Handler) Handler {
	return func(w *Response, r *http.Request) error {
		if r == nil {
			return errors.New("invalid request")
		}
		for key, wants := range hdr {
			gots := r.Header.Values(key)
			if len(gots) == 0 {
				return fmt.Errorf("missing required header %q", key)
			}
			for idx, want := range wants {
				got := gots[idx]
				if want != got {
					return fmt.Errorf("invalid header %q(%d): %q", key, idx, got)
				}
			}
		}
		return f(w, r)
	}
}
