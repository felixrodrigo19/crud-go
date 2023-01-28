package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Getenv("CGO_ENABLED")

	driverName := "sqlite3"
	dbSource := "./crud-local.db"
	db, err := sql.Open(driverName, dbSource)
	checkErr(err)

	// insert on table
	insertData(db, "dog")

	getData(db, "dog")

	deleteData(db, 2)
	db.Close()
}

func deleteData(db *sql.DB, row int) {
	// delete
	stmt, err := db.Prepare("delete from animals where id=?")
	checkErr(err)

	res, err := stmt.Exec(row)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func getData(db *sql.DB, description string) string {
	if description == "" {
		return description
	}
	stmt, err := db.Prepare("SELECT * FROM animals WHERE description=?")

	var id string
	var desc string
	var created string
	err = stmt.QueryRow(description).Scan(&id, &desc, &created)
	if err != nil {
		return ""
	}

	fmt.Println(id, desc, created)
	return desc
}

func insertData(db *sql.DB, animal string) {
	row := getData(db, animal)
	if row != "" {
		fmt.Printf("Animal %s already created on table\n", animal)
		return
	}
	if animal == "" {
		return
	}

	stmt, err := db.Prepare("INSERT INTO animals(description, created) values(?,?)")

	currentDate := time.Now()

	res, err := stmt.Exec(animal, currentDate)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
