package main

import (
	"log"
	"net/http"
)

const (
	ENDTIME = "May 16, 2018 17:30:00"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("[STARTING] On PORT 3000")
	http.ListenAndServe(":3000", nil)
}
