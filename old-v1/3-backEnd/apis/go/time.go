package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type outputTime struct {
	Unix    int    `json:"unix"`
	Natural string `json:"natural"`
	auxTime time.Time
}

var months = []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}

func timestamp(w http.ResponseWriter, r *http.Request) {
	var data outputTime
	param := strings.Split(r.URL.Path, "/")
	unix, err := strconv.Atoi(param[3])
	if err == nil { // param is a integer
		data.Unix = unix
		data.auxTime = time.Unix(int64(unix), 0)
		data.Natural = data.auxTime.Month().String() + " " + strconv.Itoa(data.auxTime.Day()) + ", " + strconv.Itoa(data.auxTime.Year())
	} else {
		param = strings.Split(param[3], " ")
		fmt.Println(param)
		if len(param) == 3 {
			ok := true
			day := param[1]
			if string(day[len(day)-1]) == "," {
				day = day[0 : len(day)-1]
			}
			dayNum, err := strconv.Atoi(day)
			if err != nil || dayNum < 1 || dayNum > 31 {
				ok = false // dayNum not integer or not in range 1-31
			}
			month := param[0]
			if !sliceContainsString(strings.ToLower(month), months) {
				ok = false
			}
			year := param[2]
			yearNum, err := strconv.Atoi(year)
			if err != nil || yearNum < 1 { // yearNum not a integer oo BC
				ok = false
			}
			if ok {
				fmt.Println("OK", month, day, year)
				monthNum, _ := strconv.Atoi(month)
				auxTime := time.Date(yearNum, time.Month(monthNum), dayNum, 0, 0, 0, 0, time.UTC)
				data.Unix = int(auxTime.Unix())
				data.Natural = month + " " + day + ", " + year
			}
		}
	}
	sendStructAsJSON(w, r, data)
}
