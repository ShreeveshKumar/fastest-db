package main
import(
	"fmt"
	"sync"
	"time"
)

type Point struct{
  latitude float64
  longitude float64
}

type CachedPoint struct{
	Point
	updatedAT time.Time
	dirty bool
	expiresIN time.Time
}

type Store struct{
	mu sync.RWMutex
	data map[string]CachedPoint
	wal WAL 
	flusher Flusher
	ttl time.Duration
 }






 





func addPoint(key string,point Point)(bool){
	s.mu.Lock() // basically system put in lock so no otheer can perform same operation 
	defer s.mu.unlock()

	
    


   if key {
	error.New("key already exits")
   }



}

func viewPoint()(){
	

}

func updatePoint()(){
	

}

func deletePoint()(){
	

}