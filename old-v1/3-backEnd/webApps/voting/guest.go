package voting

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func sumGuestVote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	choiceVoted, _ := strconv.Atoi(r.Form["options"][0])
	ipUsed := getIP(r)
	//fmt.Printf("Voted option %d from IP %s", choiceVoted, ipUsed)
	if dbAlreadyGuestVoted(choiceVoted, ipUsed) {
		http.Redirect(w, r, baseURL+"voting/alreadyGuestVoted/", 301)
		return // avoid execute next redirect
	}
	dbSumGuestVote(choiceVoted, ipUsed)
	http.Redirect(w, r, baseURL+"voting/guest/", 301)
}

func dbSumGuestVote(option int, ip string) {
	db, _ := connectDB()
	_, err := db.Exec("INSERT into voting.votes (Username, ChoiceID, ip) values (?,?,?)", "guest", option, ip)
	if err != nil {
		log.Fatal(err)
	}
}

func dbAlreadyGuestVoted(option int, ip string) bool {
	//fmt.Println(`IP == >`, ip)
	//ip = "192.168.1.3"
	odds := dbGetChoicesFromOnePoll(option)
	//fmt.Println("ODDS --> ", len(odds), odds)
	var user string
	var choiceID int
	var ipdir string
	db, _ := connectDB()
	for i := 0; i < len(odds); i++ {
		row := db.QueryRow("SELECT * FROM voting.votes WHERE ChoiceID = ? AND IP = ?", odds[i], ip)
		err := row.Scan(&user, &choiceID, &ipdir)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No Records Found")
			} else {
				log.Fatal(err)
			}
		} else {
			return true // exists vote
		}
	}
	return false
}

func dbGetChoicesFromOnePoll(choiceID int) []int {
	var result []int
	var choiceIDfromTable int
	db, _ := connectDB()
	rows, err := db.Query("SELECT ChoicesID FROM voting.choices WHERE PollID = (SELECT PollID from voting.choices WHERE ChoicesID = ?)", choiceID)
	if err != nil {
		fmt.Println("No Records Found")
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&choiceIDfromTable)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(" +1 ", choiceIDfromTable)
		result = append(result, choiceIDfromTable)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
