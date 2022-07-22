# chio

chio contains useful functions to use in conjunction with go-chi/chi router.

## decode helpers

example

```go
import (
    "github.com/go-chi/chi/v5"
    "github.com/FallenTaters/chio"
)

func main() {
    r := chi.NewRouter()
    r.Get(`/`, chio.JSON(handler))
    // ...
}

type myStruct struct{
    A int `json:"a"`
}

func handler(w http.ResponseWriter, r *http.Request, v myStruct) {
    // ...
}
```

### JSON

if the body cannot be read and decoded, http.StatusUnprocessableEntity is returned.

## encode helpers

example

```go
import (
    "github.com/go-chi/chi/v5"
    "github.com/FallenTaters/chio"
)

func main() {
    r := chi.NewRouter()
    r.Get(`/`, handler)
    // ...
}

func handler(w http.ResponseWriter, r *http.Request) {
    chio.WriteJSON(w, http.StatusOK, myStruct{3})
}
```

### Empty

writes the status code and sets the Content-Length header to 0

### WriteString

writes s to the body and sets the Content-Type to text/plain

### WriteJSON

writes v to the body as WriteJSON and sets the Content-Type to application/json. If marshalling fails, it panics.

### WriteXML

writes v to the body as WriteXML and sets the Content-Type to application/xml. If marshalling fails, it panics.

### StreamBlob

copies data from r to w and sets the Content-Type to application/octet-stream

### WriteBlob

writes v to the body and sets the Content-Type to application/octet-stream

## Middleware

example

```go
import (
    "github.com/go-chi/chi/v5"
    "github.com/FallenTaters/chio/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Recover(middleware.DefaultPanicLogger(os.Stderr)))
    // ...
}

func handler(w http.ResponseWriter, r *http.Request) {
    chio.WriteJSON(w, http.StatusOK, myStruct{3})
}
```

### BasicAuth

A simple middleware handler for adding basic http auth to a route.
Authorize should return true if the username and password are correct, otherwise 401 will be returned.
The returned contextKey and contextValue can be used to set values on the request context, for example a user struct.
Values can be retrieved easily using middleware.GetValue.

### Recover

Catches a panic and passes the value and a stack trace to f.

### SetValue

Sets a value on the request context for use in later middlewares or handlers.
If get returns false, it is assumed something went wrong and the correct response has already been written.
If get returns true, v will be set under k in the request context and the underlying handler is called.
The value can be retrieved and type-asserted easily using GetValue.

