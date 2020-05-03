package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/miguelvr/trtserver-go/pkg/trtclient"
	trtis "github.com/miguelvr/trtserver-go/pkg/trtclient/nvidia_inferenceserver"
	"github.com/miguelvr/trtserver-go/pkg/utils"
	"google.golang.org/grpc"
)

const (
	URL          = "localhost:8001"
	modelName    = "resnet50_imagenet"
	modelVersion = 1
	labelMapFile = "assets/imagenet_labels.txt"
)

type ResNetOutput struct {
	Label       int32
	Probability float32
}

func encodeRequest(input interface{}) ([][]byte, error) {
	img := input.(image.Image)
	img = imaging.Resize(img, 224, 224, imaging.Linear)
	imgArray := utils.ConvertImageToArray(img)
	imgRaw, err := utils.EncodeArray(imgArray)
	if err != nil {
		return nil, fmt.Errorf("error converting image to binarized array: %v", err)
	}

	return [][]byte{imgRaw}, nil
}

func decodeResponse(inferResponse *trtis.InferResponse) (interface{}, error) {
	if inferResponse.RawOutput == nil {
		return nil, fmt.Errorf("inference request failed: %v", inferResponse.RequestStatus)
	}

	var outputs [][]byte
	outputs = inferResponse.RawOutput
	outputBytes0 := outputs[0]
	outputBytes1 := outputs[1]

	var label int32
	err := utils.ReadBytes(outputBytes0[:4], &label)
	if err != nil {
		return nil, err
	}

	var prob float32
	err = utils.ReadBytes(outputBytes1[:4], &prob)
	if err != nil {
		return nil, err
	}

	return ResNetOutput{Label: label, Probability: prob}, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Usage: predict [IMAGE_PATH]")
	}

	imagePath := args[0]

	r, err := os.Open(labelMapFile)
	if err != nil {
		log.Fatal(err)
	}

	labelMap, err := utils.ReadLabels(r)
	if err != nil {
		log.Fatalf("error reading label map file: %v", err)
	}

	// Connect to gRPC server
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to endpoint %s: %v", URL, err)
	}

	defer conn.Close()

	clientConfig := trtclient.Config{
		RequestHeader: &trtis.InferRequestHeader{
			Input: []*trtis.InferRequestHeader_Input{
				{Name: "INPUT__0"},
			},
			Output: []*trtis.InferRequestHeader_Output{
				{Name: "OUTPUT__0"},
				{Name: "OUTPUT__1"},
			},
			BatchSize: 1,
		},
		ModelName:    modelName,
		ModelVersion: modelVersion,
		EncoderFunc:  encodeRequest,
		DecoderFunc:  decodeResponse,
	}
	// Create gen from gRPC server connection
	grpcServiceClient := trtis.NewGRPCServiceClient(conn)
	client := trtclient.New(grpcServiceClient, clientConfig)

	liveHealthResponse := client.CheckHealth("live")
	fmt.Printf("TRTIS Health - Live: %v\n", liveHealthResponse.Health)

	statusResponse := client.GetStatus()
	fmt.Printf("TRTIS Status: %v\n", statusResponse.ServerStatus.ModelStatus)

	img, err := imaging.Open(imagePath, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image %s", imagePath)
	}

	response, err := client.Inference(img)
	if err != nil {
		log.Fatalf("error during inference request: %v", err)
	}

	output := response.(ResNetOutput)
	fmt.Printf(
		"Inference Response: {\"label\": %s, \"label_id\": %d, \"score\": %f}\n",
		(*labelMap)[output.Label], output.Label, output.Probability,
	)
}
