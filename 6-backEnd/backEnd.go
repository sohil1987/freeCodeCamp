package main

// fresh -c tmp/fresh.conf
// GOOS=linux GOARCH=amd64 go build
// user stories for projects
// https://pricey-hugger.glitch.me/

import (
	"freeCodeCamp/6-backEnd/metric"
	"freeCodeCamp/6-backEnd/util"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// custom404 for all apps
	assets := util.FS404(http.Dir("./_public/404/"))
	mux.Handle("/error/", http.StripPrefix("/error/", assets))

	// assets for individual apis
	//timeAssets := http.FileServer(http.Dir("./timestamp/assets"))
	timeAssets := util.FS404(http.Dir("./metric/assets"))
	mux.Handle("/metric/", http.StripPrefix("/metric/", timeAssets))

	mux.HandleFunc("/metric/v1/", metric.RouterMetric)

	mux.HandleFunc("/", util.PageNotFound)

	server := http.Server{
		Addr:           util.ServerIP,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
