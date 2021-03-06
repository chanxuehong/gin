package response

import (
	"io"
	"net/http"
)

func newResponseWriter10001() ResponseWriter2 { return new(responseWriter10001) }

var _ io.ReaderFrom = (*responseWriter10001)(nil)
var _ http.CloseNotifier = (*responseWriter10001)(nil)

type responseWriter10001 struct {
	responseWriter00001
}

func (w *responseWriter10001) ReadFrom(r io.Reader) (n int64, err error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	n, err = w.responseWriter.(io.ReaderFrom).ReadFrom(r)
	w.written += n
	return
}
