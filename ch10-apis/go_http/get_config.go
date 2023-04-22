package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// helper method to construct the expected format of
// the basic authentication string, encoded in base64
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// helper method to implement the error-checking pattern
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	transCfg := &http.Transport{
		// ignore expired SSL certificates
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// create a new http client, with the previous defined transport config
	client := &http.Client{Transport: transCfg}

	// create a new http request, with the method, url, and headers
	request, err := http.NewRequest("GET",
		"https://csr1/restconf/data/Cisco-IOS-XE-native:native", nil)
	checkError(err)
	request.Header.Set("Accept", "application/yang-data+json")
	request.Header.Add("Authorization", "Basic "+basicAuth("ntc", "ntc123"))

	// perform the HTTP request, defined before, and store it in `result`
	result, err := client.Do(request)
	checkError(err)
	// read the body content from the response
	body, err := ioutil.ReadAll(result.Body)
	checkError(err)
	result.Body.Close()
	fmt.Printf("%s", body)
}
