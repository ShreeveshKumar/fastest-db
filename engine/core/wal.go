package main
import (
	"fmt"
	"sync"
	"os"
	"time"
	"encoding/json"
	"bufio"
)

const filepath = "spatial-db/engine/log/wal.log"

type WAL struct{
	mu sync.Mutex
	file *os.File
}

type WALRecord struct{
	OpType    string    `json:"op_type"`
    Key       string    `json:"key"`
	Value     Point     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}


func (w *WAL)Append(opType string, key string, value Point)(error){
	w.m.lock(); 
	defer w.m.unlock(); 

	record := WALRecord{
	 OpType:op_type,
	 Key:key,
	 Value:value
	 Timestamp:time.Now()
	}

	data,err = json.Marshal(record)

	if err!=nil {
		fmt.Errorf("failed to create Marhsal from record %w",err)
	}

	file,err = os.OpenFile(filepath)

	file,err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); 
	if err!= nil{
		fmt.Errorf("Error writing to file");
	}
	
	defer file.Close(); 

	file,err = file.Write(append(data,'\n')); 

	if err!=nil {
		return fmt.Errorf("failed to write to WAL file: %w", err)
	}

	return nil; 
}


func (w *WAL)Recover()([]Op, error){
	w.mu.lock(); 
	defer w.mu.unlock(); 

	file, error = os.Open(w.filepath); 

	if err != nil {
        return nil, fmt.Errorf("failed to open WAL file: %w", err)
    }

    defer file.Close()
	var ops []Op; 

    scanner := bufio.NewScanner(file); 
	for scanner.Scan() {
        line,err := scanner.Text()
        parts = strings.split(line,'|',2); 

		if len(parts) != 2 {
            continue 
        }

		opType = parts[0]; 
		jsonpoints = parts[1]; 

		var point Point; 
		err := json.UnMarshal([]byte(jsonpoints,&point)); 
		if err != nil {
            return nil, fmt.Errorf("error parsing JSON: %w", err)
        }
      

		ops = append(ops, Op{
			OpType:opType
			Value:point
		}); 

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("scanner error: %w", err)
		}

		return ops, nil
    }

}

