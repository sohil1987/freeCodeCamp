package voting

import (
	"math/rand"
	"net"
	"net/http"
	"time"
)

// GetIP returns string with IP
func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
