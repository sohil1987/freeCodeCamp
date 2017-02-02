package main

import (
	"fmt"
	"net/http"
)

type outputFile struct {
	File  string `json:"file"`
	Size  int64  `json:"size"`
	chunk []byte
}

type fileinfo interface {
	Size() int64
}

func file(w http.ResponseWriter, r *http.Request) {
	var data outputFile
	fmt.Println(`DENTRO`)
	file, handler, err := r.FormFile("inputFile")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// get file content
	/*data.chunk, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintln(string(data.chunk))*/

	var fi fileinfo
	fi, _ = file.(fileinfo)
	data.File = handler.Filename
	data.Size = fi.Size()

	sendStructAsJSON(w, r, data)
}
