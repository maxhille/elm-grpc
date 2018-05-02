package main

import (
	"log"
	"net"
	"net/http"
	"strings"

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

	log.Println("starting webserver at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END GRPCWEB OMIT

// START SERVICE OMIT
func (s service) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Printf("handling search request: %q", req.Query)

	ns := []string{"mary", "peter", "bob", "alice", "mark", "john", "karen"}
	result := []string{}

	for _, n := range ns {
		if req.Query != "" && strings.Contains(n, req.Query) {
			result = append(result, n)
		}
	}

	res := pb.SearchResponse{Items: result}
	return &res, nil
}

// END SERVICE OMIT
