package main

import (
    "database/sql"
    "fmt"
    "log"
	"strconv"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    id       int
    name     string
}

type Seat struct {
    id       int
    name     string
}

// Function to initialize the database connection
func initDB() *sql.DB {
    // Connect to the database using the username, password, and database name (replace "username:password@/dbname")
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/user")
    if err != nil {
        log.Fatal(err)
    }

    // Check the database connection
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    return db
}

// Function to insert a new user into the database
func insertUser(db *sql.DB, user User) error {
    query := "INSERT INTO user(id, name) VALUES(?, ?)"
    _, err := db.Exec(query, user.id, user.name)
    if err != nil {
        return err
    }
    fmt.Println("User inserted successfully:", user.name)
    return nil
}

// Function to insert a new seat into the database
func insertSeat(db *sql.DB, seat Seat) error {
    query := "INSERT INTO seat(id, name) VALUES(?, ?)"
    _, err := db.Exec(query, seat.id, seat.name)
    if err != nil {
        return err
    }
    fmt.Println("Seat inserted successfully:", seat.name)
    return nil
}

func insertDummyUsers() {
	db := initDB()
    defer db.Close()

	for i:= 0; i<100; i++ {
		newUser := User{
			id: i,
			name: "user-" + strconv.Itoa(i),
		}
		if err := insertUser(db, newUser); err != nil {
			log.Fatal(err)
		}
	}
}

func insertDummySeats() {
	db := initDB()
    defer db.Close()

	for i:= 0; i<100; i++ {
		newSeat := Seat{
			id: i,
			name: "seat-" + strconv.Itoa(i),
		}
		if err := insertSeat(db, newSeat); err != nil {
			log.Fatal(err)
		}
	}
}





func main() {
	insertDummyUsers()
	insertDummySeats()	
}


