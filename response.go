package chio

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Empty writes the status code and sets the Content-Length header to 0
func Empty(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Length`, `0`)
}

// WriteString writes s to the body and sets the Content-Type to text/plain
func WriteString(w http.ResponseWriter, code int, s string) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Type`, `text/plain`)
	if _, err := w.Write([]byte(s)); err != nil {
		panic(err)
	}
}

// WriteJSON writes v to the body as WriteJSON and sets the Content-Type to application/json. If marshalling fails, it panics.
func WriteJSON(w http.ResponseWriter, code int, v any) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Type`, `application/json`)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(fmt.Errorf(`error encoding JSON into response: %w`, err))
	}
}

// WriteXML writes v to the body as WriteXML and sets the Content-Type to application/xml. If marshalling fails, it panics.
func WriteXML(w http.ResponseWriter, code int, v any) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Type`, `application/xml`)
	if err := xml.NewEncoder(w).Encode(v); err != nil {
		panic(fmt.Errorf(`error encoding XML into response: %w`, err))
	}
}

// StreamBlob copies data from r to w
func StreamBlob(w http.ResponseWriter, code int, contentType string, r io.Reader) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Type`, contentType)
	if _, err := io.Copy(w, r); err != nil {
		panic(fmt.Errorf(`error streaming blob into response: %w`, err))
	}
}

// WriteBlob writes v to the body and automatically detects the Content-Type
func WriteBlob(w http.ResponseWriter, code int, v []byte) {
	w.WriteHeader(code)
	w.Header().Set(`Content-Type`, http.DetectContentType(v))
	if _, err := w.Write(v); err != nil {
		panic(fmt.Errorf(`error writing blob into response: %w`, err))
	}
}
