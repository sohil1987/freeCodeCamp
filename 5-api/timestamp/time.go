package timestamp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	//fmt.Println(`Init from Package Timestamp`)
}

// Output ...
type Output struct {
	Unix    int64 `json:"unix"`
	utc     time.Time
	Natural string `json:"utc"`
}

// RouterTime ...
func RouterTime(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	input := params[3]
	o := Output{Unix: 0, Natural: "Invalid Date"}
	if input != "" {
		unix, err := strconv.Atoi(input)
		if err == nil { // we receive a unixTime
			getOutputFromUnixTime(unix, &o)
		} else { // we receive a date string
			getOutputFromDate(input, &o)
		}
	} else { // empty string, we get current timestamp
		getOutputFromUnixTime(int(time.Now().Unix()*1000), &o)
	}
	structToJSON(w, r, o)
}

func getOutputFromUnixTime(unix int, o *Output) {
	o.Unix = int64(unix)
	o.utc = time.Unix(int64(unix/1000), 0)
	o.Natural = o.utc.Format(time.RFC1123Z)
}

func getOutputFromDate(s string, o *Output) {
	layout1 := "2006 1 2"
	layout2 := "2006-1-2"
	var err error
	valid := true
	o.utc, err = time.Parse(layout1, s)
	if err != nil {
		o.utc, err = time.Parse(layout2, s)
		if err != nil { // not-valid-date
			//log.Fatal(err)
			valid = false
		}
	}
	if valid {
		o.Unix = o.utc.Unix() * 1000
		o.Natural = o.utc.Format(time.RFC1123Z)
	}
}

func structToJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}
