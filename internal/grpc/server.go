package grpc

import (
	"context"
	"net"

	"github.com/Egorpalan/api-pvz/internal/grpc/pvz_v1"
	"github.com/Egorpalan/api-pvz/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pvz_v1.UnimplementedPVZServiceServer
	pvzUsecase *usecase.PVZUsecase
}

func NewServer(pvzUC *usecase.PVZUsecase) *Server {
	return &Server{pvzUsecase: pvzUC}
}

func (s *Server) GetPVZList(ctx context.Context, _ *pvz_v1.GetPVZListRequest) (*pvz_v1.GetPVZListResponse, error) {
	pvzs, err := s.pvzUsecase.GetAll()
	if err != nil {
		return nil, err
	}

	var result []*pvz_v1.PVZ
	for _, p := range pvzs {
		result = append(result, &pvz_v1.PVZ{
			Id:               p.ID,
			RegistrationDate: timestamppb.New(p.RegistrationDate),
			City:             p.City,
		})
	}

	return &pvz_v1.GetPVZListResponse{Pvzs: result}, nil
}

func RunGRPCServer(pvzUC *usecase.PVZUsecase) error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pvz_v1.RegisterPVZServiceServer(s, NewServer(pvzUC))

	return s.Serve(lis)
}
