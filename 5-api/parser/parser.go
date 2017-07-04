package parser

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

//var baseURL = "/" // Go local
var baseURL = "/freecodecamp/5-api/" // Go deploy

func init() {
	//fmt.Println(`Init from Package Parser`)
}

// Output ...
type Output struct {
	IP       string `json:"IP Address"`
	Language string `json:"Language"`
	Os       string `json:"Operative System"`
}

// RouterParser ...
func RouterParser(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	input := params[3]
	if strings.ToLower(input) != "whoami" {
		http.Redirect(w, r, baseURL+"parser/parser.html", 301)
	}
	var o Output
	o.IP = getIP(r)
	o.Language = r.Header.Get("Accept-Language")
	o.Os = r.Header.Get("User-Agent")
	structToJSON(w, r, o)
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}

func structToJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}
