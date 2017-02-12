package voting

import (
	"database/sql"
	"fmt"
	"log"
)

// DATABASE TYPES AND METHODS

type poll struct {
	PollID   int
	Question string
}

type choice struct {
	ChoiceID int
	PollID   int
	Choice   string
}

type vote struct {
	ChoiceID int
	Username string
	IP       string
}

type user struct {
	UserID   int
	Username string
	PassSalt []byte
	PassMd5  []byte
}

// FRONTEND

type option struct {
	OptionID int
	Option   string
	NumVotes int
}

type aPoll struct {
	PollID   int
	Question string
	Options  []option
}

var listPolls []aPoll

func dbGetGuestListPolls() []aPoll {
	polls := dbGetPolls()
	choices := dbGetChoices()
	votes := dbGetVotes()
	listPolls := make([]aPoll, len(polls))
	var o option
	for i, v := range polls { // PollID, Question
		listPolls[i].PollID = v.PollID
		listPolls[i].Question = v.Question
		o = option{NumVotes: 0}
		for j, w := range choices { // ChoiceID, PollID, Choice
			if w.PollID == v.PollID {
				o.OptionID = w.ChoiceID
				o.Option = w.Choice
				for k, x := range votes { // Username, ChoiceID, IP
					if w.ChoiceID == x.ChoiceID {
						fmt.Sprintln(i, v, j, w, k, x)
						o.NumVotes++
					}
				}
				listPolls[i].Options = append(listPolls[i].Options, o)
				o = option{NumVotes: 0}
			}
		}
	}
	//fmt.Println("LISTPOLLS --> ", listPolls)
	return listPolls
}

func dbGetVotes() []*vote {
	db, _ := connectDB()
	rows, err := db.Query("SELECT * FROM voting.votes")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return nil
		}
	}
	defer rows.Close()
	var votes []*vote
	for rows.Next() {
		v := &vote{}
		err := rows.Scan(&v.Username, &v.ChoiceID, &v.IP)
		if err != nil {
			log.Fatal(err)
		}
		votes = append(votes, v)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(`VOTES ==>`, votes)
	return votes
}

func dbGetChoices() []*choice {
	db, _ := connectDB()
	rows, err := db.Query("SELECT * FROM voting.choices")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return nil
		}
	}
	defer rows.Close()
	var choices []*choice
	for rows.Next() {
		ch := &choice{}
		err := rows.Scan(&ch.ChoiceID, &ch.PollID, &ch.Choice)
		if err != nil {
			log.Fatal(err)
		}
		choices = append(choices, ch)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(`CHOICES ==>`, choices)
	return choices
}

func dbGetPolls() []*poll {
	db, _ := connectDB()
	rows, err := db.Query("SELECT * FROM voting.polls")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return nil
		}
	}
	defer rows.Close()
	var polls []*poll
	for rows.Next() {
		p := &poll{}
		err := rows.Scan(&p.PollID, &p.Question)
		if err != nil {
			log.Fatal(err)
		}
		polls = append(polls, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(`POLLS ==>`, polls)
	return polls
}
