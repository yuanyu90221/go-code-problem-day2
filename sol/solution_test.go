package sol

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{
				head: ConvertList(&[]int{6, 5, 4, 3, 2, 1}),
				m:    2,
				n:    4,
			},
			want: ConvertList(&[]int{6, 3, 4, 5, 2, 1}),
		},
		{
			name: "Example2",
			args: args{
				head: ConvertList(&[]int{6, 5, 4, 3, 2, 1}),
				m:    2,
				n:    6,
			},
			want: ConvertList(&[]int{6, 1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
