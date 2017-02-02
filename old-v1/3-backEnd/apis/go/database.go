package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type configuration struct {
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
	API struct {
		Key   string `json:"key"`
		CseID string `json:"cseID"`
	} `json:"api"`
}

var db *sql.DB
var c configuration

func init() {
	var err error
	loadConfig() // use this or next
	//loadConfig2() // use this or previous
	connPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Db)
	db, err = sql.Open("mysql", connPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(`sql.Open() OK`)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("db.Ping() OK")
	}
}

func loadConfig() { // parse JSON with DECODER
	file, err := os.Open("secret.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	c = configuration{}
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func loadConfig2() { // parse JSON with MARSHALL
	file, err := os.Open("secret.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	// get file content
	chunk, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(chunk, &c)

}
