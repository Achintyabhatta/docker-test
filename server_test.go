package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type dataValue struct {
	Ipaddr string
}

type errorValue struct {
	Message string
}

type getResponse struct {
	Success string
	Error   errorValue
	Data    dataValue
}

type getTestData struct {
	Url           string
	Responsevalue string
}

type getSetData struct {
	Url           string
	Ipaddr        string
	Responsevalue string
}

func TestGetServerHandler(t *testing.T) {
	rootURL := "localhost:8081/DNS/iplookup?domain="
	testData := []getTestData{
		getTestData{
			Url:           "google.com",
			Responsevalue: getDummyGetResponse("google.com"),
		},
		getTestData{
			Url:           "google.com",
			Responsevalue: getDummyGetResponse("google.com"),
		},
		getTestData{
			Url:           "google",
			Responsevalue: getDummyGetResponse("google"),
		},
		getTestData{
			Url:           "dasdbasbdasdas.com",
			Responsevalue: getDummyGetResponse("dasdbasbdasdas.com"),
		},
	}
	db := database{}

	for _, value := range testData {
		req, err := http.NewRequest("GET", rootURL+value.Url, nil)
		if err != nil {
			t.Fatalf("Could not create request %v:", err)
		}
		rec := httptest.NewRecorder()
		getIPValue(db, rec, req)
		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK; got %v\n", res.StatusCode)
		}
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read response: %v \n", err)
		}
		bodyString := string(bodyBytes)
		if value.Responsevalue != bodyString {
			t.Fatalf("Returned GET Response not same as expected response\n")
		}
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()
	res, err := http.Get(fmt.Sprintf("%s/DNS/iplookup?domain=google.com", srv.URL))
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected Status Ok, Got : %v", res.StatusCode)
	}
}

func TestSetServerHandler(t *testing.T) {

	testData := []getSetData{
		getSetData{
			Url:           "google.com",
			Ipaddr:        "10.10.10.10",
			Responsevalue: getDummySetResponse("google.com", "10.10.10.10"),
		},
		getSetData{
			Url:           "google",
			Ipaddr:        "10.10.10.10",
			Responsevalue: getDummySetResponse("google", "10.10.10.10"),
		},
		getSetData{
			Url:           "google.com",
			Ipaddr:        "10.10.1",
			Responsevalue: getDummySetResponse("google.com", "10.10.1"),
		},
	}
	db := database{}
	for _, value := range testData {
		data := url.Values{}
		data.Set("domainName", value.Url)
		data.Set("ipAddr", value.Ipaddr)
		req, err := http.NewRequest("POST", "/DNS/save", strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			t.Fatalf("Could not create POST request %v:", err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(db.saveDNS)
		handler.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK; got %v\n", res.StatusCode)
		}
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read response: %v \n", err)
		}
		bodyString := string(bodyBytes)
		if value.Responsevalue != bodyString {
			t.Fatalf("Returned SET Response not same as expected response\n")
		}
	}
}
