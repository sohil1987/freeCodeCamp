package stock

import (
	"freeCodeCamp/v1/back-end/web-apps/_help"
	"time"
)

var c help.Conf

func init() {
	help.LoadConfig(help.SecretJSON, &c)
}

type quandlAPI struct {
	Dataset struct {
		Collapse            interface{}     `json:"collapse"`
		ColumnIndex         int             `json:"column_index"`
		ColumnNames         []string        `json:"column_names"`
		Data                [][]interface{} `json:"data"`
		DatabaseCode        string          `json:"database_code"`
		DatabaseID          int             `json:"database_id"`
		DatasetCode         string          `json:"dataset_code"`
		Description         string          `json:"description"`
		EndDate             string          `json:"end_date"`
		Frequency           string          `json:"frequency"`
		ID                  int             `json:"id"`
		Limit               interface{}     `json:"limit"`
		Name                string          `json:"name"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		Order               interface{}     `json:"order"`
		Premium             bool            `json:"premium"`
		RefreshedAt         string          `json:"refreshed_at"`
		StartDate           string          `json:"start_date"`
		Transform           interface{}     `json:"transform"`
		Type                string          `json:"type"`
	} `json:"dataset"`
}

// https://www.quandl.com/api/v3/datasets/WIKI/FB.json?column_index=11
type quandlAPIOriginal struct {
	Dataset struct {
		ID                  int         `json:"id"`
		DatasetCode         string      `json:"dataset_code"`
		DatabaseCode        string      `json:"database_code"`
		Name                string      `json:"name"`
		Description         string      `json:"description"`
		RefreshedAt         time.Time   `json:"refreshed_at"`
		NewestAvailableDate string      `json:"newest_available_date"`
		OldestAvailableDate string      `json:"oldest_available_date"`
		ColumnNames         []string    `json:"column_names"`
		Frequency           string      `json:"frequency"`
		Type                string      `json:"type"`
		Premium             bool        `json:"premium"`
		Limit               interface{} `json:"limit"`
		Transform           interface{} `json:"transform"`
		ColumnIndex         int         `json:"column_index"`
		StartDate           string      `json:"start_date"`
		EndDate             string      `json:"end_date"`
		Data                []struct {
			Num0 string  `json:"0"`
			Num1 float64 `json:"1"`
		} `json:"data"`
		Collapse   interface{} `json:"collapse"`
		Order      interface{} `json:"order"`
		DatabaseID int         `json:"database_id"`
	} `json:"dataset"`
}
