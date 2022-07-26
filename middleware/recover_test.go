package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FallenTaters/chio/internal/assert"
	"github.com/FallenTaters/chio/middleware"
)

func TestRecover(t *testing.T) {
	var v any
	var stck []middleware.StackTraceEntry
	var codeInPanic int
	var codeInHandler int
	mw := middleware.Recover(func(w http.ResponseWriter, r *http.Request, panicValue any, stack []middleware.StackTraceEntry) {
		v, stck = panicValue, stack
		if codeInPanic != 0 {
			w.WriteHeader(codeInPanic)
		}
	})
	var pnc any = nil
	h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if codeInHandler != 0 {
			w.WriteHeader(codeInHandler)
		}
		if pnc != nil {
			panic(pnc)
		}
	}))

	// no panic
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(rec, req)
	assert.True(t, v == nil)
	assert.True(t, stck == nil)
	assert.Equal(t, rec.Code, http.StatusOK)

	// do panic
	v, stck, codeInPanic, codeInHandler = nil, nil, 0, 0
	pnc = 0
	rec = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(rec, req)
	assert.True(t, v == 0)
	assert.DeepEqual(t, len(stck), 6)
	assert.Equal(t, stck[0].Function, `github.com/FallenTaters/chio/middleware_test.TestRecover.func2`)
	assert.NotEqual(t, stck[0].Line, 0)
	assert.True(t, strings.HasSuffix(stck[0].File, `/middleware/recover_test.go`))
	assert.Equal(t, rec.Code, http.StatusInternalServerError)

	// http.ErrAbortHandler
	v, stck, codeInPanic, codeInHandler = nil, nil, 0, 0
	pnc = http.ErrAbortHandler
	rec = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	func() {
		defer func() {
			v := recover()
			assert.True(t, v == http.ErrAbortHandler)
		}()
		h.ServeHTTP(rec, req)
	}()
	assert.True(t, v == nil)
	assert.True(t, stck == nil)
	assert.Equal(t, rec.Code, http.StatusOK)

	// overwrite status code
	v, stck, codeInPanic, codeInHandler = nil, nil, http.StatusBadRequest, http.StatusUnauthorized
	pnc = 0
	rec2 := &writeCounter{httptest.NewRecorder(), nil}
	req, err = http.NewRequest(http.MethodGet, ``, nil)
	if err != nil {
		panic(err)
	}
	h.ServeHTTP(rec2, req)
	assert.True(t, v == 0)
	assert.True(t, stck != nil)
	assert.Equal(t, rec2.Code, http.StatusBadRequest)
	assert.DeepEqual(t, rec2.writes, []int{http.StatusBadRequest})

	v, stck, codeInPanic, codeInHandler = nil, nil, 0, http.StatusUnauthorized
	rec2 = &writeCounter{httptest.NewRecorder(), nil}
	h.ServeHTTP(rec2, req)
	assert.DeepEqual(t, rec2.writes, []int{http.StatusInternalServerError})

	v, stck, codeInPanic, codeInHandler = nil, nil, 0, 0
	rec2 = &writeCounter{httptest.NewRecorder(), nil}
	h.ServeHTTP(rec2, req)
	assert.DeepEqual(t, rec2.writes, []int{http.StatusInternalServerError})
}

type writeCounter struct {
	*httptest.ResponseRecorder
	writes []int
}

func (w *writeCounter) WriteHeader(code int) {
	w.writes = append(w.writes, code)
	w.ResponseRecorder.WriteHeader(code)
}

func TestDefaultPanicLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := middleware.DefaultPanicLogger(&buf)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, `/bla`, nil)
	if err != nil {
		panic(err)
	}
	logger(rec, req, 0, []middleware.StackTraceEntry{{
		Function: `myFunc`,
		File:     `myFile`,
		Line:     100,
	}})

	assert.Equal(t, buf.String(), "GET \"/bla\" - panic: 0\n\tmyFunc\n\t\tmyFile:100\n\n")
}
