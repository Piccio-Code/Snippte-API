package main

import (
	_ "embed"
	"flag"
	. "github.com/Piccio-Code/Snippte-API"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServerFS(StaticFolder)

	mux.Handle("/static/", http.StripPrefix("/static/", hideStaticMiddleware(fileServer)))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func hideStaticMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
