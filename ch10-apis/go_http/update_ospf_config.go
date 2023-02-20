package main

import (
  "fmt"
  "encoding/base64"
  "io/ioutil"
  "net/http"
  "crypto/tls"
  "bytes"
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
  var jsonStr = []byte(`{
  "router-ospf": {
    "ospf": {
      "process-id": [
        {
          "id": 10,
          "network": [
            {
              "ip": "203.0.113.0",
              "wildcard": "0.0.0.7",
              "area": 0
            },
            {
              "ip": "203.0.113.64",
              "wildcard": "0.0.0.7",
              "area": 0
            }
          ],
          "router-id": "203.0.113.1"
        }
      ]
    }
  }
}'`)
  request, _ := http.NewRequest("PUT", "https://csr1/restconf/data/Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf", bytes.NewBuffer(jsonStr))
  request.Header.Set("Accept", "application/yang-data+json")
  request.Header.Set("Authorization","Basic " + basicAuth("ntc","ntc123"))
  request.Header.Set("Content-Type", "application/yang-data+json")

  result, err := client.Do(request)
  if (err!=nil){
    fmt.Printf("%s",err)
  }
  body, _ := ioutil.ReadAll(result.Body)
  result.Body.Close()
  fmt.Printf("%s", body)
}
