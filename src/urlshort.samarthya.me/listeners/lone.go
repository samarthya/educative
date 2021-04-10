package listners

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"urlshort.samarthya.me/utils"
)

// DefaultHandler for the server method
var DefaultHandler MyHandler = MyHandler{}

// MyHandler is a empty struct to implement the handler interface
type MyHandler struct {
	DB *utils.URLDB
}

//DumpHeaders is to dump request headers
func DumpHeaders(rq *http.Request) {
	for k, v := range rq.Header {
		log.Println(" Key ", k, " = ", v)
	}
}

// Add the url to the store
func Add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("return the form to process")
	case http.MethodPost:
		log.Println("process the url to be added")
	}

	log.Println("/add requested")
	url := r.FormValue("url")
	fmt.Println("URL: ", url)
	DumpHeaders(r)
	io.WriteString(w, "/add requested")
}

// ServeHTTP Serves the HTTP Handler
func (MyHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	log.Println(" ServeHTTP invoked")
	DumpHeaders(rq)
	switch rq.RequestURI {
	case "/":
		io.WriteString(rw, "/ was requested")
		return
	case "/add":
		Add(rw, rq)
		return
	}
	rw.Write([]byte("Listening is on...."))
}

// Listen is a Dummy listener
func Listen() {
	log.Println(" Listen invoked ")
}
