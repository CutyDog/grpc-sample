package client

import (
	"context"
	"fmt"

	accountv1 "github.com/CutyDog/grpc-sample/proto/gen/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AccountClient struct {
	client accountv1.AccountServiceClient
	conn   *grpc.ClientConn
}

func NewAccountClient(addr string) (*AccountClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to account service: %w", err)
	}

	client := accountv1.NewAccountServiceClient(conn)

	return &AccountClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *AccountClient) Close() error {
	return c.conn.Close()
}

func (c *AccountClient) GetAccountByUID(ctx context.Context, uid string) (*accountv1.Account, error) {
	resp, err := c.client.GetAccount(ctx, &accountv1.GetAccountRequest{
		Key: &accountv1.GetAccountRequest_Uid{
			Uid: uid,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return resp.Account, nil
}
