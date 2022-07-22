package middleware

import (
	"context"
	"net/http"
)

// SetValue sets a value on the request context for use in later middlewares or handlers.
// If get returns false, it is assumed something went wrong and the correct response has already been written.
// If get returns true, v will be set under k in the request context and the underlying handler is called.
// The value can be retrieved and type-asserted easily using GetValue.
func SetValue[V any](
	key any,
	get func(w http.ResponseWriter, r *http.Request) (V, bool),
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v, ok := get(w, r)
			if !ok {
				return
			}

			ctx := context.WithValue(r.Context(), key, v)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetValue retrieves a value from the request context set by SetValue middleware.
// If the value was not set or is of a different type than V, the zero value for V is returned.
func GetValue[V any](r *http.Request, key any) V {
	v, _ := r.Context().Value(key).(V)
	return v
}
