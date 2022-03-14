package main

import (
	"encoding/json"
	"fmt"
  "io/ioutil"
)

/*

  These structs are constructed to mimic our JSON data. Pay close attention
  to the datatypes used here, as well as the "json" mappings to the right -
  these are responsible for instructing the compiler which JSON structures
  map to which struct properties.

*/
type Address struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postalCode"`
}

type PhoneNumber struct {
	Phone_type string `json:"phone_type"`
	Number     string `json:"number"`
}

type JsonData struct {
	FirstName    string              `json:"firstName"`
	LastName     string              `json:"lastName"`
	IsAlive      bool                `json:"isAlive"`
	Age          int                 `json:"age"`
	Address      map[string]string   `json:"address"`
	PhoneNumbers []map[string]string `json:"phoneNumbers"`
	Children     []string            `json:"children"`
	Spouse       string              `json:"spouse"`
}

func main() {

  // We read our JSON file into a variable called "json_data"
	json_data, err := ioutil.ReadFile("data.json")
  if err != nil {
    panic(err)
  }

  // We create an instance of our JsonData struct
  var dat JsonData

  // The Unmarshal function attempts to map JSON data to our struct
	err = json.Unmarshal(json_data, &dat)
	if err != nil {
		panic(err)
	}

  // Sit back and watch the output! :)
  fmt.Printf("Name: %s, %s\n", dat.LastName, dat.FirstName)
  fmt.Println("Age: ", dat.Age)
  fmt.Printf("Located at: %s\n", dat.Address["streetAddress"])
  fmt.Println("Now listing phone numbers")
  for i := range dat.PhoneNumbers {
    fmt.Printf(
      " -- %s's %s phone number is %s\n",
      dat.FirstName,
      dat.PhoneNumbers[i]["phone_type"],
      dat.PhoneNumbers[i]["number"],
    )
  }
}
