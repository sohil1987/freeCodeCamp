package issue

import (
	"fmt"
	"freeCodeCamp/v2-beta/6-backEnd/util"
	"net/http"
	"strconv"
	"strings"
)

func init() {

}

var e myerror

type myerror struct {
	Info  string `json:"info"`
	Error string `json:"error"`
}

// RouterIssue ...
func RouterIssue(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	path := params[3:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Println("Going ....", path, len(path))
	if len(path) == 0 || len(path) > 2 {
		util.PageNotFound(w, r)
		return
	}
	d.loadJSONDataFromFile()
	project := strings.ToLower(path[0])
	var existsProject bool
	p, existsProject = d[project]
	if r.Method == "GET" {
		if existsProject {
			if len(path) == 2 {
				issue, err := strconv.Atoi(path[1])
				if err == nil {
					d.getOneIssue(w, r, project, issue)
					return
				}
				fmt.Println(err)
			}
			d.getAllIssues(w, r, project)
			return
		}
	}
	if r.Method == "POST" {
		if existsProject {
			d.insertIssue(w, r, project)
			return
		}
	}
	if r.Method == "PUT" {
		if existsProject {
			d.updateIssue(w, r, project)
			return
		}
	}
	if r.Method == "DELETE" {
		if existsProject {
			d.deleteIssue(w, r, project)
			return
		}
	}
	util.PageNotFound(w, r)
}
