package main

// fresh -c tmp/fresh.conf
// GOOS=linux GOARCH=amd64 go build

import (
	"freeCodeCamp/v1/back-end/apis/_help"
	"freeCodeCamp/v1/back-end/apis/files"
	"freeCodeCamp/v1/back-end/apis/parser"
	"freeCodeCamp/v1/back-end/apis/search"
	"freeCodeCamp/v1/back-end/apis/timestamp"
	"freeCodeCamp/v1/back-end/apis/url"
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
	timeAssets := http.FileServer(http.Dir("./timestamp/assets"))
	//timeAssets := fs404(http.Dir("./timestamp/assets"))
	mux.Handle("/time/", http.StripPrefix("/time/", timeAssets))

	parserAssets := fs404(http.Dir("./parser/assets"))
	mux.Handle("/parser/", http.StripPrefix("/parser/", parserAssets))

	fileAssets := fs404(http.Dir("./files/assets"))
	mux.Handle("/file/", http.StripPrefix("/file/", fileAssets))

	urlAssets := fs404(http.Dir("./url/assets"))
	mux.Handle("/url/", http.StripPrefix("/url/", urlAssets))

	searchAssets := fs404(http.Dir("./search/assets"))
	mux.Handle("/search/", http.StripPrefix("/search/", searchAssets))

	mux.HandleFunc("/time/v1/", timestamp.RouterTime)
	mux.HandleFunc("/parser/v1/", parser.RouterParser)
	mux.HandleFunc("/file/v1/", files.RouterFiles)
	mux.HandleFunc("/url/v1/", url.RouterURL)
	mux.HandleFunc("/search/v1/", search.RouterSearch)

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
