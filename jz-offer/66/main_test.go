package main

import (
	"reflect"
	"testing"
)

func Test_constructArr(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
