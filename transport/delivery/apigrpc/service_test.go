package apigrpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"net"
	"testing"

	pb "github.com/ArenDjango/golang-test-task/protos/rates"
	"github.com/stretchr/testify/assert"
)

type exchangeService struct {
	pb.UnimplementedRatesServiceServer
	api GarantexAPI
}

func (s *exchangeService) GetRates(ctx context.Context, req *pb.GetRatesRequest) (*pb.GetRatesResponse, error) {
	return s.api.GetRates(ctx)
}

type GarantexAPI interface {
	GetRates(ctx context.Context) (*pb.GetRatesResponse, error)
}

type MockGarantexAPI struct{}

func (m *MockGarantexAPI) GetRates(ctx context.Context) (*pb.GetRatesResponse, error) {
	return &pb.GetRatesResponse{
		AskPrice:  1.23,
		BidPrice:  1.22,
		Timestamp: "2024-05-24T00:00:00Z",
	}, nil
}

func startTestGRPCServer(t *testing.T, srv pb.RatesServiceServer) (*grpc.Server, string) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRatesServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	return s, lis.Addr().String()
}

func TestGRPCServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPI := &MockGarantexAPI{}
	server := &exchangeService{
		api: mockAPI,
	}

	grpcServer, addr := startTestGRPCServer(t, server)
	defer grpcServer.Stop()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRatesServiceClient(conn)

	req := &pb.GetRatesRequest{}
	resp, err := client.GetRates(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 1.23, resp.GetAskPrice())
	assert.Equal(t, 1.22, resp.GetBidPrice())
	assert.Equal(t, "2024-05-24T00:00:00Z", resp.GetTimestamp())
}
