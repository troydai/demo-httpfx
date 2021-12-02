package types

import "net/http"

type HttpHandler func(w http.ResponseWriter, req *http.Request)

type DataRetriver interface {
	Get(key []byte) string
}
