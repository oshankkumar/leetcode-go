package converttobase2

func basePos2(n int) string {
	if n == 0 {
		return "0"
	}

	lsb := "0"
	if n&1 == 1 {
		lsb = "1"
	}

	return basePos2(n>>1) + lsb
}

func Bits(n int64) string {
	var res string

	for i := 0; i < 64; i++ {
		res = Ter(n&1 == 0, "0", "1") + res
		n >>= 1
	}

	return res
}

func Ter[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
