package client

import (
	"bufio"
	"io"
)

type LabelMap map[int32]string

func ReadLabels(file io.Reader) (*LabelMap, error) {
	scanner := bufio.NewScanner(file)

	labelMap := LabelMap{}

	i := 0
	for scanner.Scan() {
		label := scanner.Text()
		labelMap[int32(i)] = label
		i++
	}

	return &labelMap, nil
}
