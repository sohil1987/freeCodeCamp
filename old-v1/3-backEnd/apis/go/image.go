package main

import "net/http"

func image(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, r.Host+r.URL.Path)
	http.Redirect(w, r, "https://brusbilis.com/freecodecamp/old-v1/apis/image/", 301)

}
