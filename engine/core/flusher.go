// flusher.go – 🕒 Batch Flush System
// Purpose: Periodically flush in-memory data to DB or disk.

// Core logic:

// Background goroutine with ticker.

// Picks up recently changed data from memory or WAL.

// Writes to cold storage (disk, database, file).

// Optional: removes from memory after flush if overflowed.

// ✨ Optional: Make flush frequency configurable.

package core

import (
	"fmt"
)


func sayHello(){
	fmt.Println("Hello this is good")
}
