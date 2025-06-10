package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const file_path = "../log/wal.log"; 
func addPoint(key string, point Point) (bool, error) {
	wal := &WAL{}

	file_exists = os.OpenFile(file_path,)

	// func (w *WAL) Append(opType string, key string, value Point) error {

	err := wal.Append("ADD", key, point);

	if err != nil {
		fmt.Printf("Error is caused by", err)
	}

	// till here we have appended in wal
    

	return true, nil
}

func viewPoint(key string) (Point, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if cachedPoint, exists := s.data[key]; exists {
		return cachedPoint.Point, nil
	}

	return Point{}, errors.New("point not found")
}

func updatePoint(key string, point Point) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; !exists {
		return false, errors.New("key does not exist")
	}

	s.data[key] = CachedPoint{
		Point:     point,
		updatedAT: time.Now(),
		dirty:     true,
		expiresIN: time.Now().Add(s.ttl),
	}

	return true, nil
}

func deletePoint(key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; !exists {
		return false, errors.New("key does not exist")
	}

	delete(s.data, key)
	return true, nil
}
