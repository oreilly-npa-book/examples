package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/pkg/networkstuff"
)

type routerServiceServer struct {
	pb.UnimplementedRouterServiceServer
	localRouters []*pb.Router
}

func (s *routerServiceServer) GetRouter(ctx context.Context, router_request *pb.RouterRequest) (*pb.Router, error) {
	for _, router := range s.localRouters {
		if router.Id == router_request.Id {
			return router, nil
		}
	}
	// No router was found, return an unnamed router
	return &pb.Router{}, nil
}

var exampleData = []byte(`[
	{
		"interfaces": [
			{
				"id": 1000,
				"description": "Gi 0/0/0"
			},
			{
				"id": 1001,
				"description": "Gi 0/0/1"
			}
		],
		"hostname": "Router A",
		"id": 1
	},
	{
		"interfaces": [
			{
				"id": 2000,
				"description": "Gi 0/0/0"
			},
			{
				"id": 2001,
				"description": "Gi 0/0/1"
			}
		],
		"hostname": "Router B",
		"id": 2
	}
]`)

func newServer() *routerServiceServer {
	s := &routerServiceServer{}
	if err := json.Unmarshal(exampleData, &s.localRouters); err != nil {
		log.Fatalf("Failed to load default routers: %v", err)
	}

	return s
}

// func main() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	var opts []grpc.ServerOption
// 	grpcServer := grpc.NewServer(opts...)
// 	pb.RegisterRouterServiceServer(grpcServer, newServer())
// 	grpcServer.Serve(lis)
// }
