package voting

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sumUserVote(w http.ResponseWriter, r *http.Request) {
	if !isLogged(r) {
		http.Redirect(w, r, "/voting/login/", 301)
		return
	}
	cookie, _ := r.Cookie("session")
	user := strings.Split(cookie.Value, ":")[0]
	r.ParseForm()
	choiceVoted, _ := strconv.Atoi(r.Form["options"][0])
	ipUsed := getIP(r)
	//fmt.Printf("%s Voted option %d from IP %s", user, choiceVoted, ipUsed)
	if dbAlreadyUserVoted(user, choiceVoted, ipUsed) {
		http.Redirect(w, r, "/voting/alreadyUserVoted?user="+user, 301)
		return // avoid execute next redirect
	}
	dbSumUserVote(user, choiceVoted, ipUsed)
	http.Redirect(w, r, "/voting/logged?user="+user, 301)
}

func dbSumUserVote(user string, option int, ip string) {
	db, _ := connectDB()
	_, err := db.Exec("INSERT into voting.votes (Username, ChoiceID, ip) values (?,?,?)", user, option, ip)
	if err != nil {
		log.Fatal(err)
	}
}

func dbAlreadyUserVoted(user string, option int, ip string) bool {
	//ip = "192.168.1.3"
	odds := dbGetChoicesFromOnePoll(option)
	//fmt.Println("ODDS --> ", len(odds), odds)
	var name string
	var choiceID int
	var ipdir string
	db, _ := connectDB()
	for i := 0; i < len(odds); i++ {
		row := db.QueryRow("SELECT * FROM voting.votes WHERE ChoiceID = ? AND Username = ?", odds[i], user)
		err := row.Scan(&name, &choiceID, &ipdir)
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

func dbInsertNewOption(newPoll int, newOption string) {
	fmt.Println("10-->", newPoll, newOption)
	db, _ := connectDB()
	// find out pollID from the choiceID we have
	row := db.QueryRow("SELECT PollID FROM voting.choices WHERE ChoicesID = ?", newPoll)
	var pollNumber int
	err := row.Scan(&pollNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	fmt.Println("11-->", pollNumber)
	// find out if there is newOption in the Poll
	if dbExistsOption(pollNumber, newOption) {
		return
	}
	_, err = db.Exec("INSERT into voting.choices (PollID, Choice) values (?,?)", pollNumber, newOption)
	if err != nil {
		log.Fatal(err)
	}
}

func dbExistsOption(poll int, newOption string) bool {
	db, _ := connectDB()
	row := db.QueryRow("SELECT ChoicesID FROM voting.choices WHERE PollID = ? AND Choice = ?", poll, newOption)
	var choiceID int
	err := row.Scan(&choiceID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return false
		}
		log.Fatal(err)
	}
	//fmt.Println("Existing Option")
	return true
}

func doCreateNewPoll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	question := r.Form["question"][0]
	choices := strings.Split(r.Form["choices"][0], ",")
	fmt.Println(question, choices)
	pollID := dbInsertNewPollIfNoExists(question)
	if pollID == 0 {
		return
	}
	dbInsertChoicesOnNewPoll(pollID, choices)
	user := r.Form["user"][0]
	http.Redirect(w, r, "/voting/logged?user="+user, 301)
}

func dbInsertNewPollIfNoExists(question string) int {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Exec("INSERT INTO voting.polls (Question) SELECT ? FROM dual WHERE NOT EXISTS (SELECT Question FROM voting.polls WHERE Question = ?) LIMIT 1", question, question)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	id, _ := row.LastInsertId() // if no inserts returns 0
	if id == 0 {
		return 0
	}
	return int(id)
}

func dbInsertChoicesOnNewPoll(pollID int, choices []string) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	for index := 0; index < len(choices); index++ {
		row, err := db.Exec("INSERT into voting.choices (PollID, Choice) values (?,?)", pollID, choices[index])
		if err != nil {
			log.Fatal(err)
			return
		}
		res, _ := row.RowsAffected()
		if res > 0 {
			total++
		}
	}
}
