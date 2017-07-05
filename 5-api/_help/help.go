package help

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	//fmt.Println(`Init From Help Package`)
}

// BaseURL ...
//var BaseURL = "/" // Go local
var BaseURL = "/freecodecamp/5-api/" // Go deploy

// ServerIP ...
//var ServerIP = "localhost:3000" // Go local
var ServerIP = "localhost:3501" // Go deploy

// StructToJSON ...
func StructToJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}
