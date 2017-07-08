package main

// fresh -c tmp/fresh.conf
// GOOS=linux GOARCH=amd64 go build

import (
	"freeCodeCamp/7-bonus2/_help"
	"freeCodeCamp/7-bonus2/search"
	"freeCodeCamp/7-bonus2/stock"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// custom404 for all apps
	assets := fs404(http.Dir("./_public/404/"))
	mux.Handle("/error/", http.StripPrefix("/error/", assets))

	// assets for individual apis
	searchAssets := fs404(http.Dir("./search/assets"))
	mux.Handle("/search/", http.StripPrefix("/search/", searchAssets))
	stockAssets := fs404(http.Dir("./stock/assets"))
	mux.Handle("/stock/", http.StripPrefix("/stock/", stockAssets))

	mux.HandleFunc("/search/v1/", search.RouterSearch)
	mux.HandleFunc("/stock/api/", stock.RouterStock)
	go stock.Hub.Start()
	mux.HandleFunc("/stock/socket/", stock.Socket)

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
