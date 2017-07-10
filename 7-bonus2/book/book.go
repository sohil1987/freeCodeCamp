package book

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

/* book
Go
var baseURL = "./../../" // Go local
//var baseURL = "/freecodecamp/7-bonus2/" // Go deploy
HTML Templates
"/book/loQueSea"// html local absolute path
<head> <base href="/freecodecamp/7-bonus2/"></head> // html deploy
"./book/loQueSea" // html deploy relative URL
COOKIES session.go 44
//Path:  "/book/",
//Path:  baseURL + "book/", // deploy
*/

//var baseURL = "./../../" // Go local
var baseURL = "/freecodecamp/7-bonus2/" // Go deploy
var tmpl map[string]*template.Template

func init() {
	base := "book/views/"
	guest := base + "guest.html"
	logged := base + "logged.html"
	mybooks := base + "mybooks.html"
	login := base + "login.html"
	profile := base + "profile.html"
	layout := base + "partials/layout.html"
	footer := base + "partials/footer.html"
	scripts := base + "partials/scripts.html"
	books := base + "partials/books.html"
	guestNav := base + "partials/guestNav.html"
	loggedNav := base + "partials/loggedNav.html"
	addBook := base + "partials/addBook.html"
	tmpl = make(map[string]*template.Template)
	tmpl["guest.html"] = template.Must(template.ParseFiles(guest, layout, footer, scripts, guestNav))
	tmpl["logged.html"] = template.Must(template.ParseFiles(logged, layout, footer, scripts, books, loggedNav))
	tmpl["mybooks.html"] = template.Must(template.ParseFiles(mybooks, layout, footer, scripts, books, loggedNav, addBook))
	tmpl["login.html"] = template.Must(template.ParseFiles(login, layout, footer, scripts, guestNav))
	tmpl["profile.html"] = template.Must(template.ParseFiles(profile, layout, footer, scripts, loggedNav))
}

// RouterBook ...
func RouterBook(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Original Path ==> ", r.URL.Path)
	param := strings.Split(r.URL.Path, "/")
	path := param[2:len(param)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	//path2 := r.URL.Path[len("/book/"):]
	fmt.Println("Going ....", path)
	if len(path) == 0 {
		fmt.Fprintln(w, "INDEX, NOT FOUND")
		return
	}
	switch path[0] {
	case "guest":
		guest(w, r)
	case "login":
		if r.Method == "POST" {
			doLoginOrCreate(w, r)
		} else {
			login(w, r)
		}
	case "logged":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"book/login", 301)
			return
		}
		logged(w, r)
	case "mybooks":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"book/login", 301)
			return
		}
		mybooks(w, r)
	case "profile":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"book/login", 301)
			return
		}
		if r.Method == "POST" {
			dbSaveNewProfile(w, r)
		} else {
			profile(w, r)
		}
	case "logout":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"book/login", 301)
			return
		}
		logout(w, r)
	case "addBook":
		if !isLogged(r) {
			http.Redirect(w, r, baseURL+"book/login", 301)
			return
		}
		if r.Method == "POST" {
			dbAddNewBook(w, r)
		} else {
			http.Redirect(w, r, baseURL+"book/login", 301)
		}
	case "desire":
		dbSaveDesire(w, r)
	case "api":
		getBooksFromGoogle(w)
	default:
		fmt.Fprintln(w, "404 PAGE NOT FOUND")
	}
}

func guest(w http.ResponseWriter, r *http.Request) {
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	tmpl["login.html"].ExecuteTemplate(w, "login.html", nil)
}
func wrongLogin(w http.ResponseWriter, r *http.Request) {
	data := "INVALID PASSWORD"
	tmpl["login.html"].ExecuteTemplate(w, "login.html", data)
}
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`LOGOUT`)
	dbDeleteCookie(r)
	http.Redirect(w, r, baseURL+"book/guest/", 301)
}

func logged(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	data := dbGetAllBookList(user)
	data.ActiveUser = getUser(r)
	tmpl["logged.html"].ExecuteTemplate(w, "logged.html", data)
}

func profile(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	data := dbGetUserProfile(user)
	data.ActiveUser = user
	tmpl["profile.html"].ExecuteTemplate(w, "profile.html", data)
}

func mybooks(w http.ResponseWriter, r *http.Request) {
	var data bookList
	user := getUser(r)
	data = dbGetMyBookList(user)
	dbGetMyBookRequests(&data, user)
	tmpl["mybooks.html"].ExecuteTemplate(w, "mybooks.html", data)
}

func getUser(r *http.Request) string {
	r.ParseForm()
	return r.Form["user"][0]
}
