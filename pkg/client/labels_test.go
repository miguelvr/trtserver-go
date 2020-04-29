package client

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestGetLabels(t *testing.T) {
	testCases := []struct {
		Contents string
		Expected LabelMap
	}{
		{
			Contents: "a\nb\n",
			Expected: LabelMap{0: "a", 1: "b"},
		},
		{
			Contents: "aaaa\n12345\nxxxxx",
			Expected: LabelMap{0: "aaaa", 1: "12345", 2: "xxxxx"},
		},
	}

	for _, testCase := range testCases {
		t.Run(t.Name(), func(t *testing.T) {
			reader := strings.NewReader(testCase.Contents)
			labelMap, err := ReadLabels(reader)
			assert.NoError(t, err)

			isEqual := cmp.Equal(*labelMap, testCase.Expected)
			assert.True(t, isEqual)
		})
	}
}
