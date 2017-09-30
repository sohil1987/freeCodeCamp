package files

import (
	"bytes"
	"fmt"
	"freeCodeCamp/v1/back-end/apis/_help"
	"net/http"
	"strconv"
)

func init() {
	//fmt.Println(`Init from Files Package`)
}

// Output ...
type Output struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

// RouterFiles ...
func RouterFiles(w http.ResponseWriter, r *http.Request) {
	var o Output

	file, handler, err := r.FormFile("inputFile")
	if err != nil {
		http.Redirect(w, r, help.BaseURL+"file/file.html", 301)
		return
	}
	defer file.Close()

	var buff bytes.Buffer
	size, err := buff.ReadFrom(file)
	if err != nil {
		http.Redirect(w, r, help.BaseURL+"file/file.html", 301)
		return
	}

	o.Filename = handler.Filename
	o.Size = getCapacity(float64(size))
	help.StructToJSON(w, r, o)

}

func getCapacity(size float64) string {
	var q string
	if size < 1000 {
		q = strconv.FormatFloat(size, 'f', 2, 64)
		fmt.Println(q)
		return q + " B"
	} else if size < 1000000 {
		q = strconv.FormatFloat(size/1000, 'f', 2, 64)
		fmt.Println(q)
		return q + " Kb"
	} else {
		q = strconv.FormatFloat(size/1000000, 'f', 2, 64)
		fmt.Println(q)
		return q + " Mb"
	}
}
