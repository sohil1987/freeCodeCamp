package util

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

func init() {
	C.LoadConfig(SecretJSON)
}

// SecretJSON ...
const SecretJSON = "./secret.json"

// BaseURL ...
const BaseURL = "/" // Go local
//const BaseURL = "/freecodecamp/6-backEnd/" // Go deploy

// ServerIP ...
const ServerIP = "localhost:3000" // Go local
//const ServerIP = "localhost:3503" // Go deploy

// C ...
var C Conf

// Conf ...
type Conf struct {
	Forge struct {
		APIKey string `json:"apiKey"`
	} `json:"forge"`
	Quandl struct {
		APIKey string `json:"apiKey"`
	} `json:"quandl"`
}

// StructToJSON ...
func StructToJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

// PageNotFound ...
func PageNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, BaseURL+"error/404.html", 301)
}

// FS404 ...
func FS404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			PageNotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

// GetRandomFloat64 [min, max] both included
func GetRandomFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return (rand.Float64() * (max - min)) + (min)
}

// ToFixedFloat64 (untruncated, num) -> untruncated.toFixed(num)
func ToFixedFloat64(untruncated float64, precision int) float64 {
	coef := math.Pow10(precision)
	truncated := float64(int(untruncated*coef)) / coef
	return truncated
}

// GetIP ...
func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}

// LoadConfig ...
func (C *Conf) LoadConfig(pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&C)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
