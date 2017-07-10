package voting

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

/*

Go
var baseURL = "./../../" // Go local
var baseURL = "/freecodecamp/7-bonus2/" // Go deploy

HTML Templates
"./../../voting/loQueSea" || "/voting/loQueSea "// html local
<head> <base href="/freecodecamp/7-bonus2/"></head> // html deploy
"./voting/loQueSea" // hrml deploy relative URL

JS
window.location.assign(app.getBaseUrl() + 'voting/guest/');

*/

//var baseURL = "./../../" // Go local
var baseURL = "/freecodecamp/7-bonus2/" // Go deploy
var tmpl map[string]*template.Template

func init() {
	base := "voting/views/"
	guest := base + "guest.html"
	logged := base + "logged.html"
	login := base + "login.html"
	newPoll := base + "newPoll.html"
	layout := base + "partials/layout.html"
	footer := base + "partials/footer.html"
	scripts := base + "partials/scripts.html"
	polls := base + "partials/polls.html"
	tmpl = make(map[string]*template.Template)
	tmpl["guest.html"] = template.Must(template.ParseFiles(guest, layout, footer, scripts, polls))
	tmpl["logged.html"] = template.Must(template.ParseFiles(logged, layout, footer, scripts, polls))
	tmpl["login.html"] = template.Must(template.ParseFiles(login, layout, footer, scripts))
	tmpl["newPoll.html"] = template.Must(template.ParseFiles(newPoll, layout, footer, scripts))
}

// RouterVoting ...
func RouterVoting(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Original Path ==> ", r.URL.Path)
	param := strings.Split(r.URL.Path, "/")
	path := param[2:len(param)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	//path2 := r.URL.Path[len("/voting/"):]
	//fmt.Println("Going ....", path2)
	if len(path) == 0 {
		fmt.Fprintln(w, "INDEX, NOT FOUND")
		return
	}
	//fmt.Println("Going ==> ", path[0]) //len(path), path)
	switch path[0] {
	case "guest":
		if r.Method == "POST" {
			sumGuestVote(w, r)
		} else {
			guest(w, r)
		}
	case "logged":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"voting/login", 301)
			return
		}
		if r.Method == "POST" {
			sumUserVote(w, r)
		} else {
			logged(w, r)
		}
	case "login":
		if r.Method == "POST" {
			doLoginOrCreate(w, r)
		} else {
			login(w, r)
		}
	case "alreadyGuestVoted":
		alreadyGuestVoted(w, r)
	case "alreadyUserVoted":
		alreadyUserVoted(w, r)
	case "logout":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"voting/login", 301)
			return
		}
		logout(w, r)
	case "newPoll":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"voting/login", 301)
			return
		}
		if r.Method == "POST" {
			doCreateNewPoll(w, r)
		} else {
			newPoll(w, r)
		}
	case "newColumn":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"voting/login", 301)
			return
		}
		newColumn(w, r)
	default:
		fmt.Fprintln(w, "404 PAGE NOT FOUND")
	}
}

func guest(w http.ResponseWriter, r *http.Request) {
	data := dbGetGuestListPolls()
	//fmt.Println("LIST POLL --> ", data)
	//tmpl["guest.html"].ExecuteTemplate(w, "layout", p)
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", data)
}

func logged(w http.ResponseWriter, r *http.Request) {
	data := dbGetGuestListPolls()
	tmpl["logged.html"].ExecuteTemplate(w, "logged.html", data)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl["login.html"].ExecuteTemplate(w, "login.html", nil)
}

func wrongLogin(w http.ResponseWriter, r *http.Request) {
	data := "INVALID PASSWORD"
	tmpl["login.html"].ExecuteTemplate(w, "login.html", data)
}

func alreadyGuestVoted(w http.ResponseWriter, r *http.Request) {
	var data []aPoll
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", data)
}

func alreadyUserVoted(w http.ResponseWriter, r *http.Request) {
	var data []aPoll
	tmpl["logged.html"].ExecuteTemplate(w, "logged.html", data)
}

func logout(w http.ResponseWriter, r *http.Request) {
	dbDeleteCookie(r)
	http.Redirect(w, r, baseURL+"voting/guest/", 301)
}

func newPoll(w http.ResponseWriter, r *http.Request) {
	tmpl["newPoll.html"].ExecuteTemplate(w, "newPoll.html", nil)
}

func newColumn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newPoll, _ := strconv.Atoi(r.Form["poll"][0])
	newOption := r.Form["newOption"][0]
	dbInsertNewOption(newPoll, newOption)
	cookie, _ := r.Cookie("session")
	user := strings.Split(cookie.Value, ":")[0]
	http.Redirect(w, r, baseURL+"voting/logged/?user="+user, 301)
}
