package nightlife

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

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
	Yelp struct {
		Name   string `json:"name"`
		ID     string `json:"id"`
		Secret string `json:"secret"`
		Token  string
	} `json:"yelp"`
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
		fmt.Println(`nightlife sql.Open() OK`)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("nightlife db.Ping() OK")
	}
	//fmt.Println("Token ==>", c.Yelp.Token)
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", connPath)
}

func loadConfig() { // parse JSON with MARSHALL
	file, err := os.Open("nightlife/assets/config/secret.json")
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

func getAccessToken() {
	type yelpAccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	// API V3 https://www.yelp.es/developers/documentation/v3/get_started
	apiURL := "https://api.yelp.com"
	resource := "/oauth2/token"
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", c.Yelp.ID)
	data.Add("client_secret", c.Yelp.Secret)
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	//fmt.Println(urlStr)
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp)

	decoder := json.NewDecoder(resp.Body)
	var yelp yelpAccess
	err := decoder.Decode(&yelp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(yelp.TokenType)
	fmt.Println(yelp.AccessToken)
	fmt.Println(yelp.ExpiresIn) // 15541554
}
