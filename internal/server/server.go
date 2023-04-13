package server

import (
	"context"

	pb "github.com/eliassebastian/r6index-recommendation/pkg/proto/server"
	"google.golang.org/grpc/status"
)

type RecommendationServer struct {
	pb.UnimplementedRecommendationServiceServer
}

func (s *RecommendationServer) Index(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	if in.GetId() == "" {
		return &pb.Response{}, status.Error(400, "id = empty player id")
	}

	return &pb.Response{
		Code:    200,
		Message: "OK",
	}, nil
}
