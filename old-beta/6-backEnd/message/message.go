package message

import (
	"fmt"
	"freeCodeCamp/6-backEnd/util"
	"net/http"
	"strconv"
	"strings"
)

var e myerror

type myerror struct {
	Info  string `json:"info"`
	Error string `json:"error"`
}

func init() {
}

// RouterMessage ...
func RouterMessage(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	path := params[3:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Println("Going ....", path, r.Method, len(path))
	if len(path) > 2 {
		util.PageNotFound(w, r)
		return
	}
	d.loadJSONDataFromFile()
	var board int
	if len(path) > 0 {
		var ok error
		board, ok = strconv.Atoi(path[0])
		if ok != nil {
			util.PageNotFound(w, r)
			return
		}
	}
	if r.Method == "GET" {
		if len(path) == 0 {
			d.getThreads(w, r)
			return
		}
		d.getReplies(w, r, board)
	}
	if r.Method == "POST" {
		if len(path) == 0 {
			d.createThread(w, r)
			return
		}
		d.createReply(w, r, board)
	}
	if r.Method == "PUT" {
		if len(path) == 0 {
			d.reportThread(w, r)
			return
		}
		d.reportReply(w, r, board)
	}
	if r.Method == "DELETE" {
		if len(path) == 0 {
			d.deleteThread(w, r)
			return
		}
		d.deleteReply(w, r, board)
	}
}
