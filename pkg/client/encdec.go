package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	trtis "github.com/miguelvr/trtserver-go/pkg/gen/nvidia_inferenceserver"
)

// EncodeRequest preprocesses and encodes an image into an array of bytes (assumes Little Endian)
func EncodeRequest(img image.Image) ([][]byte, error) {
	img = imaging.Resize(img, 224, 224, imaging.Linear)
	imgArray := convertImageToArray(img)
	imgRaw, err := encodeArray(imgArray)
	if err != nil {
		return nil, fmt.Errorf("error converting image to binarized array: %v", err)
	}

	return [][]byte{imgRaw}, nil
}

// DecodeResponse converts the InferResponse raw bytes into human readable data (assumes Little Endian)
func DecodeResponse(inferResponse *trtis.InferResponse) (label int32, prob float32, err error) {
	if inferResponse.RawOutput == nil {
		return 0, 0, fmt.Errorf("inference request failed: %v", inferResponse.RequestStatus)
	}

	var outputs [][]byte
	outputs = inferResponse.RawOutput
	outputBytes0 := outputs[0]
	outputBytes1 := outputs[1]

	label, err = readInt32(outputBytes0[:4])
	if err != nil {
		return 0, 0, err
	}

	prob, err = readFloat32(outputBytes1[:4])
	if err != nil {
		return 0, 0, err
	}

	return label, prob, nil
}

func convertImageToArray(img image.Image) []float32 {
	bounds := img.Bounds()
	w := bounds.Max.X
	h := bounds.Max.Y

	rawInput := make([]float32, 3*w*h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			rawInput[(i*w)+j] = float32(r>>8) / 255.0
			rawInput[(w*h)+(i*w)+j] = float32(g>>8) / 255.0
			rawInput[2*(w*h)+(i*w)+j] = float32(b>>8) / 255.0
		}
	}

	return rawInput
}

func encodeArray(array []float32) ([]byte, error) {
	var binaryArray []byte
	for _, v := range array {
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, err
		}
		binaryArray = append(binaryArray, buf.Bytes()...)
	}

	return binaryArray, nil
}

// readFloat32 converts a slice of 4 bytes to float32 (assumes Little Endian)
func readFloat32(fourBytes []byte) (float32, error) {
	buf := bytes.NewBuffer(fourBytes)
	var retval float32
	err := binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		return 0, err
	}
	return retval, nil
}

// readInt32 converts a slice of 4 bytes to int32 (assumes Little Endian)
func readInt32(fourBytes []byte) (int32, error) {
	buf := bytes.NewBuffer(fourBytes)
	var retval int32
	err := binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		return 0, err
	}
	return retval, nil
}
