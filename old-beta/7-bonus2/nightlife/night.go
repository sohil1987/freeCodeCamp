package nightlife

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

/* NIGHTLIFE
Go
var baseURL = "./../../" // Go local
//var baseURL = "/freecodecamp/7-bonus2/" // Go deploy
HTML Templates
"/nightlife/loQueSea"// html local absolute path
<head> <base href="/freecodecamp/7-bonus2/"></head> // html deploy
"./voting/loQueSea" // hrml deploy relative URL
COOKIES
//Path:  "/nightlife/",
//Path:  baseURL + "nightlife/", // deploy
*/

//var baseURL = "./../../" // Go local
var baseURL = "/freecodecamp/7-bonus2/" // Go deploy
var tmpl map[string]*template.Template

func init() {
	base := "nightlife/views/"
	guest := base + "guest.html"
	logged := base + "logged.html"
	login := base + "login.html"
	layout := base + "partials/layout.html"
	footer := base + "partials/footer.html"
	scripts := base + "partials/scripts.html"
	events := base + "partials/events.html"
	tmpl = make(map[string]*template.Template)
	tmpl["guest.html"] = template.Must(template.ParseFiles(guest, layout, footer, scripts, events))
	tmpl["logged.html"] = template.Must(template.ParseFiles(logged, layout, footer, scripts, events))
	tmpl["login.html"] = template.Must(template.ParseFiles(login, layout, footer, scripts))
}

// RouterNight ...
func RouterNight(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Original Path ==> ", r.URL.Path)
	param := strings.Split(r.URL.Path, "/")
	path := param[2:len(param)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	//path2 := r.URL.Path[len("/nightlife/"):]
	fmt.Println("Going ....", path)
	if len(path) == 0 {
		fmt.Fprintln(w, "INDEX, NOT FOUND")
		return
	}
	switch path[0] {
	case "guest":
		if r.Method == "POST" {
			guestWithSearch(w, r)
		} else {
			guest(w, r)
		}
	case "login":
		if r.Method == "POST" {
			doLoginOrCreate(w, r)
		} else {
			login(w, r)
		}
	case "logged":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"nightlife/login", 301)
			return
		}
		logged(w, r)
	case "logout":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"nightlife/login", 301)
			return
		}
		logout(w, r)
	case "renoveAccessToken":
		getAccessToken()
	case "vote":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"nightlife/login", 301)
			return
		}
		resolveVote(w, r)
	default:
		fmt.Fprintln(w, "404 PAGE NOT FOUND")
	}
}

func guest(w http.ResponseWriter, r *http.Request) {
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", nil)
}

func guestWithSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	search := r.Form["search"][0]
	//data := searchBarInLocation2()
	data := searchBarInLocation(search, "")
	data.Search = search
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", data)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl["login.html"].ExecuteTemplate(w, "login.html", nil)
}

func logged(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form["user"][0]
	search := r.Form["search"][0]
	//data := searchBarInLocation2()
	data := searchBarInLocation(search, user)
	data.User = user
	data.Search = search
	dbSaveLastSearch(user, search)
	tmpl["logged.html"].ExecuteTemplate(w, "logged.html", data)
}

func logout(w http.ResponseWriter, r *http.Request) {
	dbDeleteCookie(r)
	http.Redirect(w, r, baseURL+"nightlife/guest/", 301)
}

func wrongLogin(w http.ResponseWriter, r *http.Request) {
	data := "INVALID PASSWORD"
	tmpl["login.html"].ExecuteTemplate(w, "login.html", data)
}

func resolveVote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form["user"][0]
	search := r.Form["search"][0]
	id := r.Form["id"][0]
	if dbUserAlreadyVotedBar(user, id) {
		dbRemoveUserVoteBar(user, id)
	} else {
		dbAddUserVoteBar(user, id)
	}
	params := "?user=" + user + "&search=" + search
	http.Redirect(w, r, baseURL+"nightlife/logged/"+params, 301)
}
