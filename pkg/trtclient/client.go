package trtclient

import (
	"context"
	"log"
	"time"

	trtis "github.com/miguelvr/trtserver-go/pkg/trtclient/nvidia_inferenceserver"
)

type Client struct {
	grpcClient trtis.GRPCServiceClient
	config     Config
}

type Config struct {
	RequestHeader *trtis.InferRequestHeader
	ModelName     string
	ModelVersion  int64
	EncoderFunc   EncoderFunc
	DecoderFunc   DecoderFunc
}

type EncoderFunc func(interface{}) ([][]byte, error)

type DecoderFunc func(*trtis.InferResponse) (interface{}, error)

func New(client trtis.GRPCServiceClient, config Config) *Client {
	return &Client{client, config}
}

// mode should be either "live" or "ready"
func (c *Client) CheckHealth(mode string) *trtis.HealthResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create health request for given mode {"live", "ready"}
	healthRequest := trtis.HealthRequest{
		Mode: mode,
	}
	// Submit health request to server
	healthResponse, err := c.grpcClient.Health(ctx, &healthRequest)
	if err != nil {
		log.Fatalf("Couldn't get server health: %v", err)
	}
	return healthResponse
}

func (c *Client) GetStatus() *trtis.StatusResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create status request for a given model
	statusRequest := trtis.StatusRequest{
		ModelName: c.config.ModelName,
	}
	// Submit status request to server
	statusResponse, err := c.grpcClient.Status(ctx, &statusRequest)
	if err != nil {
		log.Fatalf("Couldn't get server status: %v", err)
	}
	return statusResponse
}

func (c *Client) Inference(input interface{}) (interface{}, error) {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rawInput, err := c.config.EncoderFunc(input)
	if err != nil {
		return nil, err
	}

	// Create inference request for specific model/version
	inferRequest := trtis.InferRequest{
		ModelName:    c.config.ModelName,
		ModelVersion: c.config.ModelVersion,
		MetaData:     c.config.RequestHeader,
		RawInput:     rawInput,
	}

	// Submit inference request to server
	inferResponse, err := c.grpcClient.Infer(ctx, &inferRequest)
	if err != nil {
		return nil, err
	}

	response, err := c.config.DecoderFunc(inferResponse)
	if err != nil {
		return nil, err
	}

	return response, nil
}
