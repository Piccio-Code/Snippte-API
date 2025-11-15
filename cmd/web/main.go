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

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO:\t", log.LstdFlags),
		errorLog: log.New(os.Stderr, "ERROR:\t", log.LstdFlags|log.Lshortfile),
	}

	mux := http.NewServeMux()

	fileServer := http.FileServerFS(StaticFolder)

	mux.Handle("/static/", http.StripPrefix("/static/", hideStaticMiddleware(fileServer)))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	app.infoLog.Printf("Starting server on %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: app.errorLog,
	}

	err := srv.ListenAndServe()
	app.errorLog.Println(err)
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
