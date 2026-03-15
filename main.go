package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if _, err := os.Stat("downloads.html"); os.IsNotExist(err) {
		log.Fatal("ERROR: downloads.html was not found in this folder!")
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received from %s for %s", r.RemoteAddr, r.URL.Path)
		
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "downloads.html")
	})

	fmt.Println("-------------------------------------------")
	fmt.Println("SERVER STARTED!")
	fmt.Println("Open this link: http://YOUR IP:9999")
	fmt.Println("-------------------------------------------")

	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}
