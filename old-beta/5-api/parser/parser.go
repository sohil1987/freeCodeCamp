package parser

import (
	"freeCodeCamp/5-api/_help"
	"net"
	"net/http"
	"strings"
)

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
		http.Redirect(w, r, help.BaseURL+"parser/parser.html", 301)
	}
	var o Output
	o.IP = getIP(r)
	o.Language = r.Header.Get("Accept-Language")
	o.Os = r.Header.Get("User-Agent")
	help.StructToJSON(w, r, o)
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}
