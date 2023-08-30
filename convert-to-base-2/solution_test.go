package converttobase2

import "testing"

func Test_basePos2(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{10, "01010"},
		{2, "010"},
		{0, "0"},
		{31, "011111"},
	}
	for _, tt := range tests {
		if got := basePos2(tt.n); got != tt.want {
			t.Errorf("basePos2(%d) = %s, want %s", tt.n, got, tt.want)
		}
	}
}

func TestNegNum(t *testing.T) {
	t.Log(Bits(0x0F))
	t.Log(Bits(-0x0F))
}
