package server

import (
	"context"

	accountv1 "github.com/CutyDog/grpc-sample/proto/gen/account/v1"
	"github.com/CutyDog/grpc-sample/services/account/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServer struct {
	accountv1.UnimplementedAccountServiceServer
	repo repo.AccountRepository
}

func NewAccountServer(repository repo.AccountRepository) *AccountServer {
	return &AccountServer{repo: repository}
}

func (s *AccountServer) GetAccount(ctx context.Context, req *accountv1.GetAccountRequest) (*accountv1.GetAccountResponse, error) {
	var account *accountv1.Account

	switch key := req.Key.(type) {
	case *accountv1.GetAccountRequest_Id:
		// IDで検索
		acc, err := s.repo.FindByID(ctx, key.Id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get account: %v", err)
		}
		if acc == nil {
			return nil, status.Error(codes.NotFound, "account not found")
		}
		account = &accountv1.Account{
			Id:        int64(acc.ID),
			Uid:       acc.UID,
			CreatedAt: timestamppb.New(acc.CreatedAt),
			UpdatedAt: timestamppb.New(acc.UpdatedAt),
		}

	case *accountv1.GetAccountRequest_Uid:
		// UIDで検索
		acc, err := s.repo.FindByUID(ctx, key.Uid)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get account: %v", err)
		}
		if acc == nil {
			return nil, status.Error(codes.NotFound, "account not found")
		}
		account = &accountv1.Account{
			Id:        int64(acc.ID),
			Uid:       acc.UID,
			CreatedAt: timestamppb.New(acc.CreatedAt),
			UpdatedAt: timestamppb.New(acc.UpdatedAt),
		}

	default:
		return nil, status.Error(codes.InvalidArgument, "either id or uid must be provided")
	}

	return &accountv1.GetAccountResponse{
		Account: account,
	}, nil
}
