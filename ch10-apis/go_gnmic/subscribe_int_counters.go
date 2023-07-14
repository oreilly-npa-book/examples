package main

import (
	"context"
	"fmt"
	"github.com/openconfig/gnmic/api"
	"google.golang.org/protobuf/encoding/prototext"
	"log"
	"time"
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

	// create a gNMI subscribeRequest
	subReq, err := api.NewSubscribeRequest(
		api.Encoding("json_ietf"),
		api.SubscriptionListMode("stream"),
		api.Subscription(
			api.Path("/interfaces/interface/state/counters"),
			api.SubscriptionMode("sample"),
			api.SampleInterval(10*time.Second),
		))
	check_error(err)

	fmt.Println(prototext.Format(subReq))
	// start the subscription
	go tg.Subscribe(ctx, subReq, "sub1")
	// start a goroutine that will stop the subscription after x seconds
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(42 * time.Second):
			tg.StopSubscription("sub1")
		}
	}()
	subRspChan, subErrChan := tg.ReadSubscriptions()
	for {
		select {
		case rsp := <-subRspChan:
			fmt.Println(prototext.Format(rsp.Response))
		case tgErr := <-subErrChan:
			log.Fatalf("subscription %q stopped: %v", tgErr.SubscriptionName, tgErr.Err)
		}
	}

}
