package main

import (
	"context"
	"fmt"
	"log"

	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/gnmic/api"
	"google.golang.org/protobuf/encoding/prototext"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// create a target
	tg, err := api.NewTarget(
		api.Name("gnmi example"),
		api.Address("clab-demo-eos-spine1:6030"),
		api.Username("admin"),
		api.Password("admin"),
		api.Insecure(true),
	)
	check_error(err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a gNMI client
	err = tg.CreateGNMIClient(ctx)
	check_error(err)
	defer tg.Close()

	// create a GetRequest
	getReq, err := api.NewGetRequest(
		api.Path("/interfaces/interface/config/description"),
		api.Encoding("json_ietf"))
	check_error(err)
	fmt.Println(prototext.Format(getReq))

	// send the created gNMI GetRequest to the created target
	getResp, err := tg.Get(ctx, getReq)
	check_error(err)
	fmt.Println(prototext.Format(getResp))
	descriptionGnmiValue := getResp.GetNotification()[0].GetUpdate()[0].GetVal()
	//fmt.Printf("%`v", descriptionGnmiValue)
	myCurrentDescriptionValue, err := value.ToScalar(descriptionGnmiValue)
	check_error(err)
	myCurrentDescriptionStr := myCurrentDescriptionValue.(string)
	fmt.Println("This is my current description: " + myCurrentDescriptionStr)
	myNewDescription := myCurrentDescriptionStr + "something_else"
	fmt.Println("This is going to be the new one: " + myNewDescription)
}
