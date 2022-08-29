package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	fmt.Println(count)
	element1 := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	element2 := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < count; i++ {
		_, err := fmt.Fprintf(writer, "%s\n", element1[i])
		if err != nil {
			return 
		}
	}

	for i := 0; i < count; i++ {
		_, err := fmt.Fprintf(writer, "%s\n", element2[i])
		if err != nil {
			return
		}
	}

	err := writer.Flush()
	if err != nil {
		return 
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
