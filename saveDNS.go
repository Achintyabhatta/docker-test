package main

import (
	"net/http"
)

// handler function that gets DomainName and IpAddress and saves it to the database
func setIPValue(db database, w http.ResponseWriter, r *http.Request) {

	domainName := r.FormValue("domainName")
	ipAddr := r.FormValue("ipAddr")

	// Validates domainName
	if !validateDomainName(domainName) {
		postDataJsonEncoder(w, "wrongDomain", false)
		return
	}
	if !validIP4(ipAddr) {
		postDataJsonEncoder(w, "wrongIP", false)
		return
	}
	db[domainName] = ipAddr
	postDataJsonEncoder(w, "dataSaved", true)
	return
}
