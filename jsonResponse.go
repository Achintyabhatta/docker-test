package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// prepares a Json response for the user
func getDataJsonEncoder(w http.ResponseWriter, ipValue string, status bool) {
	w.Header().Set("Content-Type", "application/json")
	responseValue := make(map[string]string)
	responseValue["ipAddr"] = ipValue
	errorMessage := make(map[string]string)
	errorMessage["message"] = ""
	if status == false {
		errorMessage["message"] = "Either domain name is Invalid or IP not registered!"
	}
	resp := map[string]interface{}{"success": status, "data": responseValue, "error": errorMessage}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

// prepares a JsonResponse for POST method used by the user
func postDataJsonEncoder(w http.ResponseWriter, messageType string, status bool) {
	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]string)
	if messageType == "wrongDomain" {
		message["message"] = "Given Domain Name not valid!"
	} else if messageType == "wrongIP" {
		message["message"] = "Given IPv4 Value not Valid!"
	} else if messageType == "dataSaved" {
		message["message"] = "IP Value saved Successfully!"
	}
	resp := map[string]interface{}{"success": status, "message": message}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
