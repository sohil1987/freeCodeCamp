package nightlife

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type dataAPI struct {
	Businesses []*struct {
		Going        int
		Distance     float64 `json:"-"`
		IsClosed     bool    `json:"-"`
		DisplayPhone string  `json:"-"`
		URL          string  `json:"url"`
		ImageURL     string  `json:"image_url"`
		Phone        string  `json:"-"`
		Coordinates  struct {
			Longitude float64 `json:"-"`
			Latitude  float64 `json:"-"`
		} `json:"-"`
		Location struct {
			ZipCode        string      `json:"-"`
			Address2       interface{} `json:"-"`
			Address3       interface{} `json:"-"`
			Address1       string      `json:"-"`
			City           string      `json:"city"`
			Country        string      `json:"-"`
			State          string      `json:"-"`
			DisplayAddress []string    `json:"-"`
		} `json:"location"`
		ID         string  `json:"id"`
		Rating     float64 `json:"-"`
		Categories []struct {
			Alias string `json:"-"`
			Title string `json:"-"`
		} `json:"-"`
		Name        string `json:"name"`
		ReviewCount int    `json:"-"`
		Price       string `json:"-"`
	} `json:"businesses"`
	Region struct {
		Center struct {
			Longitude float64 `json:"-"`
			Latitude  float64 `json:"-"`
		} `json:"-"`
	} `json:"-"`
	Total  int `json:"total"`
	Search string
	User   string
}

func loadDataDev() dataAPI { // parse JSON with MARSHALL
	var d dataAPI
	file, err := os.Open("nightlife/assets/config/data.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	// get file content
	chunk, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(chunk, &d)
	return d
}

/*
--------- Original Struct---------
*/

type dataAPIOriginal struct {
	Businesses []struct {
		Distance     float64 `json:"distance"`
		IsClosed     bool    `json:"is_closed"`
		DisplayPhone string  `json:"display_phone"`
		URL          string  `json:"url"`
		ImageURL     string  `json:"image_url"`
		Phone        string  `json:"phone"`
		Coordinates  struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		} `json:"coordinates"`
		Location struct {
			ZipCode        string      `json:"zip_code"`
			Address2       interface{} `json:"address2"`
			Address3       interface{} `json:"address3"`
			Address1       string      `json:"address1"`
			City           string      `json:"city"`
			Country        string      `json:"country"`
			State          string      `json:"state"`
			DisplayAddress []string    `json:"display_address"`
		} `json:"location"`
		ID         string  `json:"id"`
		Rating     float64 `json:"rating"`
		Categories []struct {
			Alias string `json:"alias"`
			Title string `json:"title"`
		} `json:"categories"`
		Name        string `json:"name"`
		ReviewCount int    `json:"review_count"`
		Price       string `json:"price,omitempty"`
	} `json:"businesses"`
	Region struct {
		Center struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		} `json:"center"`
	} `json:"region"`
	Total int `json:"total"`
}
