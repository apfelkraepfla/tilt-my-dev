package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	pb "github.com/apfelkraepfla/tilt-my-def/protos/quotespb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func TestIntegrationScenario(t *testing.T) {
	portStr := os.Getenv("MYAPP_PORT")

	// Create a gRPC client connection
	conn, err := grpc.Dial("localhost:"+portStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewQuoteServiceClient(conn)

	// Call the GetQuote function
	response, err := client.GetQuote(context.Background(), &pb.QuoteRequest{})
	if err != nil {
		t.Fatalf("Error calling GetQuote: %v", err)
	}

	// Check if the response is not nil and has a non-empty quote
	if response == nil || response.Quote == "" {
		t.Errorf("Unexpected response: %v", response)
	}
}
