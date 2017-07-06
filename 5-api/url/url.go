package url

import (
	"encoding/json"
	"freeCodeCamp/5-api/_help"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

func init() {
	//fmt.Println("Init from package url")
}

// RouterURL ...
func RouterURL(w http.ResponseWriter, r *http.Request) {
	option := strings.Split(r.URL.Path, "/")[3]
	switch strings.ToLower(option) {
	case "list":
		//fmt.Println(`list`)
		getList(w, r)
	case "new":
		//fmt.Println(`new`)
		getNewShortenedURL(w, r)
	default:
		//fmt.Println(`default`)
		getRedirectToShortenedURL(w, r)
	}
}

const filePath = "./url/url.json"
const filePath2 = "./url/url2.json"

type invalid struct {
	Url   string `json:"url"`
	Error string `json:"error"`
}

type data []record

type record struct {
	Short    int    `json:"short"`
	Original string `json:"original"`
}

func getList(w http.ResponseWriter, r *http.Request) {
	var d data
	getDataFromFile(filePath, &d)
	help.StructToJSON(w, r, d)
}

func getNewShortenedURL(w http.ResponseWriter, r *http.Request) {
	candidate := strings.Replace(r.URL.Path, "/url/v1/new/", "", 1)
	candidate = strings.Replace(candidate, "/", "//", 1)
	candidate = strings.ToLower(candidate)
	var i invalid
	i.Url = candidate
	if !isValidURL(candidate) {
		i.Error = candidate + " is not a valid format URL"
		help.StructToJSON(w, r, i)
		return
	} else if !checkURLisUp(candidate) {
		i.Error = "Can't DNS to URL " + candidate
		help.StructToJSON(w, r, i)
		return
	}
	var d data
	getDataFromFile(filePath, &d)
	if candidateExists(candidate, d) {
		i.Error = candidate + " already exists"
		help.StructToJSON(w, r, i)
		return
	}
	var rec record
	rec.Original = candidate
	// Take the first free number
	indexs := make([]int, 0) // take all indexs
	for i := range d {
		indexs = append(indexs, d[i].Short)
	}
	sort.Ints(indexs) // sort
	found := false    // look for the first free
	cont := 1
	for !found && cont < 1000 {
		if cont != indexs[cont-1] {
			rec.Short = cont
			found = true
		}
		cont++
	}
	d = append(d, rec)
	help.StructToJSON(w, r, rec)
	writeJSONtoFile(&d, filePath)
}

func getRedirectToShortenedURL(w http.ResponseWriter, r *http.Request) {
	option := strings.Split(r.URL.Path, "/")[3]
	id, err := strconv.Atoi(option)
	if err != nil { // param is not a valid number
		var i invalid
		i.Url = option
		i.Error = option + " is not a valid number."
		help.StructToJSON(w, r, i)
		return
	}
	var d data
	getDataFromFile(filePath, &d)
	for _, v := range d {
		if id == v.Short {
			http.Redirect(w, r, v.Original, 301)
			//return
		}
	}
	var i invalid
	i.Url = option
	i.Error = "Doesn't exist a shortened URL for number " + option
	help.StructToJSON(w, r, i)
}

func getDataFromFile(pathToFile string, d *data) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file) //	get file content
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeJSONtoFile(d *data, pathToFile string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(d)
}

func candidateExists(candidate string, d data) bool {
	for _, v := range d {
		if candidate == string(v.Original) {
			return true
		}
	}
	return false
}

func checkURLisUp(url string) bool {
	res, err := http.Get(url)
	if err == nil && res.StatusCode == 200 {
		return true
		//fmt.Println(url, `UP`)
	}
	return false
	//fmt.Println(url, "DOWN\n", err)
}

func isValidURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	return true
}
