package stock

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getData(w http.ResponseWriter, code string) {
	urlData := "https://www.quandl.com/api/v3/datasets/WIKI/" + code + ".json" + "?api_key=" + c.APIKey + "&column_index=11"
	//fmt.Println(urlData)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(urlData)
	//resp, err := http.Get(urlData)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		w.Write([]byte("NoT FOUND"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(body) // info sent
	// body is a string, for use we must Unmarshal over a struct
	//return
	var q quandlAPI
	err = json.Unmarshal(body, &q)
	// prepare data for send via json to browser
	sd := convertDatasetToStockData(q.Dataset.Data)
	sendStructAsJSON(w, sd)
}

func convertDatasetToStockData(data [][]interface{}) [][]float64 {
	var sd [][]float64
	sd = make([][]float64, len(data))
	for j := range sd {
		sd[j] = make([]float64, 2)
	}
	for i := range data { //sd lo creo ya del reves, del final hacia adelante
		sd[len(data)-i-1][0] = getTimestamp(data[i][0].(string)) //.(float64))
		sd[len(data)-i-1][1] = data[i][1].(float64)
	}
	return sd
}

func getTimestamp(date string) float64 {
	var t int64
	params := strings.Split(date, "-")
	day, _ := strconv.Atoi(params[2])
	month, _ := strconv.Atoi(params[1])
	year, _ := strconv.Atoi(params[0])
	auxTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	t = int64(auxTime.UnixNano() / int64(time.Millisecond))
	return float64(t)
}
