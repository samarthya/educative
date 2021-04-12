package listeners

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"strings"

	"urlshort.samarthya.me/utils"
)

const (
	//LOCALHOST identifies the localhost string
	LOCALHOST = "http://localhost:8181/get/"
	// ADD uri
	ADD = "/add"
	// STATIC uri
	STATIC = "/static"
	// GET uri
	GET = "/get"
	// ContentType type of content
	ContentType = "Content-Type"
)

// MyHandler is a empty struct to implement the handler interface
type MyHandler struct {
	DB  *utils.URLDB
	Tmp *template.Template
}

//DumpHeaders is to dump request headers
func DumpHeaders(rq *http.Request) {
	for k, v := range rq.Header {
		log.Println(" Key ", k, " = ", v)
	}
}

// Redirect to the request coming in
func (m *MyHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	key := LOCALHOST
	key += r.URL.Path[5:]
	url, err := m.DB.Get(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// FileServerHandler the file server request
func (m *MyHandler) FileServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("read the contents")
	switch {
	case strings.HasSuffix(r.RequestURI, "index.html"):
		if m.Tmp != nil {
			m.Tmp.ExecuteTemplate(w, "index", nil)
		}
	default:
		log.Println("handle the content...")
		extension := r.RequestURI[strings.LastIndex(r.RequestURI, "."):]
		if b, err := ioutil.ReadFile(r.RequestURI); err == nil {
			w.Header().Set(ContentType, mime.TypeByExtension(extension))
			if strings.Contains(mime.TypeByExtension(extension), "image") {
				w.Write(b)
			} else {
				io.WriteString(w, string(b))
			}
		}
	}
}

// Add the url to the store
func (m *MyHandler) Add(w http.ResponseWriter, r *http.Request) {
	log.Println("/add requested", r.Method)
	// Only if you want to see the headers
	// DumpHeaders(r)

	switch r.Method {
	case http.MethodGet:
		log.Println("return the form to process")
		if m.Tmp != nil {
			m.Tmp.ExecuteTemplate(w, "input-form", nil)
		}
	case http.MethodPost:
		log.Println("process the url to be added")
		url := r.FormValue("url")
		fmt.Println("URL: ", url)
		key := fmt.Sprintf("%s%d", LOCALHOST, m.DB.Len())
		m.DB.Add(key, url)
		io.WriteString(w, key)
	}
}

// ServeHTTP Serves the HTTP Handler
func (m *MyHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	log.Println("ServeHTTP...")
	// Enable for header dump
	// DumpHeaders(rq)
	log.Println("URI > ", rq.RequestURI)

	switch rq.Method {
	case http.MethodGet:
		// If it is Add or get the data
		if strings.Index(rq.RequestURI, ADD) == 0 {
			m.Add(rw, rq)
		} else if strings.Index(rq.RequestURI, STATIC) == 0 {
			m.FileServerHandler(rw, rq)
		} else if strings.Index(rq.RequestURI, GET) == 0 {
			m.Redirect(rw, rq)
		} else {

		}
	case http.MethodPost:
		m.Add(rw, rq)
	}

}

// Listen is a Dummy listener
func Listen() {
	log.Println(" Listen invoked ")
}
