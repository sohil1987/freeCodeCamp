package main

import (
	"freeCodeCamp/old-v1/3-backEnd/webApps/book"
	"freeCodeCamp/old-v1/3-backEnd/webApps/nightlife"
	"freeCodeCamp/old-v1/3-backEnd/webApps/stock"
	"freeCodeCamp/old-v1/3-backEnd/webApps/voting"

	"net/http"
)

func main() {
	//pwd, _ := os.Getwd()
	//fmt.Println("We are here ->", pwd)
	//mainNormal()
	mainMux()
}

func mainMux() {
	mux := http.NewServeMux()

	// assets for all apps
	assets := http.FileServer(http.Dir("_public"))
	mux.Handle("/", http.StripPrefix("/", assets))

	// assets for individual apps
	votingRes := http.FileServer(http.Dir("voting/assets"))
	mux.Handle("/voting/assets/", http.StripPrefix("/voting/assets/", votingRes))
	nightRes := http.FileServer(http.Dir("nightlife/assets"))
	mux.Handle("/nightlife/assets/", http.StripPrefix("/nightlife/assets/", nightRes))
	bookRes := http.FileServer(http.Dir("book/assets"))
	mux.Handle("/book/assets/", http.StripPrefix("/book/assets/", bookRes))
	stockRes := http.FileServer(http.Dir("stock/assets/"))
	mux.Handle("/stock/", http.StripPrefix("/stock/", stockRes))

	mux.HandleFunc("/voting/", voting.RouterVoting)
	mux.HandleFunc("/nightlife/", nightlife.RouterNight)
	mux.HandleFunc("/book/", book.RouterBook)
	mux.HandleFunc("/stock/api/", stock.RouterStock)
	go stock.Hub.Start()
	mux.HandleFunc("/stock/socket/", stock.Socket)

	mux.HandleFunc("/pintelest/", nodePintelest)

	server := http.Server{
		Addr:    "localhost:3006",
		Handler: mux,
	}
	server.ListenAndServe()
}

func mainNormal() {
	// assets for all apps
	assets := http.FileServer(http.Dir("_public"))
	http.Handle("/", http.StripPrefix("/", assets))

	// assets for individual apps
	votingRes := http.FileServer(http.Dir("voting/assets"))
	http.Handle("/voting/assets/", http.StripPrefix("/voting/assets/", votingRes))

	book := http.FileServer(http.Dir("./book/"))
	nightlife := http.FileServer(http.Dir("./nightlife/"))
	stock := http.FileServer(http.Dir("./stock/"))

	http.Handle("/book/", http.StripPrefix("/book", book))
	http.Handle("/nightlife/", http.StripPrefix("/nightlife", nightlife))
	http.Handle("/stock/", http.StripPrefix("/stock", stock))

	// any /voting/* will redirect to voting.Voting
	http.HandleFunc("/voting/", voting.RouterVoting)

	// any /pintelest/* will redirect to voting.Voting
	http.HandleFunc("/pintelest/", nodePintelest)

	server := http.Server{
		Addr: "127.0.0.1:3006",
	}
	server.ListenAndServe()
}

func nodePintelest(w http.ResponseWriter, r *http.Request) {
	url := "https://brusbilis.com/freecodecamp/old-v1/webapps/pintelest/"
	http.Redirect(w, r, url, 301)
}
