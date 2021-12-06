package main

import (
	"testing"
)
import "math/rand"

//func init() {
//	rand.Seed(time.Now().Unix())
//}

func BenchmarkStd(b *testing.B) {
	a := rand.Int31() //nolint:gosec
	var _ int
	for i := 0; i < b.N; i++ {
		_ = ^a
	}
}

func BenchmarkMY(b *testing.B) {
	a := rand.Int31() //nolint:gosec
	var _ int
	for i := 0; i < b.N; i++ {
		_ = -a - 1
	}
}
