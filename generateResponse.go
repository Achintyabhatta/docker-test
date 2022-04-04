package main

func getDummyGetResponse(domain string) string {
	ipVal := searchIP(domain)
	if !validateDomainName(domain) || ipVal == "" {
		response := "{\"data\":{\"ipAddr\":\"\"},\"error\":{\"message\":\"Either domain name is Invalid or IP not registered!\"},\"success\":false}"
		return response
	}
	response := "{\"data\":{\"ipAddr\":" + "\"" + string(ipVal) + "\"" + "},\"error\":{\"message\":\"\"},\"success\":true}"
	return response
}

func getDummySetResponse(domain string, ipaddr string) string {
	if !validateDomainName(domain) {
		response := "{\"message\":{\"message\":\"Given Domain Name not valid!\"},\"success\":false}"
		return response
	}
	if !validIP4(ipaddr) {
		response := "{\"message\":{\"message\":\"Given IPv4 Value not Valid!\"},\"success\":false}"
		return response
	}
	response := "{\"message\":{\"message\":\"IP Value saved Successfully!\"},\"success\":true}"
	return response
}
