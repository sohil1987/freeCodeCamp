package util

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path"
)

func init() {
	fmt.Println("init from util package")
}

// SecretJSON ...
const SecretJSON = "./secret.json"

// BaseURL ...
const BaseURL = "/" // Go local
//const BaseURL = "/freecodecamp/6-backEnd/" // Go deploy

// ServerIP ...
const ServerIP = "localhost:3000" // Go local
//const ServerIP = "localhost:3503" // Go deploy

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

// ToFixedFloat64 (untruncated, num) -> untruncated.toFixed(num)
func ToFixedFloat64(untruncated float64, precision int) float64 {
	coef := math.Pow10(precision)
	truncated := float64(int(untruncated*coef)) / coef
	return truncated
}
