package tracker

import (
	"encoding/json"
	"freeCodeCamp/v2-beta/5-api/_help"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	//fmt.Println(`Init from package tracker`)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const recordsFile = "./tracker/tracker.json"
const usersFile = "./tracker/users.json"
const lenID = 8
const securityLoops = 100

var r *rand.Rand

type users []user

type user struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type records []record

type record struct {
	User        user   `json:"user"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Date        string `json:"date"`
}

type invalid struct {
	Error string `json:"error"`
}

// RouterTracker ...
func RouterTracker(w http.ResponseWriter, r *http.Request) {
	option := strings.Split(r.URL.Path, "/")[3]
	option = strings.ToLower(option)
	switch option {
	case "users":
		if r.Method == "GET" {
			showAllUsers(w, r)
		} else {
			redirectToIndex(w, r)
		}
	case "log":
		if r.Method == "GET" {
			showUserLogs(w, r)
		} else {
			redirectToIndex(w, r)
		}
	case "newuser":
		if r.Method == "POST" {
			addNewUser(w, r)
		} else {
			redirectToIndex(w, r)
		}
	case "add":
		if r.Method == "POST" {
			addNewExercise(w, r)
		} else {
			redirectToIndex(w, r)
		}
	default:
		redirectToIndex(w, r)
	}
}

func redirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, help.BaseURL+"tracker/tracker.html", 301)
}

func showAllUsers(w http.ResponseWriter, r *http.Request) {
	var u users
	readUsersFromFile(usersFile, &u)
	help.StructToJSON(w, r, u)
}

func showUserLogs(w http.ResponseWriter, r *http.Request) {
	var data records
	var i invalid
	readDataFromFile(recordsFile, &data)
	var userLogs records
	layout := "2006-01-02" // 15:04:05"
	user := r.URL.Query().Get("user")
	if user == "" {
		redirectToIndex(w, r)
		return
	}
	var u users
	readUsersFromFile(usersFile, &u)
	found := false
	for _, v := range u {
		if v.Username == user {
			found = true
		}
	}
	if !found {
		i.Error = user + " not exists"
		help.StructToJSON(w, r, i)
		return
	}

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	if from == "" && to == "" && limit == 0 {
		for _, v := range data {
			if v.User.Username == user {
				userLogs = append(userLogs, v)
			}
		}
		help.StructToJSON(w, r, userLogs)
		return
	}
	if from == "" {
		from = "1971-01-01" // 00:00:00"
	}
	if to == "" {
		to = time.Now().Format(layout)
	}

	fromTime, err := time.Parse(layout, from)
	if err != nil {
		i.Error = from + " is not a valid date"
		help.StructToJSON(w, r, i)
		return
	}
	toTime, err := time.Parse(layout, to)
	if err != nil {
		i.Error = to + " is not a valid date"
		help.StructToJSON(w, r, i)
		return
	}
	if toTime.Sub(fromTime) < 0 {
		i.Error = from + " is a later date than " + to + " . Turn them around"
		help.StructToJSON(w, r, i)
		return
	}
	for _, v := range data {
		if v.User.Username == user {
			if v.Duration <= limit && toTime.Sub(fromTime) > 0 {
				userLogs = append(userLogs, v)
			}
		}
	}
	help.StructToJSON(w, r, userLogs)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		redirectToIndex(w, r)
		return
	}
	var u users
	var i invalid
	var one user
	readUsersFromFile(usersFile, &u)
	for _, v := range u {
		if v.Username == username {
			i.Error = username + " already exists"
			help.StructToJSON(w, r, i)
			return
		}
	}
	one.ID = createNewID()
	one.Username = username
	u = append(u, one)
	writeUsersToFile(usersFile, &u)
	help.StructToJSON(w, r, u)
}

func addNewExercise(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	layout := "2006-01-02" // 15:04:05"
	username := r.Form.Get("username")
	var i invalid
	var u users
	var rec record
	var data records
	readDataFromFile(recordsFile, &data)
	readUsersFromFile(usersFile, &u)
	exists := false
	if username == "" {
		redirectToIndex(w, r)
		return
	}
	for _, v := range u {
		if v.Username == username {
			rec.User.Username = username
			rec.User.ID = v.ID
			exists = true
		}
	}
	if !exists {
		i.Error = username + " not exists"
		help.StructToJSON(w, r, i)
		return
	}
	description := r.Form.Get("description")
	duration, err := strconv.Atoi(r.Form.Get("duration"))
	if err != nil {
		duration = 0
	}
	date := r.Form.Get("date")
	if username == "" || description == "" || duration == 0 {
		redirectToIndex(w, r)
		return
	}
	if date == "" {
		date = time.Now().Format(layout)
	} else {
		_, err := time.Parse(layout, date)
		if err != nil {
			redirectToIndex(w, r)
			return
		}
	}
	rec.Description = description
	rec.Duration = duration
	rec.Date = date
	showLogs := "tracker/v1/log?user=" + rec.User.Username
	data = append(data, rec)
	writeDatatoFile(recordsFile, &data)
	http.Redirect(w, r, help.BaseURL+showLogs, 301)
}

func createNewID() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHJLMNOPQRSTUVWXYZ"
	result := ""
	var data records
	readDataFromFile(recordsFile, &data)
	isNew := true
	found := false
	sec := 0
	for !found && sec < securityLoops {
		for i := 0; i < lenID; i++ {
			result += string(chars[r.Intn(len(chars))])
		}
		for _, v := range data {
			if v.User.ID == result {
				isNew = false
			}
		}
		if isNew == true {
			found = true
		}
		sec++
	}
	return result
}

func readDataFromFile(recordsFile string, data *records) {
	file, err := os.Open(recordsFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
}

func readUsersFromFile(usersFile string, u *users) {
	file, err := os.Open(usersFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeUsersToFile(usersFile string, u *users) {
	f, err := os.Create(usersFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(u)
}

func writeDatatoFile(recordsFile string, data *records) {
	f, err := os.Create(recordsFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(data)
}
