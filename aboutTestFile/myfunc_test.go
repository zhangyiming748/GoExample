package aboutTestFile

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestMySquare(t *testing.T) {
	inputs := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := [...]int{1, 4, 9, 16, 25, 36, 49, 64, 81}

	for i := 0; i < len(inputs); i++ {
		ret := MySquare(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d\texpected is %d\tbut actual is %d\n", inputs[i], expected[i], ret)
		}
	}

}
func TestMyCube(t *testing.T) {
	inputs := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := [...]int{1, 8, 27, 64, 125, 216, 343, 512, 729}

	for i := 0; i < len(inputs); i++ {
		ret := MyCube(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d\texpected is %d\tbut actual is %d\n", inputs[i], expected[i], ret)
		}
	}
}
func BenchmarkMySquare(b *testing.B) {
	rand.Seed(time.Now().Unix())
	//与测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
		j := rand.Int()
		ret := MyCube(j)
		log.Printf("result is %d", ret)
	}
	b.StopTimer()
	//与测试无关的代码
}
func BenchmarkMyCube(b *testing.B) {
	rand.Seed(time.Now().Unix())
	//与测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
		j := rand.Int()
		ret := MyCube(j)
		log.Printf("result is %d", ret)
	}
	b.StopTimer()
	//与测试无关的代码
}
