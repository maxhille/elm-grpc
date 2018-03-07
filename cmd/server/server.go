package example

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/maxhille/elm-grpc/build"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// START OMIT
type server struct {
	items []string
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, server{items: []string{"Mary", "Peter", "Bob", "Alice"}})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return nil, fmt.Errorf("not implemented yet")
}

// END OMIT
