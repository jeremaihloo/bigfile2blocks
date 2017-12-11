package cores_test

import (
	"math"
	"testing"
)

func TestBlockNum(t *testing.T) {
	fileSize := 38000
	blockSize := 123
	num := math.Ceil(float64(fileSize / blockSize))
	t.Log(num)
}
