package homework

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name       string
		input      []int64
		wantResult []int64
	}{
		{
			name:       "should return reversed slice",
			input:      []int64{1, 2, 5, 15},
			wantResult: []int64{15, 5, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reverse(tt.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
