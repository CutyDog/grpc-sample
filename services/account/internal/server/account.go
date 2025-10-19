package server

import (
	"context"

	accountv1 "github.com/CutyDog/grpc-sample/proto/gen/account/v1"
	"gorm.io/gorm"
)

type AccountServer struct {
	accountv1.UnimplementedAccountServiceServer
	db *gorm.DB
}

func NewAccountServer(db *gorm.DB) *AccountServer {
	return &AccountServer{db: db}
}

func (s *AccountServer) GetAccount(ctx context.Context, req *accountv1.GetAccountRequest) (*accountv1.GetAccountResponse, error) {
	// TODO: Implement GetAccount logic
	return &accountv1.GetAccountResponse{}, nil
}
