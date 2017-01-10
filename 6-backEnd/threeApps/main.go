package main

import "net/http"

func main() {
	public := http.FileServer(http.Dir("./menu/"))
	voting := http.FileServer(http.Dir("./_voting/"))
	night := http.FileServer(http.Dir("./_night/"))
	book := http.FileServer(http.Dir("./_book/"))

	http.Handle("/", http.StripPrefix("/", public))
	http.Handle("/voting/", http.StripPrefix("/voting", voting))
	http.Handle("/night/", http.StripPrefix("/night", night))
	http.Handle("/book/", http.StripPrefix("/book", book))
	http.ListenAndServe(":3007", nil)
}
