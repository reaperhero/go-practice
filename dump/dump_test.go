package dump

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

// rum demo:
//     go run ./dump/_examples/slice.go
func TestDump(t *testing.T) {
	dump.P(
		nil, true,
		12, int8(12), int16(12), int32(12), int64(12),
		uint(22), uint8(22), uint16(22), uint32(22), uint64(22),
		float32(23.78), float64(56.45),
		'c', byte('d'),
		"string",
		[]byte("abc"),
		[]int{1, 2, 3},
		[]string{"ab", "cd"},
		[]interface{}{
			"ab",
			234,
			[]int{1, 3},
			[]string{"ab", "cd"},
		},
	)
}
