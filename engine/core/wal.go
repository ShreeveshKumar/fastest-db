


package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const filepath = "spatial-db/engine/log/wal.log"

type WAL struct {
	mu   sync.Mutex
	file *os.File
}

type WALRecord struct {
	OpType    string    `json:"op_type"`
	Key       string    `json:"key"`
	Value     Point     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}



type Op struct {
	OpType string
	Value  Point
}

func (w *WAL) Append(opType string, key string, value Point) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	record := WALRecord{
		OpType:    opType,
		Key:       key,
		Value:     value,
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return fmt.Errorf("failed to write to WAL file: %w", err)
	}

	return nil
}

func (w *WAL) Recover() ([]Op, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open WAL file: %w", err)
	}
	defer file.Close()

	var ops []Op
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)

		if len(parts) != 2 {
			continue
		}

		opType := parts[0]
		jsonPoints := parts[1]

		var point Point
		err := json.Unmarshal([]byte(jsonPoints), &point)
		if err != nil {
			return nil, fmt.Errorf("error parsing JSON: %w", err)
		}

		ops = append(ops, Op{
			OpType: opType,
			Value:  point,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return ops, nil
}

func main() {
	// Example usage
	wal := &WAL{}
	err := wal.Append("insert", "key1", Point{latitude: 1, longitude: 2})
	if err != nil {
		fmt.Println("Error appending to WAL:", err)
	}

	ops, err := wal.Recover()
	if err != nil {
		fmt.Println("Error recovering from WAL:", err)
	} else {
		fmt.Println("Recovered operations:", ops)
	}
}


