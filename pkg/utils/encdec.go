package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
)

func ConvertImageToArray(img image.Image) []float32 {
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

func EncodeArray(array []float32) ([]byte, error) {
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

// ReadBytes converts a slice of 4 bytes to float32 (assumes Little Endian)
func ReadBytes(fourBytes []byte, v interface{}) error {
	buf := bytes.NewBuffer(fourBytes)

	switch v.(type) {
	case *int8:
		v = v.(*int8)
	case *int16:
		v = v.(*int16)
	case *int32:
		v = v.(*int32)
	case *float32:
		v = v.(*float32)
	case *float64:
		v = v.(*float64)
	default:
		return fmt.Errorf("cannot decode bytes to type %T", v)
	}

	if err := binary.Read(buf, binary.LittleEndian, v); err != nil {
		return err
	}
	return nil
}
