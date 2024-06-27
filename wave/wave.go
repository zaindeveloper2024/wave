package wave

import (
	"net/http"
)

type Wave struct {
	server *http.ServeMux
}

func New() (w *Wave) {
	w = &Wave{
		server: newServer(),
	}
	return
}

func newServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func (w *Wave) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	w.server.HandleFunc(pattern, handler)
}
func (w *Wave) Start(address string) error {
	return http.ListenAndServe(address, w.server)
}
