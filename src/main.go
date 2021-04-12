package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/magiconair/properties"
	"urlshort.samarthya.me/listeners"

	db "urlshort.samarthya.me/utils"
)

const (
	//TEMPLATES folder that contains the templates
	TEMPLATES = "templates.dir"
)

var s *http.Server

var myHandler = &listeners.MyHandler{DB: db.NewDB()}
var urlDB *db.URLDB

var props *properties.Properties

func init() {
	log.Println(" Initialized ")
	s = &http.Server{
		Addr:           ":8181",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	props = properties.MustLoadFile("./config.properties", properties.UTF8)

}

func getTemplateRoot() string {
	if v, found := props.Get(TEMPLATES); found {
		return v
	}
	return "template/"
}

func pwd() {
	if dl, err := os.ReadDir("."); err == nil {
		for _, v := range dl {
			if v.IsDir() {
				log.Println("inside", v.Name())
			}
		}
	}
}

var webTemplates listeners.FilesArray
var htmlTemplates *template.Template

// main method
func main() {
	log.Println("Starting server....", s)
	fmt.Println("--- URL Shortner ---")

	pwd()

	if s := getTemplateRoot(); s != "" {
		webTemplates = listeners.FileList(s)
		htmlTemplates, err := listeners.LoadTemplates(webTemplates)

		// load the templates
		if err == nil {
			myHandler.Tmp = htmlTemplates
		}
	}

	// listeners.Listen()
	log.Fatal(s.ListenAndServe())
}
