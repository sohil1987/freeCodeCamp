package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func urlRedirect(w http.ResponseWriter, r *http.Request) {
	if string(r.URL.Path[len(r.URL.Path)-1]) == "/" { //remove last /
		r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
	}
	r.URL.Path = strings.Replace(r.URL.Path, ":/", "://", -1)
	param := strings.Split(r.URL.Path, "/")
	if len(param) > 4 {
		if param[3] == "new" {
			if len(param) > 5 {
				posURL := strings.Split(r.URL.Path, "new/")
				if isValidURL(posURL[1]) {
					if !dbGetRecordByOriginal(posURL[1]) {
						rec := dbInsertNewURL(posURL[1])
						base := "https://brusbilis.com/freecodecamp/8-dynamic/apis/v1/url/"
						rec.Short = base + strconv.Itoa(rec.Idurl)
						sendStructAsJSON(w, r, rec)
						return
					}
				}
			}
			//return /// Remove to allow nonvalid url redirect to menu
		}
	}
	if len(param) == 4 {
		num, err := strconv.Atoi(param[3])
		if err == nil {
			ok, urlToGo := isValidNumber(num)
			if ok {
				http.Redirect(w, r, urlToGo, 301)
			}
		}
	}
	//http.Redirect(w, r, "/url/url.html", 301) // for local dev
	http.Redirect(w, r, "/freecodecamp/8-dynamic/apis/url/url.html", 301) //for deploy

}

func isValidURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	return true
}

func isValidNumber(num int) (bool, string) {
	//rows := dbGetRecordByIdurl(num)
	/*for _, r := range rows {
		fmt.Printf("%d, %s\n", r.Idurl, r.Original)
	}*/
	r := dbGetRecordByIdurl2(num)
	//fmt.Printf("%d, %s\n", r.Idurl, r.Original)
	if r.Original != "" {
		return true, r.Original
	}
	return false, ""
}

// DATABASE TYPES AND METHODS

type record struct {
	Idurl    int
	Original string `json:"original url"`
	Short    string `json:"short"`
}

func dbGetRecordByOriginal(Original string) bool {
	var r record
	err := db.QueryRow("SELECT * FROM url WHERE Original = ?", Original).Scan(&r.Idurl, &r.Original)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No URL saved with that url.")
		return false
	case err != nil:
		log.Fatal(err)
	}
	return true
}

func dbGetRecordByIdurl(num int) []*record {
	rows, err := db.Query("SELECT * FROM url WHERE Idurl = ?", num)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return nil
		}
	}
	defer rows.Close()

	var records []*record
	for rows.Next() {
		r := &record{}
		err := rows.Scan(&r.Idurl, &r.Original)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, r)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func dbGetRecordByIdurl2(num int) record {
	var r record
	err := db.QueryRow("SELECT * FROM url WHERE Idurl = ?", num).Scan(&r.Idurl, &r.Original)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No URL saved with that ID.")
	case err != nil:
		log.Fatal(err)
	}
	return r
}

func dbInsertNewURL(newurl string) record {
	var r record
	result, err := db.Exec("INSERT into url(Original) VALUE (?)", newurl)
	if err != nil {
		log.Fatal(err)
	}
	aux, _ := result.LastInsertId()
	r.Idurl = int(aux)
	r.Original = newurl
	return r
}
