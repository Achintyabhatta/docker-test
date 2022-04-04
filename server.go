package main

import (
	"log"
	"net/http"
)

type database map[string]string // A map that will mimic a database while server is running

// handler function that gets DomainName and IpAddress and saves it to the database
func (db database) saveDNS(w http.ResponseWriter, r *http.Request) {
	setIPValue(db, w, r)
}

// handler function to do IpLookup and return response to the user
func (db database) ipLookup(w http.ResponseWriter, r *http.Request) {
	getIPValue(db, w, r)
}

func main() {
	if err := http.ListenAndServe(":8084", handler()); err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	db := database{}
	mux := http.NewServeMux()
	mux.HandleFunc("/DNS/save", db.saveDNS)
	mux.HandleFunc("/DNS/iplookup", db.ipLookup)
	return mux
}
