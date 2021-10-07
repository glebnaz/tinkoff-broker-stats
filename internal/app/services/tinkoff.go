package services

import (
	"context"

	pb "github.com/glebnaz/tinkoff-broker-stats/pkg/pb/api/v1"
	"github.com/glebnaz/tinkoff-broker-stats/pkg/tinkoff"
)

type Service struct {
	api tinkoff.API
}

func (s Service) GetAccounts(ctx context.Context, request *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	acs, err := s.api.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.GetAccountsResponse_Account, len(acs))

	for i := range acs {
		account := pb.GetAccountsResponse_Account{
			Id:   acs[i].ID,
			Type: string(acs[i].Type),
		}

		res[i] = &account
	}

	return &pb.GetAccountsResponse{List: res}, nil
}

func NewService(api tinkoff.API) *Service {
	return &Service{api: api}
}
