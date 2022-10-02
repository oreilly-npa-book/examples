package main

import (
	"context"
	"fmt"

	"github.com/karimra/gnmic/api"
	"google.golang.org/protobuf/encoding/prototext"

)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Device struct {
	Hostname   string `yaml:"hostname"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Insecure   bool   `yaml:"insecure"`
}

type Data struct {
	Prefix   string `yaml:"prefix,omitempty"`
	Path     string `yaml:"path"`
	Encoding string `yaml:"encoding,omitempty"`
	Value    string `yaml:"value,omitempty"`
}


func main() {
  tg, err := api.NewTarget(
		api.Name("gnmi"),
		api.Address("clab-demo-eos-spine1:6030"),
		api.Username("admin"),
		api.Password("admin"),
		api.Insecure(true),
	)
  check(err)

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  err = tg.CreateGNMIClient(ctx)
  check(err)
  defer tg.Close()

  getReq, err := api.NewGetRequest(
    api.Path("/"),
    api.Encoding("json_ietf"))
  check(err)

  getResp, err := tg.Get(ctx, getReq)
  check(err)

  fmt.Println(prototext.Format(getResp))
}

