package aboutReadFile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Readlink(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		lines = append(lines, string(a))

	}
	return lines
}
