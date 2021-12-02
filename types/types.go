package types

import "net/http"

type HttpHandler interface {
	Path() string
	Handle(w http.ResponseWriter, req *http.Request)
}

type DataRetriver interface {
	Get(key []byte) string
}
