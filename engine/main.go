package main; 
import (
	"fmt"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to Home page")
}



func main(){
   http.HandleFunc("/",homePage); 

   const port= ":8000";
   fmt.Printf("🚀 Server is running at http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil);
	
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}	
}
