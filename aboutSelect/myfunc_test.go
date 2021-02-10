package aboutSelect

import (
	"testing"
)

func TestMySelect(t *testing.T) {
	MySelect()
}
func BenchmarkMySelect(b *testing.B) {

	//与测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
		MySelect()

	}
	b.StopTimer()
	//与测试无关的代码
}
