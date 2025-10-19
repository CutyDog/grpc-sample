package graph

import (
	"github.com/CutyDog/grpc-sample/services/graphql/client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AccountClient *client.AccountClient
}
