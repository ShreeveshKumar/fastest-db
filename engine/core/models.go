package main; 
import(
	"time"
	"sync"
	"os"
)

type GeoFeature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	Time Timestamp `json:"timestamp"`
	GeoHash string `json:"geohash`
}

type Properties struct {
	Name         string  `json:"name"`
	Crop         string  `json:"crop"`
	AreaHectares float64 `json:"area_hectares"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"` 
}


type Timestamp struct{
	Time time.Time `json:"time"`
}


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
