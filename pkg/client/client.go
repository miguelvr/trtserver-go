package client

import (
	"context"
	"log"
	"time"

	trtis "github.com/miguelvr/trtserver-go/pkg/gen/nvidia_inferenceserver"
)

type Client struct {
	trtis.GRPCServiceClient
}

func New(client trtis.GRPCServiceClient) *Client {
	return &Client{client}
}

// mode should be either "live" or "ready"
func (c *Client) HealthRequest(mode string) *trtis.HealthResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create health request for given mode {"live", "ready"}
	healthRequest := trtis.HealthRequest{
		Mode: mode,
	}
	// Submit health request to server
	healthResponse, err := c.Health(ctx, &healthRequest)
	if err != nil {
		log.Fatalf("Couldn't get server health: %v", err)
	}
	return healthResponse
}

func (c *Client) StatusRequest(modelName string) *trtis.StatusResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create status request for a given model
	statusRequest := trtis.StatusRequest{
		ModelName: modelName,
	}
	// Submit status request to server
	statusResponse, err := c.Status(ctx, &statusRequest)
	if err != nil {
		log.Fatalf("Couldn't get server status: %v", err)
	}
	return statusResponse
}

func (c *Client) InferRequest(rawInput [][]byte, modelName string) (*trtis.InferResponse, error) {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create request header which describes inputs, outputs, and batch size
	inferRequestHeader := &trtis.InferRequestHeader{
		Input: []*trtis.InferRequestHeader_Input{
			{Name: "INPUT__0"},
		},
		Output: []*trtis.InferRequestHeader_Output{
			{Name: "OUTPUT__0"},
			{Name: "OUTPUT__1"},
		},
		BatchSize: 1,
	}

	// Create inference request for specific model/version
	inferRequest := trtis.InferRequest{
		ModelName:    modelName,
		ModelVersion: 1,
		MetaData:     inferRequestHeader,
		RawInput:     rawInput,
	}

	// Submit inference request to server
	inferResponse, err := c.Infer(ctx, &inferRequest)
	if err != nil {
		return nil, err
	}

	return inferResponse, nil
}
