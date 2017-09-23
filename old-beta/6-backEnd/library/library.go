package library

import (
	"fmt"
	"freeCodeCamp/6-backEnd/util"
	"net/http"
	"strconv"
	"strings"
)

func init() {

}

var e myerror

type myerror struct {
	Info  string `json:"info,omitempty"`
	Error string `json:"error,omitempty"`
}

// RouterLibrary ...
func RouterLibrary(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	path := params[3:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Println("Going ....", path, len(path))
	if len(path) == 0 || len(path) > 2 {
		util.PageNotFound(w, r)
		return
	}
	if path[0] != "books" {
		util.PageNotFound(w, r)
		return
	}
	var id = -1
	if len(path) == 2 {
		var err error
		id, err = strconv.Atoi(path[1])
		if err != nil || id < 0 { // avoid negative IDs
			//fmt.Println(err)
			e.Error = "ID is not a valid number"
			util.StructToJSON(w, r, e)
			return
		}
	}
	b.loadJSONDataFromFile()
	if r.Method == "GET" {
		if len(path) == 1 {
			fmt.Println("----- Show all ----")
			b.getAll(w, r)
		} else {
			fmt.Println("----- Show", id, "-----")
			b.getOne(w, r, id)
		}
		return
	}
	if r.Method == "POST" {
		if len(path) == 1 {
			fmt.Println("----- Add book ----")
			b.addBook(w, r)
		} else {
			fmt.Println("----- Add comment to", id, "-----")
			b.addComment(w, r, id)
		}
		return
	}
	if r.Method == "DELETE" {
		if len(path) == 1 {
			fmt.Println("----- Delete all ----")
			b.deleteAll(w, r)
		} else {
			fmt.Println("----- Delete", id, "-----")
			b.deleteOne(w, r, id)
		}
		return
	}
	util.PageNotFound(w, r)

}

/*
  add_header Cache-Control no-store;
  add_header Cache-Control no-cache;

*/
