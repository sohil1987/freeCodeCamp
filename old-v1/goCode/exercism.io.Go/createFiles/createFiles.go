package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func createFiles() {
	nom := readFiles("./onlyNames.txt")
	fil := readFiles("./fileNames.txt")
	fmt.Println(len(strings.Split(nom, "\n")))
	fmt.Println(len(strings.Split(fil, "\n")))
	names := strings.Split(nom, "\n")
	files := strings.Split(fil, "\n")
	files = files[:len(files)-1] // remove last empty element
	for i, v := range files {
		fmt.Println(names[i], v)
		// DANGER, BE CAREFUL
		//writeFile(names[i], v)
		//
	}
}

func writeFile(name, fileName string) {
	var s = `
package main

import (
	"fmt"
)

func ` + name + `() {
	fmt.Println("` + name + `")
}`
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(s)
}

func readFiles(str string) string {
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
