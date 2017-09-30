package message

import (
	"encoding/json"
	"fmt"
	"freeCodeCamp/v2-beta/6-backEnd/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const dataFile = "./message/message.json"

func init() {
}

var d dataSet
var h thread
var m reply

type dataSet []thread

type thread struct {
	Board        int     `json:"board"`
	Title        string  `json:"title"`
	Text         string  `json:"text"`
	CreationDate string  `json:"creationDate"`
	LastPost     string  `json:"lastPost"`
	Password     string  `json:"password"`
	IsReported   bool    `json:"isReported"`
	Replies      []reply `json:"replies"`
}

type reply struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	Password   string `json:"password"`
	IsReported bool   `json:"isReported"`
}

func (d *dataSet) getThreads(w http.ResponseWriter, r *http.Request) {
	util.StructToJSON(w, r, d)
}

func (d *dataSet) createThread(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := d.getFirstFreeThreadIndex()
	var aux thread
	aux.Board = index
	aux.CreationDate = time.Now().Format(time.RFC1123Z)
	aux.IsReported = false
	aux.LastPost = time.Now().Format(time.RFC1123Z)
	aux.Replies = []reply{}
	aux.Text = r.Form.Get("text")
	aux.Title = r.Form.Get("title")
	if r.Form.Get("password") != "" {
		aux.Password = r.Form.Get("password")
	}
	*d = append(*d, aux)
	d.writeJSONtoFile()
	e.Info = "Thread successfully created"
	e.Error = ""
	util.StructToJSON(w, r, e)
}

func (d *dataSet) reportThread(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	found := false
	board, err := strconv.Atoi(r.Form.Get("board"))
	if err != nil {
		e.Info = ""
		e.Error = fmt.Sprintf("id %s is not a valid ID for a new board", r.Form.Get("board"))
		util.StructToJSON(w, r, e)
		return
	}
	for i, v := range *d {
		if v.Board == board {
			found = true
			(*d)[i].IsReported = true
			d.writeJSONtoFile()
		}
	}
	if found {
		e.Info = "Thread successfully reported"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Thread num %d not exists", board)
	}
	util.StructToJSON(w, r, e)
}

func (d *dataSet) deleteThread(w http.ResponseWriter, r *http.Request) {
	found := false
	r.ParseForm()
	board, err := strconv.Atoi(r.Form.Get("board"))
	if err != nil {
		e.Info = ""
		e.Error = fmt.Sprintf("id %s is not a valid ID for a new board", r.Form.Get("board"))
		util.StructToJSON(w, r, e)
		return
	}
	have := r.Form.Get("password")
	for i, v := range *d {
		if v.Board == board {
			found = true
			need := v.Password
			if need == have {
				aux := *d
				aux = append(aux[:i], aux[i+1:]...)
				*d = aux
			} else {
				e.Info = ""
				e.Error = "Incorrect Password"
				util.StructToJSON(w, r, e)
				return
			}
			d.writeJSONtoFile()
		}
	}
	if found {
		e.Info = "Thread successfully deleted"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Thread num %d not exists", board)
	}
	util.StructToJSON(w, r, e)
}

func (d *dataSet) getReplies(w http.ResponseWriter, r *http.Request, b int) {
	for _, v := range *d {
		if v.Board == b {
			util.StructToJSON(w, r, v)
			return
		}
	}
	e.Info = ""
	e.Error = fmt.Sprintf("Board num %d not found", b)
	util.StructToJSON(w, r, e)
}

func (d *dataSet) createReply(w http.ResponseWriter, r *http.Request, b int) {
	r.ParseForm()
	found := false
	for i, v := range *d {
		if v.Board == b {
			found = true
			(*d)[i].LastPost = time.Now().Format(time.RFC1123Z)
			var aux reply
			aux.ID = d.getFirstFreeReplyIndex()
			aux.IsReported = false
			aux.Text = r.Form.Get("text")
			if r.Form.Get("password") != "" {
				aux.Password = r.Form.Get("password")
			}
			(*d)[i].Replies = append((*d)[i].Replies, aux)
			d.writeJSONtoFile()
		}
	}
	if found {
		e.Info = "Reply successfully created"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Board num %d not exists", b)
	}
	util.StructToJSON(w, r, e)
}

func (d *dataSet) reportReply(w http.ResponseWriter, r *http.Request, b int) {
	found := false
	for i, v := range *d {
		for j, t := range v.Replies {
			if t.ID == b {
				found = true
				(*d)[i].Replies[j].IsReported = true
				d.writeJSONtoFile()
			}
		}
	}
	if found {
		e.Info = "Reply successfully reported"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Reply num %d not exists", b)
	}
	util.StructToJSON(w, r, e)
}

func (d *dataSet) deleteReply(w http.ResponseWriter, r *http.Request, b int) {
	found := false
	r.ParseForm()
	have := r.Form.Get("password")
	for i, v := range *d {
		for j, t := range v.Replies {
			fmt.Printf("board = %d , reply = %d\n", v.Board, t.ID)
			if t.ID == b {
				found = true
				need := t.Password
				if need == have {
					aux := (*d)[i].Replies
					aux = append(aux[:j], aux[j+1:]...)
					(*d)[i].Replies = aux
				} else {
					e.Info = ""
					e.Error = "Incorrect Password"
					util.StructToJSON(w, r, e)
					return
				}
				d.writeJSONtoFile()
				fmt.Sprintln(i, j)
			}
		}
	}
	if found {
		e.Info = "Reply successfully deleted"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Reply num %d not exists", b)
	}
	util.StructToJSON(w, r, e)
}

func (d *dataSet) loadJSONDataFromFile() {
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
	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Fatalln(err)
	}
}

func (d *dataSet) writeJSONtoFile() {
	f, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(&d)
}

func (d *dataSet) getFirstFreeThreadIndex() int {
	var indexs []int
	for _, v := range *d {
		fmt.Println(v.Board)
		indexs = append(indexs, v.Board)
	}
	for i := 1; i < len(indexs)+2; i++ {
		if !contains(indexs, i) {
			return i
		}
	}
	return len(indexs) + 1
}

func (d *dataSet) getFirstFreeReplyIndex() int {
	var indexs []int
	for _, v := range *d {
		for _, t := range v.Replies {
			indexs = append(indexs, t.ID)
		}
	}
	for i := 1; i < len(indexs)+2; i++ {
		if !contains(indexs, i) {
			return i
		}
	}
	return len(indexs) + 1
}

func contains(indexs []int, num int) bool {
	for _, index := range indexs {
		if index == num {
			return true
		}
	}
	return false
}
