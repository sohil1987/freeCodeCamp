package nightlife

import (
	"database/sql"
	"fmt"
	"log"
)

func dbGetLastSearch(user string) string {
	var search string
	row := db.QueryRow("SELECT LastSearch FROM nightlife.users WHERE Username=?", user)
	err := row.Scan(&search)
	if err != nil {
		log.Fatalln(err)
	}
	return search
}

func dbSaveLastSearch(user, search string) {
	_, err := db.Exec("UPDATE nightlife.users SET LastSearch=? WHERE Username=?", search, user)
	if err != nil {
		log.Fatalln(err)
	}
}

func dbUserAlreadyVotedBar(user, id string) bool {
	var username string
	rows, err := db.Query("SELECT Username FROM nightlife.votes WHERE BarId=?", id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&username)
		if user == username {
			return true
		}
	}
	return false
}

func dbAddUserVoteBar(user, id string) {
	_, err := db.Exec("INSERT INTO nightlife.votes (Username, barID) values (?, ?)", user, id)
	if err != nil {
		log.Fatalln(err)
	}
}

func dbRemoveUserVoteBar(user, id string) {
	rows, err := db.Exec("DELETE FROM nightlife.votes WHERE Username = ? AND BarId = ?", user, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
		} else {
			log.Fatal(err)
		}
	}
	res, _ := rows.RowsAffected()
	fmt.Sprintln("DELETED", res)
}
