package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/maxhille/elm-grpc/build"
	"google.golang.org/grpc"
)

// START OMIT
type server struct {
	items []string
}

func main() {
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, server{items: []string{"Mary", "Peter", "Bob", "Alice"}})

	// wrap gRPC server into 'normal' go http handlers
	webGrpc := grpcweb.WrapServer(s)
	http.HandleFunc("/example.SearchService/", webGrpc.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return nil, fmt.Errorf("not implemented yet")
}

// END OMIT
