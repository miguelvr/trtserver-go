package utils

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBytes(t *testing.T) {
	cases := []struct {
		name  string
		t     string
		value interface{}
	}{
		{
			name:  `1 (int8)`,
			value: int8(1),
		},
		{
			name:  `-1 (int8)`,
			value: int8(-1),
		},
		{
			name:  `1 (int16)`,
			value: int16(1),
		},
		{
			name:  `-1 (int16)`,
			value: int16(-1),
		},
		{
			name:  `1 (int32)`,
			value: int32(1),
		},
		{
			name:  `-1 (int32)`,
			value: int32(-1),
		},
		{
			name:  `1 (float32)`,
			value: float32(1),
		},
		{
			name:  `-1 (float32)`,
			value: float32(-1),
		},
		{
			name:  `1 (float64)`,
			value: float64(1),
		},
		{
			name:  `-1 (float64)`,
			value: float64(-1),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := binary.Write(&buf, binary.LittleEndian, tt.value)
			assert.NoError(t, err, "could not encode value to buffer")

			switch tt.value.(type) {
			case int8:
				var retVal int8
				err = ReadBytes(buf.Bytes(), &retVal)
				assert.NoError(t, err)
				assert.Equal(t, tt.value, retVal)
			case int16:
				var retVal int16
				err = ReadBytes(buf.Bytes(), &retVal)
				assert.NoError(t, err)
				assert.Equal(t, tt.value, retVal)
			case int32:
				var retVal int32
				err = ReadBytes(buf.Bytes(), &retVal)
				assert.NoError(t, err)
				assert.Equal(t, tt.value, retVal)
			case float32:
				var retVal float32
				err = ReadBytes(buf.Bytes(), &retVal)
				assert.NoError(t, err)
				assert.Equal(t, tt.value, retVal)
			case float64:
				var retVal float64
				err = ReadBytes(buf.Bytes(), &retVal)
				assert.NoError(t, err)
				assert.Equal(t, tt.value, retVal)
			}
		})
	}
}
