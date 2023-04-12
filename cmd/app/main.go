package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/eliassebastian/r6index-recommendation/internal/server"
	pb "github.com/eliassebastian/r6index-recommendation/pkg/proto/server"
	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("---Recommendation Service Starting---")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(grpcServer, &server.RecommendationServer{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-ctx.Done()
		log.Printf("got signal %v, attempting graceful shutdown", ctx.Err())
		stop()

		grpcServer.GracefulStop()

		// Wait for 5 seconds before forceful shutdown
		<-time.After(5 * time.Second)
		wg.Done()
	}()

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}

	wg.Wait()
	log.Println("clean shutdown")
}
