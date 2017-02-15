package nightlife

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type session struct {
}

func isLogged(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		return false
	}
	aux := strings.Split(cookie.Value, ":")
	//fmt.Println(`LOL ==> USER, SESSIONID`, aux[0], aux[1])
	ok := dbSearchCookie(aux[0], aux[1])
	if ok {
		return true
	}
	return false
}

func setSessionCookie(w http.ResponseWriter, r *http.Request, user string) {
	n := getRandomInt(0, 900000000)
	sessionID := user + ":" + sha2(user+strconv.Itoa(n))
	/*http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: sessionID,
	})*/
	cookie := http.Cookie{
		Name:  "session",
		Value: sessionID,
		Path:  baseURL + "nightlife/",
	}
	http.SetCookie(w, &cookie)
	err := dbSaveCookie(user, sessionID)
	if err != nil {
		log.Fatal(err)
	}
}

func sha2(str string) string {
	bytes := []byte(str)
	// Converts string to sha2
	h := sha256.New()                   // new sha256 object
	h.Write(bytes)                      // data is now converted to hex
	code := h.Sum(nil)                  // code is now the hex sum
	codestr := hex.EncodeToString(code) // converts hex to string
	return codestr
}

/*  DATABASE METHODS */

func dbSaveCookie(user, sessionID string) error {
	db, err := connectDB()
	_, err = db.Exec("INSERT INTO nightlife.sessions (Username, SessionID)     values (?, ?)", user, sessionID)
	if err != nil {
		return err
	}
	return nil
}

func dbDeleteCookie(r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}
	username := strings.Split(cookie.Value, ":")[0]
	db, err := connectDB()
	rows, err := db.Exec("DELETE FROM nightlife.sessions WHERE USername = ?", username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
		} else {
			log.Fatal(err)
		}
	}
	res, _ := rows.RowsAffected()
	fmt.Sprintln("DELETED", res)
}

func dbSearchCookie(user, sessionID string) bool {
	var u, s string
	//fmt.Println(`User, SESSIONID`, user, sessionID)
	db, err := connectDB()
	row := db.QueryRow("SELECT * FROM nightlife.sessions WHERE Username = ? AND SessionID = ?", user, user+":"+sessionID)
	defer db.Close()
	err = row.Scan(&u, &s)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
		} else {
			log.Fatal(err)
		}
	} else {
		return true // exists vote
	}
	return false
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
