package main

import (
  "fmt"
  "encoding/base64"
  "io/ioutil"
  "net/http"
  "crypto/tls"
)

func basicAuth(username, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
  transCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
  }
  client := &http.Client{Transport: transCfg}
  request, _ := http.NewRequest("GET", "https://csr1/restconf/data/ietf-interfaces:interfaces", nil)
  request.Header.Set("Accept", "application/yang-data+json")
  request.Header.Add("Authorization","Basic " + basicAuth("ntc","ntc123"))

  result, err := client.Do(request)
  if (err!=nil){
    fmt.Printf("%s",err)
  }
  body, _ := ioutil.ReadAll(result.Body)
  result.Body.Close()
  fmt.Printf("%s", body)
}
