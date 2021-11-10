package main

import "testing"

func Test_hasPath(t *testing.T) {
	type args struct {
		matrix [][]byte
		str    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{
				matrix: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				str: "BCCE",
			},
			want: true,
		},
		{
			name: "Case2",
			args: args{
				matrix: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				str: "ASAE",
			},
			want: false,
		},
		{
			name: "Case3",
			args: args{
				matrix: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				str: "ABCEFSADEESE",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.matrix, tt.args.str); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
