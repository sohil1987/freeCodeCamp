//https://rosettacode.org/wiki/SQL-based_authentication#Go

package nightlife

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Comment
)

func doLoginOrCreate(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	defer db.Close()
	r.ParseForm()
	user := r.Form["user"][0]
	pwd := r.Form["pass"][0]
	err = createUser(db, user, pwd)
	if err != nil { // user already exists
		fmt.Println(err)
		err = authenticateUser(db, user, pwd)
		if err != nil { // invalid password
			fmt.Println(err)
			wrongLogin(w, r)
			return
		}
	}
	setSessionCookie(w, r, user)
	search := dbGetLastSearch(user)
	http.Redirect(w, r, baseURL+"nightlife/logged?user="+user+"&search="+search, 301)
}

func createUser(db *sql.DB, user, pwd string) error {
	salt := make([]byte, 16)
	rand.Reader.Read(salt)
	_, err := db.Exec("INSERT INTO nightlife.users (Username, PassSalt, PassMd5)     values (?, ?, ?)", user, salt, saltHash(salt, pwd))
	if err != nil {
		return fmt.Errorf("User %s already exits", user)
	}
	return nil
}

func authenticateUser(db *sql.DB, user, pwd string) error {
	var salt, hash []byte
	row := db.QueryRow("SELECT PassSalt, PassMd5 from nightlife.users WHERE Username=?", user)
	defer db.Close()
	if err := row.Scan(&salt, &hash); err != nil {
		return fmt.Errorf("User %s unknown", user)
	}
	if !bytes.Equal(saltHash(salt, pwd), hash) {
		return fmt.Errorf("User %s invalid password", user)
	}
	return nil
}

func saltHash(salt []byte, pwd string) []byte {
	h := md5.New()
	h.Write(salt)
	h.Write([]byte(pwd))
	return h.Sum(nil)
}

func dbGetAllUsers(db *sql.DB) {
	rows, _ := db.Query("SELECT Username, PassSalt, PassMd5 FROM nightlife.users")
	defer rows.Close()
	var user string
	var salt, hash []byte
	for rows.Next() {
		rows.Scan(&user, &salt, &hash)
		fmt.Printf("%s %x %x\n", user, salt, hash)
	}
	// clear table to run program again
	//db.Exec("TRUNCATE table users")
}
