package main

import (
	"net/http"
)

// handler function to do IpLookup and return response to the user
func getIPValue(db database, w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")

	// Validates domainName
	if !validateDomainName(domain) {
		getDataJsonEncoder(w, "", false)
		return
	}
	// Attemp to retrive the IP from the database
	savedIP, ok := db[domain]

	// If IP not found in the database
	if !ok {
		ipValue := searchIP(domain) // perform a IP Lookup and save it in the database
		if ipValue == "" {
			getDataJsonEncoder(w, "", false)
			//w.Write([]byte("No IP found for this Domain"))
			return
		}
		db[domain] = ipValue
		getDataJsonEncoder(w, ipValue, true) // response written to the user
		return
	}
	// if IP is found in the database
	db[domain] = savedIP
	getDataJsonEncoder(w, savedIP, true) // Response written to the user
}
