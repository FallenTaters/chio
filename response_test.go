package chio_test

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FallenTaters/chio"
	"github.com/FallenTaters/chio/internal/assert"
)

func TestEmpty(t *testing.T) {
	rec := httptest.NewRecorder()

	chio.Empty(rec, http.StatusOK)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, rec.Body.Len(), 0)
}

func TestWriteString(t *testing.T) {
	rec := httptest.NewRecorder()

	chio.WriteString(rec, http.StatusOK, `bla`)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.DeepEqual(t, rec.Header().Get(`Content-Type`), `text/plain`)
	assert.Equal(t, rec.Body.String(), `bla`)
}

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()

	jsonStruct := struct {
		A int `json:"a"`
	}{}

	chio.WriteJSON(rec, http.StatusOK, jsonStruct)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.DeepEqual(t, rec.Header().Get(`Content-Type`), `application/json`)
	assert.Equal(t, rec.Body.String(), "{\"a\":0}\n")
}

func TestWriteXML(t *testing.T) {
	rec := httptest.NewRecorder()

	xmlStruct := struct {
		XMLName xml.Name `xml:"a"`
		A       int      `xml:"a"`
	}{}

	chio.WriteXML(rec, http.StatusOK, xmlStruct)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.DeepEqual(t, rec.Header().Get(`Content-Type`), `application/xml`)
	assert.Equal(t, rec.Body.String(), "<a><a>0</a></a>")
}

func TestStreamBlob(t *testing.T) {
	rec := httptest.NewRecorder()

	expected := `bla`

	chio.StreamBlob(rec, http.StatusOK, strings.NewReader(expected))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.DeepEqual(t, rec.Header().Get(`Content-Type`), `application/octet-stream`)
	assert.Equal(t, rec.Body.String(), "bla")
}

func TestWriteBlob(t *testing.T) {
	rec := httptest.NewRecorder()

	expected := `bla`

	chio.WriteBlob(rec, http.StatusOK, []byte(expected))
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.DeepEqual(t, rec.Header().Get(`Content-Type`), `application/octet-stream`)
	assert.Equal(t, rec.Body.String(), "bla")
}
