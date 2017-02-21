package stock

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getInitialCode(w http.ResponseWriter) {
	list := readFile(c.FileDB)
	s := strings.Split(list, "\n")
	sendStructAsJSON(w, s)
}

func addCode(w http.ResponseWriter, r *http.Request, code string) {
	addCodeToFile(code)
	getInitialCode(w)
}

func delCode(w http.ResponseWriter, r *http.Request, code string) {
	delCodeFromFile(code)
	getInitialCode(w)
}

func delCodeFromFile(del string) {
	var newContent = readFile(c.FileDB)
	file, err := os.Create(c.FileDB)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	list := strings.Split(newContent, "\n")
	for k, v := range list {
		if v == strings.ToUpper(del) {
			list = append(list[:k], list[k+1:]...)
			file.WriteString(strings.Join(list, "\n"))
			return
		}
	}
	file.WriteString(newContent) // here if del not in slice
}

func codeExist(add string) bool {
	urlTest := "https://www.quandl.com/api/v3/datasets/WIKI/" + add + "/metadata.json"
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(urlTest + "?api_key=" + c.APIKey)
	//resp, err := http.Get(urlTest + "?api_key=" + c.APIKey)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func addCodeToFile(add string) {
	var newContent = readFile(c.FileDB)
	file, err := os.Create(c.FileDB)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if sliceContainsString(strings.ToUpper(add), strings.Split(newContent, "\n")) {
		file.WriteString(newContent)
		return
	}
	if newContent == "" { // avoid insert white line when file is empty
		file.WriteString(strings.ToUpper(add))
	} else {
		file.WriteString(newContent + "\n" + strings.ToUpper(add))
	}
}

func readFile(str string) string {
	file, err := os.Open(str)
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	data := string(bs)
	return data
}

func sliceContainsString(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
