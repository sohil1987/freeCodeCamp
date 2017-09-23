package stock

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

/*

Go
var baseURL = "./../../" // Go local
var baseURL = "/freecodecamp/old-v1/webapps/" // Go deploy

HTML Templates
"./../../voting/loQueSea" || "/voting/loQueSea "// html local
<head> <base href="/freecodecamp/old-v1/webapps/"></head> // html deploy
"./voting/loQueSea" // hrml deploy relative URL

JS
window.location.assign(app.getBaseUrl() + 'voting/guest/');

*/

// RouterStock ...
func RouterStock(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Original Path ==> ", r.URL.Path)
	param := strings.Split(r.URL.Path, "/")
	path := param[2:len(param)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	r.ParseForm()
	var code string
	val, ok := r.Form["code"] // check if add param exists
	if !ok {
		getInitialCode(w)
		return
	}
	code = val[0]
	//fmt.Println("Path ....", path)
	if len(path) <= 1 {
		fmt.Fprintln(w, "INDEX, NOT FOUND")
		return
	}
	fmt.Println("Going ==> ", path[1]) //len(path), path)
	switch path[1] {
	case "get":
		if codeExist(code) {
			getData(w, code)
			return
		}
		getInitialCode(w)
	case "add":
		list := readFile(c.APIStockMarket.FileDB)
		listSize := strings.Split(list, "\n")
		if codeExist(code) && len(listSize) < 3 { // max 3 stocks
			addCode(w, r, code)
			return
		}
		getInitialCode(w)
	case "del":
		if codeExist(code) {
			delCode(w, r, code)
			return
		}
		getInitialCode(w)
	default:
		http.Redirect(w, r, "/stock/stock.html", 301)
		//fmt.Fprintln(w, "404 PAGE NOT FOUND")
	}
}

func sendStructAsJSON(w http.ResponseWriter, data interface{}) {
	//fmt.Println("Data pre JSON --> ", data)
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Data JSON --> ",string(dataJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

// Socket2 ... from std library
func Socket2(ws *websocket.Conn) {
	type message struct {
		Message string `json:"message"`
	}
	for {
		var m message
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", m.Message)
		// send a response
		m2 := message{"Ok, Go on"}
		err := websocket.JSON.Send(ws, m2)
		if err != nil {
			fmt.Println(`FAIL`)
			log.Println(err)
			break
		}
		if err == nil {
			fmt.Println(`LOLAZO`)
		}
	}
}
