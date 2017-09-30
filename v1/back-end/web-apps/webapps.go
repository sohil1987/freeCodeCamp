package main

// fresh -c tmp/fresh.conf
// GOOS=linux GOARCH=amd64 go build

import (
	"freeCodeCamp/v1/back-end/web-apps/_help"
	"freeCodeCamp/v1/back-end/web-apps/book"
	"freeCodeCamp/v1/back-end/web-apps/nightlife"
	"freeCodeCamp/v1/back-end/web-apps/pintelest"
	"freeCodeCamp/v1/back-end/web-apps/stock"
	"freeCodeCamp/v1/back-end/web-apps/voting"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// custom404 for all apps
	assets := fs404(http.Dir("./_public/404"))
	mux.Handle("/error/", http.StripPrefix("/error/", assets))

	// assets for individual apis
	stockAssets := fs404(http.Dir("./stock/assets"))
	mux.Handle("/stock/", http.StripPrefix("/stock/", stockAssets))
	nightRes := http.FileServer(http.Dir("./nightlife/assets"))
	mux.Handle("/nightlife/assets/", http.StripPrefix("/nightlife/assets/", nightRes))
	bookRes := http.FileServer(http.Dir("book/assets"))
	mux.Handle("/book/assets/", http.StripPrefix("/book/assets/", bookRes))
	votingRes := http.FileServer(http.Dir("voting/assets"))
	mux.Handle("/voting/assets/", http.StripPrefix("/voting/assets/", votingRes))

	pintelestRes := http.FileServer(http.Dir("./pintelest/assets"))
	mux.Handle("/pintelest/assets/", http.StripPrefix("/pintelest/assets/", pintelestRes))

	mux.HandleFunc("/stock/api/", stock.RouterStock)
	go stock.Hub.Start()
	mux.HandleFunc("/stock/socket/", stock.Socket)
	mux.HandleFunc("/nightlife/", nightlife.RouterNight)
	mux.HandleFunc("/book/", book.RouterBook)
	mux.HandleFunc("/voting/", voting.RouterVoting)
	mux.HandleFunc("/pintelest/", pintelest.RouterPintelest)

	mux.HandleFunc("/", pageNotFound)

	server := http.Server{
		Addr:           help.ServerIP,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, help.BaseURL+"error/404.html", 301)
	//fmt.Fprintf(w, r.URL.Path)
}

func fs404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			pageNotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}
