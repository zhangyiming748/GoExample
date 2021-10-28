package interview

import "testing"

func isString(s string) bool {
	sb := []byte(s)
	start := 0
	end := len(sb) - 1
	if len(s) == 1 {
		return true
	}
	for start != end {
		if sb[start] == sb[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
func TestUnit(t *testing.T) {
	s1 := "ababc"
	s2 := "aba"
	ret1 := isString(s1)
	ret2 := isString(s2)
	t.Logf("%v\n%v\n", ret1, ret2)
}
