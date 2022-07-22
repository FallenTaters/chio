// Package assert contains test helpers
package assert

import (
	"reflect"
	"testing"
)

// Equal asserts equality
func Equal[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		t.Errorf(`%#v doesn't equal %#v`, a, b)
	}
}

// NotEqual asserts non-equality
func NotEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf(`%#v doesn't equal %#v`, a, b)
	}
}

// DeepEqual asserts equality
func DeepEqual[T any](t *testing.T, a, b T) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf(`%#v doesn't deep-equal %#v`, a, b)
	}
}

// True asserts that v is true
func True(t *testing.T, v bool) {
	t.Helper()
	if !v {
		t.Errorf(`should be true but isn't`)
	}
}

// False asserts that v is false
func False(t *testing.T, v bool) {
	t.Helper()
	if v {
		t.Errorf(`should be false but isn't`)
	}
}
