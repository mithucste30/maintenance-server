package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var port = ":80"

type Message struct {
	Content string
}

func main() {
	// env overriding port
	if os.Getenv("SERVER_PORT") != "" {
		port = os.Getenv("SERVER_PORT")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", middleware(fs))

	log.Print(fmt.Sprintf("Listening on %s...", port))
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// everything is down
		w.WriteHeader(http.StatusServiceUnavailable)

		// filter dotted
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			return
		}

		// pass through files
		next.ServeHTTP(w, r)
	})
}
