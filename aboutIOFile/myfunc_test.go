package aboutIOFile

import "testing"

var (
	file1 = "golang.txt"
	file2 = "other.txt"
)

func TestCreate(t *testing.T) {
	Create(file1)
}
func TestAppending(t *testing.T) {
	Appending(file1)
}
func TestReadAndAppend(t *testing.T) {
	ReadAndAppend(file1)
}
func TestF2F(t *testing.T) {
	F2F(file1, file2)
}
func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Create(file1)
	}
}
func BenchmarkAppending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Appending(file1)
	}
}
func BenchmarkReadAndAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAndAppend(file1)
	}
}
func BenchmarkF2F(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F2F(file1, file2)
	}
}
