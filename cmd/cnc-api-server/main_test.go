package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ExpectMethod(t *testing.T) {
	v := func(http.ResponseWriter, *http.Request) error { return nil }
	resp := httptest.NewRecorder()
	h := WithExpectedMethod(http.MethodPost, v)
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req := httptest.NewRequest(http.MethodPut, "/activate/", nil)
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req = httptest.NewRequest(http.MethodGet, "/activate/", nil)
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req = httptest.NewRequest(http.MethodPost, "/activate/", nil)
	if err := h(resp, req); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_ExpectAuth(t *testing.T) {
	v := func(http.ResponseWriter, *http.Request) error { return nil }
	hdr := http.Header{
		"x-avbl-auth": []string{"test"},
	}
	resp := httptest.NewRecorder()
	h := WithExpectedHeaders(hdr, v)
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req := httptest.NewRequest(http.MethodPut, "/activate/", nil)
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req.Header.Add("x-avbl-auth", "nya")
	if err := h(resp, nil); err == nil {
		t.Error("expected error")
	}

	req.Header.Set("x-avbl-auth", "test")
	if err := h(resp, req); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
