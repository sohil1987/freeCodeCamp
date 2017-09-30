package stock

import (
	"encoding/json"
	"fmt"
	"freeCodeCamp/v2-beta/6-backEnd/util"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Sprintln(util.C)
}

/*
http://localhost:3000/stock/v1/stockPrices?stock=goog&stock=msft&like=true
*/

var price float64

var ex1 export1
var ex2 export2

type export1 struct {
	Stock string `json:"stock"`
	Price string `json:"price"`
	Likes int    `json:"likes"`
}

type export2 struct {
	StockData []exportStock
}

type exportStock struct {
	Stock    string `json:"stock"`
	Price    string `json:"price"`
	RelLikes int    `json:"rel_likes"`
}

var e error

type error struct {
	Error string
}

// RouterStock ...
func RouterStock(w http.ResponseWriter, r *http.Request) {
	e = error{}
	params := strings.Split(r.URL.Path, "/")
	path := params[3:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Println("Going ....", path)
	if len(path) == 0 || len(path) > 1 {
		util.PageNotFound(w, r)
		return
	}
	if path[0] != "stockPrices" {
		util.PageNotFound(w, r)
		return
	}
	r.ParseForm()
	stocks := r.Form["stock"]  // get param as []string
	like := r.Form.Get("like") // get param as string
	if like != "true" {
		like = "false"
	}
	fmt.Println(stocks, like)
	var prices [2]float64
	for i, v := range stocks {
		if v != "" {
			ok := getDataFromAPIToStruct(w, stocks[i])
			if !ok {
				e.Error = fmt.Sprintf("Stock %s not found", stocks[i])
				util.StructToJSON(w, r, e)
				return
			}
			getInterfacesTypes(data)
			prices[i] = price
			//prices[i] = util.ToFixedFloat64(util.GetRandomFloat64(1, 1000), 2)
		}
	}
	ex1 = export1{}
	ex2.StockData = make([]exportStock, 2)
	db.loadJSONDataFromFile(dataFile)
	if len(stocks) == 1 {
		ex1.Stock = strings.ToUpper(stocks[0])
		ex1.Price = strconv.FormatFloat(prices[0], 'g', -1, 64)
		if like == "true" {
			db.insertLike(r, ex1.Stock)
		}
		ex1.Likes = db.getLikes(ex1.Stock)
		util.StructToJSON(w, r, ex1)
		return
	} else if len(stocks) == 2 {
		ex2.StockData[0].Stock = strings.ToUpper(stocks[0])
		ex2.StockData[0].Price = strconv.FormatFloat(prices[0], 'g', -1, 64)
		ex2.StockData[1].Stock = strings.ToUpper(stocks[1])
		ex2.StockData[1].Price = strconv.FormatFloat(prices[1], 'g', -1, 64)
		if like == "true" {
			db.insertLike(r, ex2.StockData[0].Stock)
			db.insertLike(r, ex2.StockData[1].Stock)
		}
		likes1 := db.getLikes(ex2.StockData[0].Stock)
		likes2 := db.getLikes(ex2.StockData[1].Stock)
		ex2.StockData[0].RelLikes = likes1 - likes2
		ex2.StockData[1].RelLikes = likes2 - likes1
		util.StructToJSON(w, r, ex2)
		return
	} else {
		util.PageNotFound(w, r)
		return
	}
}

var data dataSet

type dataSet interface{}

func getInterfacesTypes(f interface{}) {
	switch vf := f.(type) {
	case map[string]interface{}:
		//fmt.Println("is a map:")
		for k, v := range vf {
			switch vv := v.(type) {
			case string:
				//fmt.Printf("%v: is string - %q\n", k, vv)
			case int:
				//fmt.Printf("%v: is int - %q\n", k, vv)
			case float64:
				//fmt.Printf("%v: is float64 - %g\n", k, vv)
			default:
				fmt.Sprintln(k, v, vv)
				//fmt.Printf("%v: ", k)
				getInterfacesTypes(v)
			}
		}
	case []interface{}:
		//fmt.Println("is an array:")
		for k, v := range vf {
			switch vv := v.(type) {
			case string:
				//fmt.Printf("%v: is string - %q\n", k, vv)
			case int:
				//fmt.Printf("%v: is int - %q\n", k, vv)
			case float64:
				//fmt.Printf("%v: is float64 - %g\n", k, vv)
				if k == 4 {
					//fmt.Println("Found ==> ", vv)
					price = vv
				}
			default:
				fmt.Sprintln(k, v, vv)
				//fmt.Printf("%v: ", k)
				getInterfacesTypes(v)
			}
		}
	}
}

func getDataFromAPIToStruct(w http.ResponseWriter, company string) bool {
	urlData := "https://www.quandl.com/api/v3/datasets/WIKI/" + company + "/data.json?api_key=" /*+ util.C.Quandl.APIKey */ + "&limit=1"
	//fmt.Println(urlData)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(urlData)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		//w.Write([]byte("Stock Not FOUND"))
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return true
}
