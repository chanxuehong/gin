package response

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func newResponseWriter00011() ResponseWriter2 { return new(responseWriter00011) }

var _ http.Hijacker = (*responseWriter00011)(nil)
var _ http.CloseNotifier = (*responseWriter00011)(nil)

type responseWriter00011 struct {
	responseWriter00001
	hijacked bool
}

func (w *responseWriter00011) Reset(writer http.ResponseWriter) {
	w.responseWriter00001.Reset(writer)
	w.hijacked = false
}

func (w *responseWriter00011) WriteHeader(code int) {
	if w.hijacked {
		log.Println("gin: response.WriteHeader on hijacked connection")
		return
	}
	w.responseWriter00001.WriteHeader(code)
}

func (w *responseWriter00011) Write(data []byte) (n int, err error) {
	if w.hijacked {
		log.Println("gin: response.Write on hijacked connection")
		return 0, http.ErrHijacked
	}
	return w.responseWriter00001.Write(data)
}

func (w *responseWriter00011) Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
	rwc, buf, err = w.responseWriter.(http.Hijacker).Hijack()
	if err == nil {
		w.hijacked = true
	}
	return
}
