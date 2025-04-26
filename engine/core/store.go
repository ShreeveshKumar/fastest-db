package main

import (
	"errors"
	"sync"
	"time"
)

type Point struct {
	latitude  float64
	longitude float64
}

type CachedPoint struct {
	Point
	updatedAT time.Time
	dirty     bool
	expiresIN time.Time
}

type Flusher struct {
	// Define the fields for Flusher
}

type Store struct {
	mu      sync.RWMutex
	data    map[string]CachedPoint
	wal     WAL
	flusher Flusher
	ttl     time.Duration
}

var s = Store{
	data: make(map[string]CachedPoint),
	ttl:  5 * time.Minute, // Example TTL value
}

func addPoint(key string, point Point) (bool, error) {
	s.mu.Lock() // basically system put in lock so no other can perform same operation
	defer s.mu.Unlock()

	if _, exists := s.data[key]; exists {
		return false, errors.New("key already exists")
	}

	s.data[key] = CachedPoint{
		Point:     point,
		updatedAT: time.Now(),
		dirty:     false,
		expiresIN: time.Now().Add(s.ttl),
	}

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
