package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	listners "urlshort.samarthya.me/listeners"
	db "urlshort.samarthya.me/utils"
)

var s *http.Server

var myHandler = listners.MyHandler{DB: &db.URLDB{}}
var urlDB *db.URLDB

func init() {
	log.Println(" Initialized ")
	s = &http.Server{
		Addr:           ":8181",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

// main method
func main() {
	fmt.Println(" URL Shortner ")
	listners.Listen()
	log.Fatal(s.ListenAndServe())
}
