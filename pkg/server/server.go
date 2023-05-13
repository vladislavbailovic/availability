package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	http.ResponseWriter
	Done bool
}

func (x *Response) Write(v []byte) (int, error) {
	x.Done = true
	return x.ResponseWriter.Write(v)
}

func (x *Response) WriteHeader(c int) {
	if !x.Done {
		x.ResponseWriter.WriteHeader(c)
	}
}

type Handler func(*Response, *http.Request) error

func Handle(f Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &Response{ResponseWriter: w}
		if err := f(resp, r); err != nil {
			log.Printf("ERROR [%s %s %s]: %v",
				r.RemoteAddr, r.Method, r.URL.Path, err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("DEBUG: [%s %s %s] OK",
			r.RemoteAddr, r.Method, r.URL.Path)
		resp.WriteHeader(http.StatusOK)
	}
}

func ExtractNumberFromPathAt(r *http.Request, at int) int {
	splits := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(splits) <= at {
		return 0
	}
	id, err := strconv.Atoi(splits[at])
	if err != nil {
		return 0
	}
	return id
}
