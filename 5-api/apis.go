package main

// fresh -c tmp/fresh.conf

import (
	"freeCodeCamp/5-api/timestamp"
	"net/http"
	"os"
	"path"
)

//var baseURL = "/" // Go local
var baseURL = "/freecodecamp/5-api/" // Go deploy

func main() {
	//fmt.Println("Init Timestamp API")
	mux := http.NewServeMux()

	// custom404 for all apps
	assets := fs404(http.Dir("./_public/404/"))
	mux.Handle("/error/", http.StripPrefix("/error/", assets))

	// assets for individual apis
	//timeAssets := http.FileServer(http.Dir("./timestamp/assets"))
	timeAssets := fs404(http.Dir("./timestamp/assets"))
	mux.Handle("/time/", http.StripPrefix("/time/", timeAssets))

	mux.HandleFunc("/time/v1/", timestamp.TimeRouter)
	mux.HandleFunc("/", pageNotFound)
	server := http.Server{
		Addr:    "localhost:3501",
		Handler: mux,
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, baseURL+"error/404.html", 301)
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
