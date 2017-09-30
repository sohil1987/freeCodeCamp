package main

// fresh -c tmp/fresh.conf
// GOOS=linux GOARCH=amd64 go build
// user stories for projects
// https://pricey-hugger.glitch.me/

import (
	"freeCodeCamp/v2-beta/6-backEnd/issue"
	"freeCodeCamp/v2-beta/6-backEnd/library"
	"freeCodeCamp/v2-beta/6-backEnd/message"
	"freeCodeCamp/v2-beta/6-backEnd/metric"
	"freeCodeCamp/v2-beta/6-backEnd/stock"
	"freeCodeCamp/v2-beta/6-backEnd/util"
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
	metricAssets := util.FS404(http.Dir("./metric/assets"))
	mux.Handle("/metric/", http.StripPrefix("/metric/", metricAssets))
	stockAssets := util.FS404(http.Dir("./stock/assets"))
	mux.Handle("/stock/", http.StripPrefix("/stock/", stockAssets))
	libraryAssets := util.FS404(http.Dir("./library/assets"))
	mux.Handle("/library/", http.StripPrefix("/library/", libraryAssets))
	issueAssets := util.FS404(http.Dir("./issue/assets"))
	mux.Handle("/issue/", http.StripPrefix("/issue/", issueAssets))
	messageAssets := util.FS404(http.Dir("./message/assets"))
	mux.Handle("/message/", http.StripPrefix("/message/", messageAssets))

	mux.HandleFunc("/metric/v1/", metric.RouterMetric)
	mux.HandleFunc("/stock/v1/", stock.RouterStock)
	mux.HandleFunc("/library/v1/", library.RouterLibrary)
	mux.HandleFunc("/issue/v1/", issue.RouterIssue)
	mux.HandleFunc("/message/v1/", message.RouterMessage)

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
