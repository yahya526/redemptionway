package util

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func TestUtil(t *testing.T) {
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
