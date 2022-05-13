package homework

import "testing"

func Test_average(t *testing.T) {
	tests := []struct {
		name  string
		input [15]float32
		want  float32
	}{
		{
			name:  "Should return avarage 3.5",
			input: [15]float32{1, 2, 3, 4, 5, 6},
			want:  3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := average(tt.input); gotResult != tt.want {
				t.Errorf("average() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
