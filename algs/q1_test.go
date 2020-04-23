package algs

import "testing"

func TestFind2ndGreatest(t *testing.T) {
	type args struct {
		l []int
	}
	tests := q1.wants
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find2ndGreatest(tt.input); got != tt.output {
				t.Errorf("Find2ndGreatest() = %v, want %v", got, tt.output)
			}
		})
	}
}
