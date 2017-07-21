package library

import (
	"encoding/json"
	"fmt"
	"freeCodeCamp/6-backEnd/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const dataFile = "./library/library.json"

func init() {}

var b books

type books []book

type comments []string

type book struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Comments     comments `json:"comments,omitempty"`
	CommentCount int      `json:"commentCount,omitempty"`
}

func (b *books) loadJSONDataFromFile() {
	file, err := os.Open(dataFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Fatalln(err)
	}
}

func (b *books) writeJSONtoFile() {
	f, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(&b)
}

func (b *books) getAll(w http.ResponseWriter, r *http.Request) {
	var records books
	var aux book
	for _, v := range *b {
		aux = book{}
		aux.ID = v.ID
		aux.Title = v.Title
		aux.Comments = []string{}
		aux.CommentCount = 0
		records = append(records, aux)
	}
	util.StructToJSON(w, r, records)
}

func (b *books) getOne(w http.ResponseWriter, r *http.Request, id int) {
	var aux book
	found := false
	for _, v := range *b {
		if v.ID == id {
			aux = v
			found = true
		}
	}
	if found {
		util.StructToJSON(w, r, aux)
	} else {
		e.Error = fmt.Sprintf("Book with ID = %d not found", id)
		util.StructToJSON(w, r, e)
	}
}

func (b *books) addBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var aux book
	aux.ID = b.getFirstFreeIndex()
	aux.Title = r.Form.Get("title")
	aux.Comments = []string{}
	aux.CommentCount = len(aux.Comments)
	*b = append(*b, aux)
	b.writeJSONtoFile()
}

func (b *books) addComment(w http.ResponseWriter, r *http.Request, id int) {
	r.ParseForm()
	found := false
	for i, v := range *b {
		if v.ID == id {
			(*b)[i].CommentCount++
			(*b)[i].Comments = append(v.Comments, r.Form.Get("comment"))
			fmt.Println(b)
			b.writeJSONtoFile()
			found = true
		}
	}
	if found {
		e.Info = "Book Successfully added"
		util.StructToJSON(w, r, e)
	} else {
		e.Error = fmt.Sprintf("Book with ID = %d not found", id)
		util.StructToJSON(w, r, e)
	}
}

func (b *books) deleteAll(w http.ResponseWriter, r *http.Request) {
	fixed := (*b)[0]
	(*b) = books{}
	*b = append(*b, fixed)
	b.writeJSONtoFile()
}

func (b *books) deleteOne(w http.ResponseWriter, r *http.Request, id int) {
	var fixed books
	for _, v := range *b {
		if v.ID != id {
			fixed = append(fixed, v)
		}
	}
	fixed.writeJSONtoFile()
}

func (b *books) getFirstFreeIndex() int {
	var indexs []int
	for _, v := range *b {
		indexs = append(indexs, v.ID)
	}
	for i := 0; i < len(indexs)+1; i++ {
		if !contains(indexs, i) {
			return i
		}
	}
	return -1
}

func contains(indexs []int, num int) bool {
	for _, index := range indexs {
		if index == num {
			return true
		}
	}
	return false
}
