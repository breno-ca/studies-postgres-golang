package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type writer struct {
	http.ResponseWriter
	status int
}

func (w *writer) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

const (
	// Symbolic constants
	bar string = "|"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &writer{ResponseWriter: w}

		next.ServeHTTP(writer, r)

		log.Println(bar, fit(r.Method), bar, writer.status, bar, r.URL)
	})
}

func fit(value string) string {
	return fmt.Sprintf("%-7s", value)
}
