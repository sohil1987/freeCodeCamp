package pintelest

import (
	"fmt"
	"freeCodeCamp/7-bonus2/_help"
	"html/template"
	"net/http"
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

const dataFile = "./pintelest/search.json" // "./pintelest/data.json"
const secretFile = "./secret.json"
const loginCallback = "http://brusbilis.com/freecodecamp/7-bonus2/pintelest/twitter"

var tmpl map[string]*template.Template

func init() {
	var base = "pintelest/views/"
	addPic := base + "addPic.html"
	guest := base + "guest.html"
	myPics := base + "myPics.html"
	pics := base + "pics.html"
	profile := base + "profile.html"
	head := base + "partials/head.html"
	end := base + "partials/end.html"
	navbarLogged := base + "partials/navbarLogged.html"
	navbar := base + "partials/navbar.html"
	tmpl = make(map[string]*template.Template)
	tmpl["guest.html"] = template.Must(template.ParseFiles(navbar, head, end, guest))
	tmpl["profile.html"] = template.Must(template.ParseFiles(navbarLogged, head, end, profile))
	tmpl["pics.html"] = template.Must(template.ParseFiles(navbarLogged, head, end, pics))
	tmpl["myPics.html"] = template.Must(template.ParseFiles(navbarLogged, head, end, myPics))
	tmpl["addPic.html"] = template.Must(template.ParseFiles(navbarLogged, head, end, addPic))

}

// RouterPintelest ...
func RouterPintelest(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	path := params[2:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Printf("Going ..... %s logueado como %v \n", path, c.User.Name)
	if len(path) == 0 {
		http.Redirect(w, r, "https://brusbilis.com/freecodecamp/7-bonus2/pintelest/guest/", 301)
		return
	}
	switch path[0] {
	case "guest":
		if !c.isLogged() {
			mydb.dbGetAllPics()
			guest(w, r)
			return
		}
		pageNotFound(w, r)
	case "login":
		if !c.isLogged() {
			login(w, r)
			return
		}
		pageNotFound(w, r)
	case "twitter":
		if !c.isLogged() {
			callback(w, r)
			return
		}
		pageNotFound(w, r)
	case "logout":
		if c.isLogged() {
			logout(w, r)
			return
		}
		pageNotFound(w, r)
	case "profile":
		if c.isLogged() {
			profile(w, r)
			return
		}
		pageNotFound(w, r)
	case "pics":
		mydb.dbGetAllPics()
		if c.isLogged() {
			pics(w, r)
			return
		}
		pageNotFound(w, r)
	case "myPics":
		mydb.dbGetUserPics(c.User.UserID)
		if c.isLogged() {
			myPics(w, r)
			return
		}
		pageNotFound(w, r)
	case "addPic":
		if c.isLogged() {
			if r.Method == "POST" {
				addPicToDB(w, r)
			} else {
				addPic(w, r)
			}
			return
		}
		pageNotFound(w, r)
	case "user":
		mydb.dbGetUserPics(path[1])
		if c.isLogged() {
			pics(w, r)
		} else {
			guest(w, r)
		}
		pageNotFound(w, r)
	case "delete":
		if len(path) >= 2 { // avoid delete without id param
			if c.isLogged() && path[1] != "" {
				mydb.deletePic(path[1])
				mydb.dbGetUserPics(c.User.UserID)
				myPics(w, r)
				return
			}
			return
		}
		pageNotFound(w, r)
	case "vote":
		if len(path) >= 2 {
			if c.isLogged() && path[1] != "" {
				mydb.vote(c.User.UserID, path[1])
				mydb.dbGetAllPics()
				pics(w, r)
				//http.Redirect(w, r, help.BaseURL+"pintelest/pics", 301)
				return
			}
			return
		}
		pageNotFound(w, r)
	default:
		pageNotFound(w, r)
	}
}

func addPicToDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newLink := r.Form.Get("linkPic")
	newText := r.Form.Get("textPic")
	mydb.insertUser(newText, newLink)
	http.Redirect(w, r, "https://brusbilis.com/freecodecamp/7-bonus2/pintelest/myPics", 301)
}

func guest(w http.ResponseWriter, r *http.Request) {
	tmpl["guest.html"].ExecuteTemplate(w, "guest.html", &imgs)
}

func profile(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"imgs": imgs,
		"conf": c,
	}
	tmpl["profile.html"].ExecuteTemplate(w, "profile.html", &m)

}

func pics(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"imgs": imgs,
		"conf": c,
	}
	tmpl["pics.html"].ExecuteTemplate(w, "pics.html", &m)
}

func myPics(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`MYPICS`)
	m := map[string]interface{}{
		"imgs": imgs,
		"conf": c,
	}
	tmpl["myPics.html"].ExecuteTemplate(w, "myPics.html", &m)
}

func addPic(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"imgs": imgs,
		"conf": c,
	}
	tmpl["addPic.html"].ExecuteTemplate(w, "addPic.html", &m)
}

func login(w http.ResponseWriter, r *http.Request) {
	c.authTwitter(w, r)
}

func callback(w http.ResponseWriter, r *http.Request) {
	c.callbackTwitter(w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c.logoutTwitter(w, r)
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, help.BaseURL+"error/404.html", 301)
}
