package main

import (
	"log"
	"net"

	"github.com/eliassebastian/r6index-recommendation/internal/server"
	pb "github.com/eliassebastian/r6index-recommendation/pkg/proto/server"
	"google.golang.org/grpc"
)

func main() {
	// ...
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(grpcServer, &server.RecommendationServer{})
	log.Fatalln(grpcServer.Serve(listener))

}
