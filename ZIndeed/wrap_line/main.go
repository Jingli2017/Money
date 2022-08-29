package main

import (
	"fmt"
	"strings"
)

//e.g. ["1p3acres", "is", "a", "good", "place", "to", "communicate"], 12
//=> {"1p3acres-is", "a-good-place", "for", "communicate"}

func WrapLine(words []string, limit int) []string {
	output := make([]string, 0)
	index := 0

	for index < len(words) {
		count := len(words[index])
		last := index + 1
		for last < len(words) && count + 1 + len(words[last]) <= limit {
			count += 1 + len(words[last])
			last++
		}

		var b strings.Builder
		for i := index; i < last-1; i++ {
			b.WriteString(words[i])
			b.WriteString("-")
		}
		b.WriteString(words[last-1])
		output = append(output, b.String())

		index = last
	}

	return output
}

//e.g. ["123 45 67 8901234 5678", "12345 8 9 0 1 23"], 10 => {"123--45-67", "8901234", "5678-12345", "8-9-0-1-23"}
//["123 45 67 8901234 5678", "12345 8 9 0 1 23"], 15 => {"123----45----67", "8901234----5678", "12345---8--9--0", "23"}
func main() {
	res_1 := WrapLineJustify([]string{"123", "45", "67", "8901234", "5678", "12345", "8", "9", "0", "1", "23"}, 10)
	fmt.Println(res_1)
	res_2 := WrapLineJustify([]string{"123", "45", "67", "8901234", "5678", "12345", "8", "9", "0", "1", "23"}, 15)
	fmt.Println(res_2)
}

func WrapLineJustify(words []string, limit int) []string {
	output := make([]string, 0)
	index := 0

	for index < len(words) {
		count := len(words[index])
		last := index + 1
		for last < len(words) && count + 1 + len(words[last]) <= limit {
			count += 1 + len(words[last])
			last++
		}

		var b strings.Builder

		gap := last - index - 1
		space := 0
		rest := 0
		if gap != 0 {
			space = (limit - count) / gap
			rest = (limit - count) % gap
		}

		for i := index; i < last-1; i++ {
			b.WriteString(words[i])
			b.WriteString("-")
			for j := 0; j < space; j++ {
				b.WriteString("-")
			}
			if rest != 0 {
				b.WriteString("-")
				rest--
			}
		}
		b.WriteString(words[last-1])

		output = append(output, b.String())
		index = last
	}

	return output
}