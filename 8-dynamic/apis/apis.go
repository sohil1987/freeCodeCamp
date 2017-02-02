package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/time/", timestamp)
	mux.HandleFunc("/v1/parser/", parser)
	mux.HandleFunc("/v1/url/", urlRedirect)
	mux.HandleFunc("/v1/image/", image)
	mux.HandleFunc("/v1/file/", file)

	public := http.FileServer(http.Dir("./static/"))
	mux.Handle("/", http.StripPrefix("/", public))

	server := http.Server{
		Addr:    "127.0.0.1:3005",
		Handler: mux,
	}
	server.ListenAndServe()
}

func sliceContainsString(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func sendStructAsJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	//fmt.Println("Data pre JSON --> ", data)
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Data JSON --> ",string(dataJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}
