package main

import (
	"net"
	"net/http"
	"strings"
)

type outputParser struct {
	IP       string `json:"Ip Address"`
	Language string `json:"Language"`
	OS       string `json:"Operative System"`
}

func parser(w http.ResponseWriter, r *http.Request) {
	param := strings.Split(r.URL.Path, "/")
	var data outputParser
	if strings.ToLower(param[3]) == "whoami" {
		data.IP = getIP(r)
		data.Language = r.Header.Get("Accept-Language")
		data.OS = r.Header.Get("User-Agent")
		sendStructAsJSON(w, r, data)
	} else {
		http.Redirect(w, r, "/parser/parser.html", 301)
	}
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}
