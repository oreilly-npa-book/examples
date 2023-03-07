package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// implementation of gRPC for Go maintained by Google
	"google.golang.org/grpc"

	// our own package that was built before, and updated to point
	// to "./networkstuff" via a replacement in the `go.mod` file
	pb "github.com/pkg/networkstuff"
)

// custom struct extending the pb.RouterServiceServer plus a
// local cache of Routers
type routerServiceServer struct {
	pb.UnimplementedRouterServiceServer
	localRouters []*pb.Router
}

// method necessary to match the interface definition for
// pb.RouterServiceServer, with the same signature
func (s *routerServiceServer) GetRouter(ctx context.Context,
	router_request *pb.RouterRequest) (*pb.Router, error) {
	// Use the local cache of Routers to match by the
	// router identifier
	for _, router := range s.localRouters {
		if router.Id == router_request.Id {
			return router, nil
		}
	}
	// No router was found, return a nameless router
	return &pb.Router{}, nil
}

// server contains the data to expose via grpc
var server = &routerServiceServer{
	localRouters: []*pb.Router{
		&pb.Router{
			Id:       1,
			Hostname: "Router A",
			Interfaces: []*pb.Interface{
				&pb.Interface{
					Id:          1000,
					Description: "Gi 0/0/0",
				},
				&pb.Interface{
					Id:          1001,
					Description: "Gi 0/0/1",
				},
			},
		},
		&pb.Router{
			Id:       2,
			Hostname: "Router B",
			Interfaces: []*pb.Interface{
				&pb.Interface{
					Id:          2000,
					Description: "Gi 0/0/0",
				},
				&pb.Interface{
					Id:          2001,
					Description: "Gi 0/0/1",
				},
			},
		},
	},
}

func main() {
	// Create a TCP server listener in 50051 port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Bootstrap a gRPC server with defaults
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register the custom RouterServiceServer implementation to
	// the gRPC server
	pb.RegisterRouterServiceServer(grpcServer, server)
	// Attach the gRPC server to the TCP port 50051 opened before
	grpcServer.Serve(lis)
}
