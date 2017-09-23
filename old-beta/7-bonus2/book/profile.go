package book

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type userProfile struct {
	FullName   string
	City       string
	State      string
	ActiveUser string
}

/*
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}
*/

func dbSaveNewProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form["user"][0]
	full := r.Form["full"][0]
	city := r.Form["city"][0]
	state := r.Form["state"][0]
	var err error
	if full != "" {
		_, err = db.Exec("UPDATE book.users SET FullName=? WHERE Username=?", full, user)
	}
	if city != "" {
		_, err = db.Exec("UPDATE book.users SET City=? WHERE Username=?", city, user)
	}
	if state != "" {
		_, err = db.Exec("UPDATE book.users SET State=? WHERE Username=?", state, user)
	}
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, baseURL+"/book/profile/?user="+user, 301)
}

func dbGetUserProfile(user string) userProfile {
	//up := userProfile{}
	var up userProfile
	var s1, s2, s3 sql.NullString
	row := db.QueryRow("SELECT FullName, City, State FROM book.users WHERE Username=?", user)
	err := row.Scan(&s1, &s2, &s3)
	//err = row.Scan(&up.FullName, &up.City, &up.State)
	if err != nil {
		fmt.Println(err)
	}
	if s1.Valid {
		up.FullName = s1.String
	} else {
		up.FullName = ""
	}
	if s2.Valid {
		up.City = s2.String
	} else {
		up.City = ""
	}
	if s3.Valid {
		up.State = s3.String
	} else {
		up.State = ""
	}
	return up
}
