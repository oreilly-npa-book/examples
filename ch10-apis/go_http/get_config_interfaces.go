package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// helper method to implement the error-checking pattern
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	request, err := http.NewRequest("GET", "https://csr1/restconf/data/ietf-interfaces:interfaces", nil)
	checkError(err)
	request.Header.Set("Accept", "application/yang-data+json")
	request.Header.Add("Authorization", "Basic "+basicAuth("ntc", "ntc123"))

	result, err := client.Do(request)
	checkError(err)
	body, err := ioutil.ReadAll(result.Body)
	checkError(err)
	result.Body.Close()
	fmt.Printf("%s", body)
}
