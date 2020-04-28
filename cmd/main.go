package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	trtclient "github.com/miguelvr/trtserver-go/pkg/client"
	trtis "github.com/miguelvr/trtserver-go/pkg/gen/nvidia_inferenceserver"
	"google.golang.org/grpc"
)

const (
	URL       = "localhost:8001"
	modelName = "resnet50_imagenet"
)

func main() {
	imagePath := "assets/white_shark.jpg"

	// Connect to gRPC server
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to endpoint %s: %v", URL, err)
	}

	defer conn.Close()

	// Create gen from gRPC server connection
	grpcServiceClient := trtis.NewGRPCServiceClient(conn)
	client := trtclient.New(grpcServiceClient)

	liveHealthResponse := client.HealthRequest("live")
	fmt.Printf("TRTIS Health - Live: %v\n", liveHealthResponse.Health)

	statusResponse := client.StatusRequest(modelName)
	fmt.Printf("TRTIS Status: %v\n", statusResponse.ServerStatus.ModelStatus)

	img, err := imaging.Open(imagePath, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	inferReq, err := trtclient.EncodeRequest(img)
	if err != nil {
		log.Fatalf("failed to encode inference request: %v", err)
	}

	inferResponse, err := client.InferRequest(inferReq, modelName)
	if err != nil {
		log.Fatalf("error during inference request: %v", err)
	}

	label, prob, err := trtclient.DecodeResponse(inferResponse)
	if err != nil {
		log.Fatalf("could not decode response: %v", err)
	}

	fmt.Printf("Inference Response: {\"label\": %d, \"score\": %f}", label, prob)
}
