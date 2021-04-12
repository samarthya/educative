package utils

import (
	"errors"
	"log"
	"sync"
)

// URLDB url datastore
type URLDB struct {
	sync.Mutex
	m map[string]string
}

// Interface exposes the methods for the db
type Interface interface {
	// Add the key and value
	Add(string, string) bool
	// Get the value from the map
	Get(string) (string, error)
	//Len returns the length of the store
	Len() int
}

//Len returns length of the map
func (u *URLDB) Len() int {
	return len(u.m)
}

// NewDB Creates a new data base
func NewDB() *URLDB {
	return &URLDB{m: make(map[string]string, 1)}
}

// Add the url in the DB
func (u *URLDB) Add(k string, v string) bool {
	u.Lock()
	defer u.Unlock()
	if n, found := u.m[k]; found {
		log.Println("value will be overwritten", n)
	}
	log.Println("inserting value", v)
	u.m[k] = v
	return true
}

// Get the url in the DB
func (u *URLDB) Get(k string) (string, error) {
	u.Lock()
	defer u.Unlock()

	if n, found := u.m[k]; found {
		log.Println("value will be overwritten", n)
		return n, nil
	}

	log.Println("key not found", k)

	return "", errors.New("key missing")
}
