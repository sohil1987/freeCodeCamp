package pintelest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func init() {
}

var imgs Images

// Images ...
type Images []Image

// Image ...
type Image struct {
	IDPic      int    `json:"idPic"`
	Link       string `json:"link"`
	IDAuthor   string `json:"idTwitter"`
	LogoAuthor string `json:"logoAuthor"`
	NameAuthor string `json:"username"`
	Text       string `json:"text"`
	Likes      int    `json:"likes"`
}

func (imgs *Images) loadData(filePath string) { // parse JSON with MARSHALL
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	// get file content
	chunk, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(chunk, &imgs)
	fmt.Println(reflect.TypeOf(imgs))
}
