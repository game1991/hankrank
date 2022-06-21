package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stepTmp, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	steps, err := strconv.ParseInt(string(stepTmp), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}

	pathTmp, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}

	vally, graph := countingValleys(int32(steps), string(pathTmp))
	fmt.Println(vally)
	fmt.Println("explain:", graph)
}

func countingValleys(step int32, path string) (int32, []byte) {
	l, v := 0, 0
	result := make([]byte, int(step)*2)

	result[0] = '-'
	result = append(result, '\n')

	for _, i := range path {
		if i == 'D' {
			result[i+1] = '\\'
			l--
		} else if i == 'U' {
			result[i+1] = '/'
			l++
			if l == 1 {
				v++
			}
		}
	}

	result = append(result, '-')
	return int32(v)
}
