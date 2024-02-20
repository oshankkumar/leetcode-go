package evaluaterpn

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tt := range tests {
		if got := EvalRPN(tt.tokens); got != tt.want {
			t.Errorf("EvalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
		}
	}
}
