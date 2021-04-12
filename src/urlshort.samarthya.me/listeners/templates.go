package listeners

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// FilesArray is an array of string for File
type FilesArray []string

//FileList list of files
func FileList(td string) (m FilesArray) {
	log.Println("reading ", td)
	m = make(FilesArray, 0)

	if d, err := os.ReadDir(td); err == nil {
		for _, v := range d {
			if !v.IsDir() {
				log.Println("found", v.Name())
				m = append(m, fmt.Sprintf("%s%s", td, v.Name()))
			}
		}
	}
	return
}

// LoadTemplates Loads templates
func LoadTemplates(p FilesArray) (*template.Template, error) {
	return template.ParseFiles(p...)
}
