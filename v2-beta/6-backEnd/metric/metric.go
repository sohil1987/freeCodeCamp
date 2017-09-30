package metric

//https://forum.freecodecamp.org/t/metric-imperial-conversion-project-with-testable-user-stories-guinea-pigs-needed/72754

import (
	"fmt"
	"freeCodeCamp/v2-beta/6-backEnd/util"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	//fmt.Println("init from metric package")
}

var validLetters = "abcdefghijklmnopqrstuvwxyz"
var validNumber = "0123456789"
var units = []string{"lbs", "kg", "gal", "l", "mi", "km"}
var unitsWords = map[string]string{"lbs": "pounds", "kg": "kilograms", "gal": "gallons", "l": "liters", "mi": "miles", "km": "kilometers"}

var o output

type output struct {
	InitNum    float64
	InitUnit   string
	ReturnNum  float64
	ReturnUnit string
	Text       string
}

var e errorOutput

type errorOutput struct {
	Text string
}

// RouterMetric ...
func RouterMetric(w http.ResponseWriter, r *http.Request) {
	o = output{}
	e = errorOutput{}
	params := strings.Split(r.URL.Path, "/")
	path := params[3:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	fmt.Println("Going ....", path)
	if len(path) == 0 || len(path) > 1 {
		util.PageNotFound(w, r)
		return
	}
	if path[0] != "convert" {
		util.PageNotFound(w, r)
		return
	}
	r.ParseForm()
	input := r.Form.Get("input")
	parseInput(strings.ToLower(input))
	if e.Text != "" {
		util.StructToJSON(w, r, e)
		return
	}
	calculateReturns()
	util.StructToJSON(w, r, o)
}

func calculateReturns() {
	switch o.InitUnit {
	case "lbs":
		o.ReturnUnit = "kg"
		o.ReturnNum = o.InitNum * 0.453592
	case "kg":
		o.ReturnUnit = "lbs"
		o.ReturnNum = o.InitNum * (1 / 0.453592)
	case "gal":
		o.ReturnUnit = "l"
		o.ReturnNum = o.InitNum * 3.78541
	case "l":
		o.ReturnUnit = "gal"
		o.ReturnNum = o.InitNum * (1 / 3.78541)
	case "mi":
		o.ReturnUnit = "km"
		o.ReturnNum = o.InitNum * 1.60934
	case "km":
		o.ReturnUnit = "mi"
		o.ReturnNum = o.InitNum * (1 / 1.60934)
	}
	o.InitNum = util.ToFixedFloat64(o.InitNum, 5)
	o.ReturnNum = util.ToFixedFloat64(o.ReturnNum, 5)
	o.Text = fmt.Sprintf("%g %s converts to %g %s", o.InitNum, unitsWords[o.InitUnit], o.ReturnNum, unitsWords[o.ReturnUnit])
}

func parseInput(input string) {
	var part1, part2 string
	var found = false
	for i, char := range input {
		if strings.Contains(validLetters, string(char)) && !found {
			part1 = input[0:i]
			part2 = input[i:len(input)]
			found = true
		}
	}
	// get initUnit
	for _, unit := range units {
		if unit == part2 {
			o.InitUnit = unit
		}
	}
	if o.InitUnit == "" {
		e.Text = "invalid unit"
	}
	// get initNum
	if part1 == "" && e.Text == "" {
		o.InitNum = 1
		return
	}
	ops := make([]float64, 0)
	numbers := strings.Split(part1, "/")
	for _, number := range numbers {
		aux, err := strconv.ParseFloat(number, 64)
		if err != nil {
			if e.Text == "" {
				e.Text = "invalid number"
			} else {
				e.Text = "invalid number and unit"
			}
		}
		ops = append(ops, aux)
	}
	o.InitNum = ops[0]
	if len(ops) > 1 {
		for i, op := range ops {
			if i > 0 {
				o.InitNum = o.InitNum / op
			}
		}
	}
	if o.InitNum < 0 {
		if e.Text == "" {
			e.Text = "invalid number"
		} else {
			e.Text = "invalid number and unit"
		}
	}
}
