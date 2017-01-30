package main

import "net/http"

func main() {
	public := http.FileServer(http.Dir("./_public/"))

	book := http.FileServer(http.Dir("./book/"))
	nightlife := http.FileServer(http.Dir("./nightlife/"))
	pintelest := http.FileServer(http.Dir("./pintelest/"))
	stock := http.FileServer(http.Dir("./stock/"))
	voting := http.FileServer(http.Dir("./voting/"))

	http.Handle("/", http.StripPrefix("/", public))
	http.Handle("/book/", http.StripPrefix("/book", book))
	http.Handle("/nightlife/", http.StripPrefix("/nightlife", nightlife))
	http.Handle("/pintelest/", http.StripPrefix("/pintelest", pintelest))
	http.Handle("/stock/", http.StripPrefix("/stock", stock))
	http.Handle("/voting/", http.StripPrefix("/voting", voting))

	server := http.Server{
		Addr: "localhost:3007",
	}
	server.ListenAndServe()
}
