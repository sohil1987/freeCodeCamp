package pintelest

import (
	"database/sql"
	"fmt"
	"freeCodeCamp/7-bonus2/_help"
	"log"

	_ "github.com/go-sql-driver/mysql" // justify
)

type myDB struct{}

var db *sql.DB

var mydb myDB

func (mydb *myDB) initDB() {
	connPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Db)
	//fmt.Println(connPath)
	var err error
	db, err = sql.Open("mysql", connPath)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("pintelest sql.Open() OK")
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("pintelest db.Ping() OK")
	}
}

func (mydb *myDB) dbGetAllPics() {
	sql := "SELECT idPic, pics.idTwitter, users.nameTwitter AS username, users.logoTwitter AS logoAuthor, text, link, likes FROM pintelest.pics JOIN users ON pics.idTwitter = users.idTwitter"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	imgs = Images{}
	var i Image
	for rows.Next() {
		err := rows.Scan(&i.IDPic, &i.IDAuthor, &i.NameAuthor, &i.LogoAuthor, &i.Text, &i.Link, &i.Likes)
		if err != nil {
			log.Fatal(err)
		}
		imgs = append(imgs, i)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (mydb *myDB) dbGetUserPics(userIDTwitter string) {
	sql := "SELECT idPic, pics.idTwitter, users.nameTwitter AS username, users.logoTwitter AS logoAuthor, text, link, likes FROM pintelest.pics JOIN users ON pics.idTwitter = users.idTwitter WHERE pics.idTwitter = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userIDTwitter)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	imgs = Images{}
	var i Image
	for rows.Next() {
		err := rows.Scan(&i.IDPic, &i.IDAuthor, &i.NameAuthor, &i.LogoAuthor, &i.Text, &i.Link, &i.Likes)
		if err != nil {
			log.Fatal(err)
		}
		imgs = append(imgs, i)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (mydb *myDB) insertUser(text, link string) {
	sql := "INSERT INTO users(idTwitter, nameTwitter,logoTwitter) SELECT ?, ?, ? FROM dual WHERE NOT EXISTS (SELECT idTwitter FROM users WHERE idTwitter = ?)  LIMIT 1"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.User.UserID, c.User.Name, c.User.AvatarURL, c.User.UserID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("New User saved ......")
		if !help.IsValidURL(link) {
			link = "https://brusbilis.com/freecodecamp/assets/images/photoNot.png"
		}
		mydb.insertPic(c.User.UserID, text, link)
	}

}

func (mydb *myDB) insertPic(id, text, link string) {
	sql := "INSERT INTO pics (idTwitter, text, link) VALUES(?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, text, link)
	if err != nil {
		log.Fatal(err)
	}
}

func (mydb *myDB) deletePic(idpic string) {
	var sql = "DELETE FROM pics WHERE idPic = ? "
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(idpic)
	if err != nil {
		log.Fatal(err)
	}
}

func (mydb *myDB) vote(userIDTwitter, idPic string) {
	fmt.Println(userIDTwitter, " vota a la pic ", idPic)
	sql := "SELECT COUNT(*) as times FROM likes WHERE idPic = ? AND idTwitter = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var option int
	err = stmt.QueryRow(idPic, userIDTwitter).Scan(&option)
	if err != nil {
		log.Fatalln(err)
	}
	mydb.insertVote(userIDTwitter, idPic, option)
}

func (mydb *myDB) insertVote(userIDTwitter, idPic string, option int) {
	var sql string
	if option == 0 {
		sql = "INSERT INTO likes VALUES (?, ?)"
	}
	if option == 1 {
		sql = "DELETE FROM likes WHERE idPic = ? AND idTwitter = ?"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(idPic, userIDTwitter)
	if err != nil {
		log.Fatal(err)
	}
	mydb.updateLikesToPics(idPic, option)
}

func (mydb *myDB) updateLikesToPics(idPic string, option int) {
	var sql string
	if option == 0 {
		sql = "UPDATE pics SET likes = likes +1 WHERE idPic = ?"
	}
	if option == 1 {
		sql = "UPDATE pics SET likes = likes -1 WHERE idPic = ?"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(idPic)
	if err != nil {
		log.Fatal(err)
	}
}
