package utils

import "testing"

func TestCalculateUpdatedSize(t *testing.T) {
	type args struct {
		value      int
		percentage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateUpdatedSize(tt.args.value, tt.args.percentage); got != tt.want {
				t.Errorf("CalculateUpdatedSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
