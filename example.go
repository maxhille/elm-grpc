package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/context"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/maxhille/elm-grpc/build"
	"google.golang.org/grpc"
)

type service struct{}

func main() {
	bindGRPCWEB()
}

// START GRPC OMIT
func bindGRPC() {
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, service{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// END GRPC OMIT

// START GRPCWEB OMIT
func bindGRPCWEB() {
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, service{})

	// wrap gRPC server into 'normal' go http handlers
	webGrpc := grpcweb.WrapServer(s)
	http.HandleFunc("/SearchService/", webGrpc.ServeHTTP)

	// serve index.html
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END GRPCWEB OMIT

// START SERVICE OMIT
func (s service) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	// items: []string{"Mary", "Peter", "Bob", "Alice"}
	log.Printf("handling search request: %q", req.Query)
	return nil, fmt.Errorf("not implemented yet")
}

// END SERVICE OMIT
