package aboutIOFile

import "testing"

var(
	file1="golang.txt"
	file2="other.txt"
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
	F2F(file1,file2)
}