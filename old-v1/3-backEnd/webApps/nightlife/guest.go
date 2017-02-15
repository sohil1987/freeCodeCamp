package nightlife

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func searchBarInLocation(location, user string) dataAPI {
	base := "https://api.yelp.com/v3/businesses/search?term=bar&location="
	baseurl := base + location
	help := fmt.Sprintf("Bearer %s", c.Yelp.Token)
	//fmt.Println(help)
	req, err := http.NewRequest("GET", baseurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", help)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var dataYelp dataAPI
	/* Option 1 */
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&dataYelp)
	if err != nil {
		panic(err)
	}
	/* Option 2 */
	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &dataYelp)*/

	// print json response pretty formatted
	//dataJSON, _ := json.MarshalIndent(dataYelp, "", " ")
	//fmt.Println(string(dataJSON))

	for _, v := range dataYelp.Businesses {
		v.Going = dbGetPeopleGoingABar(v.ID)
		//fmt.Println( v.ID, "-->", v.Going)
	}

	return dataYelp
}

func searchBarInLocation2() dataAPI {
	d := loadDataDev()
	return d
}

func dbGetPeopleGoingABar(id string) int {
	var going int
	db, err := connectDB()
	row := db.QueryRow("SELECT COUNT(*) as count FROM  nightlife.votes WHERE BarID=?", id)
	defer db.Close()
	err = row.Scan(&going)
	if err != nil {
		log.Fatal(err)
	}
	return going
}
