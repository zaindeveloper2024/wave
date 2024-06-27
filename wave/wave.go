package wave

import (
	"net/http"
)

type Wave struct {
	handler http.Handler
}

func New() (w *Wave) {
	w = &Wave{
		handler: newServer(),
	}
	return
}

func newServer() http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux
	return handler
}

func (w *Wave) Start(address string) error {
	return http.ListenAndServe(address, w.handler)
}
