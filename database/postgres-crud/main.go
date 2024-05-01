// basic implementation of crud using golang standard library
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Person struct {
	Username string
	Email    string
}

func main() {
	// put this on env for production app
	connStr := "user=postgres dbname=postgres_test password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	defer func() {
		db.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createPersonTable(db)
	insertPerson(db, Person{Username: "frastio10", Email: "hi@frast.dev"})
	insertPerson(db, Person{Username: "jamal246", Email: "hi@jamal.tata"})

	deletePerson(db, 8)

	updatedId := updatePerson(db, 1, &Person{Username: "ota", Email: "otata@gmail.com"})
	fmt.Println("Person updated: ", updatedId)

	people := getAllPerson(db)
	fmt.Println("People: ", people)

	person := getPersonById(db, 1)
	fmt.Println("Person: ", person)
}

func deletePerson(db *sql.DB, id int) {
	q := `
    DELETE FROM person WHERE id = $1
  `

	_, err := db.Exec(q, id)
	if err != nil {
		log.Println("Failed to delete person: ", err)
	}
}

func updatePerson(db *sql.DB, id int, person *Person) int {

	q := `
        UPDATE person 
        SET username = $1, email = $2 
        WHERE id = $3
  `

	_, err := db.Exec(q, person.Username, person.Email, id)
	if err != nil {
		log.Fatal("Failed to update person: ", err)
	}

	return id
}

func getPersonById(db *sql.DB, id int) Person {
	q := `SELECT username, email FROM person WHERE id = $1`
	row := db.QueryRow(q, id)

	var username string
	var email string
	err := row.Scan(&username, &email)
	if err != nil {
		log.Fatal("Failed to find person: ", err)
	}

	return Person{Username: username, Email: email}
}

// sounds weird but ok
func insertPerson(db *sql.DB, p Person) int {
	q := `INSERT INTO person (username, email) VALUES ($1, $2) RETURNING id`

	var pk int
	err := db.QueryRow(q, p.Username, p.Email).Scan(&pk)
	if err != nil {
		log.Fatal("Person Insert Error: ", err)
	}

	return pk
}

func getAllPerson(db *sql.DB) []Person {
	rows, err := db.Query("SELECT username, email FROM person")
	if err != nil {
		log.Fatal("Cannot get person list: ", err)
	}

	defer rows.Close()

	people := []Person{}

	for rows.Next() {
		var username string
		var email string

		err = rows.Scan(&username, &email)
		if err != nil {
			log.Fatal("Failed to scan person: ", err)
		}

		people = append(people, Person{Username: username, Email: email})

	}

	return people

}

func createPersonTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS person (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created timestamp DEFAULT NOW()
  )`

	_, err := db.Exec(q)
	if err != nil {
		log.Fatal("Failed to create person table: ", err)
	}

}
