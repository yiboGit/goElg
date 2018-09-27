package main 

import "eglass.com/sta"
import (
	"net/http"
	"log"
	"encoding/json"
)


func main() {
	sta.Scan("/Users/amyli/go/src/eglass.com")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// recover()
		}()
		stats, error := sta.Scan(r.URL.Path)
		if (error != nil) {
			http.Error(w, "", 400)
			log.Printf("%v", error)
			return
		}
		log.Println("files: ", stats)
		json.NewEncoder(w).Encode(stats)
	})
	log.Printf("listen on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}