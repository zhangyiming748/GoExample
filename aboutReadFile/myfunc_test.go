package aboutReadFile

import (
	"log"
	"testing"
)

func TestReadlink(t *testing.T) {
	testf := "tf.txt"
	for i, v := range Readlink(testf) {
		log.Printf("第%d行的文本是%s\n", i, v)
	}
}
