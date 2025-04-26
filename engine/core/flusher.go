// flusher.go – 🕒 Batch Flush System
// Purpose: Periodically flush in-memory data to DB or disk.

// Core logic:

// Background goroutine with ticker.

// Picks up recently changed data from memory or WAL.

// Writes to cold storage (disk, database, file).

// Optional: removes from memory after flush if overflowed.

// ✨ Optional: Make flush frequency configurable.

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	userisname = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
)

const connectionStr = "user=wait password=wait dbname=postgres host=localhost port=5432 sslmode=disable"

func writeToDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	fmt.Printf(dbUser, dbPass)

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("err connectinhg to DB")
	}

    // query = 



}
