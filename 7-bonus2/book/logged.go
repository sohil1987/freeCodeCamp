package book

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func dbGetAllBookList(user string) bookList {
	//rows, err := db.Query("SELECT * FROM book.books")
	rows, err := db.Query("SELECT * FROM book.books WHERE owner<>?", user)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var bl bookList
	var u unit
	for rows.Next() {
		rows.Scan(&u.BookID, &u.Title, &u.Thumbnail, &u.Owner)
		bl.Units = append(bl.Units, u)
	}
	//fmt.Println(bl.Units[len(bl.Units)-1].Title)  // title of last
	return bl
}

func dbGetMyBookList(user string) bookList {
	rows, err := db.Query("SELECT * FROM book.books WHERE owner = ?", user)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var bl bookList
	var u unit
	for rows.Next() {
		rows.Scan(&u.BookID, &u.Title, &u.Thumbnail, &u.Owner)
		bl.Units = append(bl.Units, u)
	}
	//fmt.Println(bl)
	//fmt.Println(bl.Units[len(bl.Units)-1].Title)  // title of last
	bl.ActiveUser = user
	return bl
}

/*
type bookList struct {
	Units      []unit
	ActiveUser string
	MyRequests   []request
	ForMeRequests[]request
}

type unit struct {
	BookID    int
	Title     string
	Thumbnail string
	Owner     string
}

type request struct {
	Candidate string
	BookID    int
	Title     string
}

*/

func dbGetMyBookRequests(bl *bookList, user string) {
	var rs []request
	var r request
	list := dbGetBooks()
	//fmt.Println(list)

	// FOR ME
	rows, err := db.Query("SELECT * FROM book.desires WHERE BookID IN (SELECT bookID FROM book.books WHERE owner = ?)", user)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		//rows.Scan(&r.BookID)
		rows.Scan(&r.BookID, &r.Candidate)
		r.Title = list[r.BookID]
		rs = append(rs, r)
	}
	bl.ForMeRequests = rs

	rs = []request{}
	// MY REQUESTS
	rows, err = db.Query("SELECT * FROM book.desires  WHERE candidate = ?", user)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&r.BookID, &r.Candidate)
		r.Title = list[r.BookID]
		rs = append(rs, r)
	}
	bl.MyRequests = rs
}

func dbGetBooks() map[int]string {
	rows, err := db.Query("SELECT bookID, title FROM book.books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	list := make(map[int]string)
	var k int
	var v string
	for rows.Next() {
		rows.Scan(&k, &v)
		list[k] = v
	}
	return list
}

func dbAddNewBook(w http.ResponseWriter, r *http.Request) {
	// check if thumbnail is a valid URL
	r.ParseForm()
	title := r.Form["title"][0]
	thumbnail := r.Form["imgurl"][0]
	owner := r.Form["user"][0]
	if !isValidURL(thumbnail) {
		http.Redirect(w, r, baseURL+"book/mybooks/?user="+owner, 301)
		return
	}
	fmt.Sprintln(title, thumbnail, owner)
	_, err := db.Exec("INSERT INTO book.books (title,thumbnail,owner) values (?,?,?)", title, thumbnail, owner)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, baseURL+"book/mybooks/?user="+owner, 301)
}

func dbSaveDesire(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookID := r.Form["bookID"][0]
	candidate := r.Form["user"][0]
	_, err := db.Exec("INSERT INTO book.desires (bookID, candidate) values (?,?)", bookID, candidate)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, baseURL+"book/mybooks/?user="+candidate, 301)

}

func isValidURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	return true
}
