package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/CutyDog/grpc-sample/services/graphql/client"
	"github.com/CutyDog/grpc-sample/services/graphql/graph"
)

func main() {
	// Get configuration from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	accountServiceAddr := os.Getenv("ACCOUNT_SERVICE_ADDR")
	if accountServiceAddr == "" {
		accountServiceAddr = "account:9090"
	}

	// Initialize gRPC clients
	accountClient, err := client.NewAccountClient(accountServiceAddr)
	if err != nil {
		log.Fatalf("failed to create account client: %v", err)
	}
	defer accountClient.Close()

	// Initialize GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			AccountClient: accountClient,
		},
	}))

	// Setup HTTP handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL server ready at http://localhost:%s/", port)
	log.Printf("🎮 Playground available at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
