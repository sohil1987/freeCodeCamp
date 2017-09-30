package stock

import (
	"encoding/json"
	"freeCodeCamp/v2-beta/6-backEnd/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
}

const dataFile = "./stock/stock.json"

var db database

type database map[string]stock

type stock struct {
	Price float64  `json:"-"`
	IPs   []string `json:"ips"`
	Likes int      `json:"likes"`
}

func (db *database) loadJSONDataFromFile(pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &db)
	if err != nil {
		log.Fatalln(err)
	}
}

func (db database) getLikes(name string) int {
	return db[name].Likes //len(db[name].IPs)
}

func (db database) insertLike(r *http.Request, name string) {
	ip := util.GetIP(r)
	st := db[name]
	found := false
	for _, v := range st.IPs {
		if v == ip {
			found = true
		}
	}
	if !found {
		st.IPs = append(st.IPs, ip)
		st.Likes = st.Likes + 1
		db[name] = st
		db.writeJSONtoFile(dataFile)
	}
}

func (db *database) writeJSONtoFile(pathToFile string) {
	f, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(&db)
}
