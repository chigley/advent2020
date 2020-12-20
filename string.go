package advent2020

func MinStr(x, y string) string {
	if x < y {
		return x
	}
	return y
}

func Reverse(s string) string {
	ret := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ret[len(s)-i-1] = s[i]
	}
	return string(ret)
}
