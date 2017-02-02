package main

import "net/http"

func main() {
	public := http.FileServer(http.Dir("./_public/"))

	book := http.FileServer(http.Dir("./book/"))
	nightlife := http.FileServer(http.Dir("./nightlife/"))
	//pintelest := http.FileServer(http.Dir("./pintelest/"))
	stock := http.FileServer(http.Dir("./stock/"))
	voting := http.FileServer(http.Dir("./voting/"))

	http.Handle("/", http.StripPrefix("/", public))
	http.Handle("/book/", http.StripPrefix("/book", book))
	http.Handle("/nightlife/", http.StripPrefix("/nightlife", nightlife))
	//http.Handle("/pintelest/", http.StripPrefix("/pintelest", pintelest))
	http.HandleFunc("/pintelest/", nodePintelest)

	http.Handle("/stock/", http.StripPrefix("/stock", stock))
	http.Handle("/voting/", http.StripPrefix("/voting", voting))

	server := http.Server{
		Addr: "localhost:3006",
	}
	server.ListenAndServe()
}

func nodePintelest(w http.ResponseWriter, r *http.Request) {
	url := "https://brusbilis.com/freecodecamp/old-v1/webapps/pintelest/"
	http.Redirect(w, r, url, 301)
}
