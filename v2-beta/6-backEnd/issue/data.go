package issue

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

const dataFile = "./issue/issue.json"

func init() {
}

var d dataSet

type dataSet map[string]project

var p project

type project struct {
	Issues []issue `json:"issues,omitempty"`
}

type issue struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	CreationDate string `json:"creationDate"`
	LatestUpdate string `json:"latestUpdate"`
	Author       string `json:"author"`
	Assignee     string `json:"assignee"`
	IsOpen       bool   `json:"isOpen"`
	//Status       string `json:"status"`
}

func (d *dataSet) getAllIssues(w http.ResponseWriter, r *http.Request, project string) {
	util.StructToJSON(w, r, p)
}

func (d *dataSet) getOneIssue(w http.ResponseWriter, r *http.Request, project string, id int) {
	for _, v := range p.Issues {
		if id == v.ID {
			util.StructToJSON(w, r, v)
			return
		}
	}
	e.Error = fmt.Sprintf("Issue id = %d not found", id)
	e.Info = ""
	util.StructToJSON(w, r, e)

}

func (d *dataSet) insertIssue(w http.ResponseWriter, r *http.Request, project string) {
	r.ParseForm() // or r.URL.Query().Get("title")
	var aux issue
	aux.ID = p.getFirstFreeIndex()
	aux.Title = r.Form.Get("title")
	aux.Text = r.Form.Get("text")
	aux.CreationDate = time.Now().Format(time.RFC1123Z)
	aux.LatestUpdate = aux.CreationDate
	aux.Author = r.Form.Get("createdBy")
	aux.Assignee = r.Form.Get("assignedTo")
	aux.IsOpen = true
	p.Issues = append(p.Issues, aux)
	(*d)[project] = p
	d.writeJSONtoFile()
	e.Error = ""
	e.Info = "Issue Successfully added"
	util.StructToJSON(w, r, e)
}

func (d *dataSet) updateIssue(w http.ResponseWriter, r *http.Request, projectName string) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("issueID"))
	if err != nil {
		e.Error = fmt.Sprintf("Issue with ID = %d not found", id)
		e.Info = ""
	}
	found := false
	for i, v := range p.Issues {
		if v.ID == id {
			found = true
			if r.Form.Get("title") != "" {
				(*d)[projectName].Issues[i].Title = r.Form.Get("title")
			}
			if r.Form.Get("text") != "" {
				(*d)[projectName].Issues[i].Text = r.Form.Get("text")
			}
			if r.Form.Get("createdBy") != "" {
				(*d)[projectName].Issues[i].Author = r.Form.Get("createdBy")
			}
			if r.Form.Get("assignedTo") != "" {
				(*d)[projectName].Issues[i].Assignee = r.Form.Get("assignedTo")
			}
			if r.Form.Get("close") == "true" {
				(*d)[projectName].Issues[i].IsOpen = false
			}
			(*d)[projectName].Issues[i].LatestUpdate = time.Now().Format(time.RFC1123Z)
		}
	}
	if found {
		e.Info = "Issue Successfully updated"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Issue with ID = %d not found", id)
	}
	util.StructToJSON(w, r, e)
	d.writeJSONtoFile()

}

func (d *dataSet) deleteIssue(w http.ResponseWriter, r *http.Request, projectName string) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		e.Error = fmt.Sprintf("Issue with ID = %d not found", id)
		e.Info = ""
	}
	found := false
	var aux project //[]issue
	for _, v := range p.Issues {
		if v.ID != id {
			aux.Issues = append(aux.Issues, v)
		} else {
			found = true
		}
	}
	if found {
		e.Info = "Issue Successfully deleted"
		e.Error = ""
	} else {
		e.Info = ""
		e.Error = fmt.Sprintf("Issue with ID = %d not found", id)
	}
	util.StructToJSON(w, r, e)
	(*d)[projectName] = aux
	d.writeJSONtoFile()
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

func (p *project) getFirstFreeIndex() int {
	var indexs []int
	for _, v := range p.Issues {
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
