package aboutPanicAndRecover

import "testing"

func TestPanic(t *testing.T)  {
	var i, j int = 1, 2
	trouble(i)
	trouble(j)
}