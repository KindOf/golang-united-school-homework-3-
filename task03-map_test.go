package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	tests := []struct {
		name       string
		input      map[int]string
		wantResult []string
	}{
		{
			name:       "should return sorted by key values slice",
			input:      map[int]string{2: "a", 0: "b", 1: "c"},
			wantResult: []string{"b", "c", "a"},
		}, {
			name:       "should return sorted by key values slice 2",
			input:      map[int]string{10: "aa", 0: "bb", 500: "cc"},
			wantResult: []string{"bb", "aa", "cc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := sortMapValues(tt.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
