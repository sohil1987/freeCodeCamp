package voting

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // comment
)

type configuration struct {
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
}

var db *sql.DB
var c configuration
var connPath string

func init() {
	loadConfig()
	connPath = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Db)
	//fmt.Println(connPath)
	db, err := sql.Open("mysql", connPath)
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

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", connPath)
}

func loadConfig() { // parse JSON with DECODER
	file, err := os.Open("voting/assets/config/secret.json")
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
