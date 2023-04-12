package server

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/eliassebastian/r6index-recommendation/pkg/proto/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterRecommendationServiceServer(server, &RecommendationServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestRecommendationServiceServer_Index(t *testing.T) {
	// test cases
	// TODO: add more test cases
	tests := []struct {
		testName string
		req      *pb.Request
		res      *pb.Response
	}{
		{
			"typical player",
			&pb.Request{Id: "6844b415-aa94-43c9-8823-9389e4816902", Level: 211, Kost: 0.76, Rank: 35, RankPoints: 3424},
			&pb.Response{Code: 200, Message: "OK"},
		},
		{
			"another typical player",
			&pb.Request{Id: "460a3311-fe2f-489c-ba95-73370cbaddfa", Level: 448, Kost: 0.66, Rank: 35, RankPoints: 2344},
			&pb.Response{Code: 200, Message: "OK"},
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRecommendationServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			response, err := client.Index(ctx, tt.req)

			if response != nil {
				if response.GetCode() != tt.res.GetCode() {
					t.Error("response: expected", tt.res.GetCode(), "received", response.GetCode())
				}
			}

			if err != nil {
				// if er, ok := status.FromError(err); ok {
				// 	if er.Code() != tt.errCode {
				// 		t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
				// 	}
				// 	if er.Message() != tt.errMsg {
				// 		t.Error("error message: expected", tt.errMsg, "received", er.Message())
				// 	}
				// }
			}
		})
	}
}
