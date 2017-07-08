package help

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	//fmt.Println(`Init From Help Package`)
}

// ServerIP ...
//const ServerIP = "localhost:3000" // Go local
const ServerIP = "localhost:3502" // Go deploy

// BaseURL ...
//const BaseURL = "/" // Go local
//const BaseURL = "./../../" // Go local
const BaseURL = "/freecodecamp/7-bonus2/" // Go deploy

// SecretJSON ...
const SecretJSON = "./secret.json"

// StructToJSON ...
func StructToJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

// Data ...
type Data []interface{}

// GetJSONDataFromFile ...
func GetJSONDataFromFile(pathToFile string, d *Data) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Fatalln(err)
	}
}

// Conf ...
type Conf struct {
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
	APIImage struct {
		Key   string `json:"key"`
		CseID string `json:"cseID"`
	} `json:"apiImage"`
	APIStockMarket struct {
		Key    string `json:"apiKey"`
		FileDB string `json:"fileDB"`
	} `json:"apiStockMarket"`
}

// LoadConfig ...
func LoadConfig(pathToFile string, c *Conf) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
