package chio_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FallenTaters/chio"
	"github.com/FallenTaters/chio/internal/assert"
)

func TestJSON(t *testing.T) {
	h := chio.JSON(func(w http.ResponseWriter, r *http.Request, i int) {
		assert.Equal(t, i, 3)
		chio.WriteString(w, http.StatusOK, `success`)
	})

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, ``, strings.NewReader(``))
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusUnprocessableEntity)
	assert.Equal(t, rec.Body.String(), ``)

	rec = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, ``, strings.NewReader(`3`))
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, rec.Body.String(), `success`)
}
