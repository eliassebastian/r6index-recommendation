package server

import (
	"context"

	pb "github.com/eliassebastian/r6index-recommendation/pkg/proto/server"
)

type RecommendationServer struct {
	pb.UnimplementedRecommendationServiceServer
}

func (s *RecommendationServer) Index(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	return &pb.Response{
		Code:    200,
		Message: "OK",
	}, nil
}
