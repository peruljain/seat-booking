package main

import (
    "database/sql"
    "fmt"
    "log"
	"sync"
	"time"

    _ "github.com/go-sql-driver/mysql"
)

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

func bookSeat(db *sql.DB, userId int) error {
	// Start a transaction
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    // Defer a rollback in case anything fails. The rollback will be ignored if the transaction is successfully committed later.
    defer tx.Rollback()

    // 1. Find an available seat where user_id is NULL
    var seatID int
    query := `SELECT id FROM seat WHERE user_id IS NULL ORDER BY id LIMIT 1 FOR SHARE`
    err = tx.QueryRow(query).Scan(&seatID)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No available seats.")
            return nil  // No available seat, not necessarily an error
        }
        return err
    }

    // 2. Update the found seat with the given user_id
    updateQuery := `UPDATE seat SET user_id = ? WHERE id = ?`
    _, err = tx.Exec(updateQuery, userId, seatID)
    if err != nil {
        return err
    }

    // Commit the transaction
    if err := tx.Commit(); err != nil {
        return err
    }

    fmt.Printf("Seat %d has been assigned to user %d\n", seatID, userId)
    return nil
}

func main() {
	db := initDB()
	defer db.Close()

	var wg sync.WaitGroup

	start := time.Now() 

	for i:=0;  i<100; i++ {
		wg.Add(1)
		// time.Sleep(1 * time.Millisecond)
		go func(userId int) {
            defer wg.Done()
            err := bookSeat(db, userId)
            if err != nil {
                log.Printf("Failed to book seat for user %d: %v", userId, err)
            }
        }(i)

	}

	wg.Wait()
	elapsed := time.Since(start) 
	fmt.Printf("All seats booked in %s\n", elapsed)

}