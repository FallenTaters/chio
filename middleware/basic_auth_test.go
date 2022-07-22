package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FallenTaters/chio/internal/assert"
	"github.com/FallenTaters/chio/middleware"
)

func TestBasicAuth(t *testing.T) {
	value, ok := ``, false
	var username, password string
	mw := middleware.BasicAuth(`basicAuthKey`, func(u, p string) (string, bool) {
		username, password = u, p
		return value, ok
	})
	called, keyVal := false, ``
	h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		keyVal = middleware.GetValue[string](r, `basicAuthKey`)
	}))

	// no basic auth
	req, err := http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(httptest.NewRecorder(), req)
	assert.Equal(t, username, ``)
	assert.Equal(t, password, ``)
	assert.False(t, called)
	assert.Equal(t, keyVal, ``)

	// auth failed
	username, password = ``, ``
	called, keyVal = false, ``
	req, err = http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(`myUser`, `myPass`)
	h.ServeHTTP(httptest.NewRecorder(), req)
	assert.Equal(t, username, `myUser`)
	assert.Equal(t, password, `myPass`)
	assert.False(t, called)
	assert.Equal(t, keyVal, ``)

	// auth successful
	username, password = ``, ``
	called = false
	value, ok = `value`, true
	req, err = http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(`myUser`, `myPass`)
	h.ServeHTTP(httptest.NewRecorder(), req)
	assert.Equal(t, username, `myUser`)
	assert.Equal(t, password, `myPass`)
	assert.True(t, called)
	assert.Equal(t, keyVal, `value`)
}
