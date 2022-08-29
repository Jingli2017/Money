package main

import (
	"strings"
)

/*
Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
 */
func main() {
}

func fullJustify(words []string, maxWidth int) []string {
	output := make([]string, 0)
	index := 0
	for index < len(words) {
		totalChars := len(words[index])
		last := index + 1

		for last < len(words) && totalChars + len(words[last]) + 1 <= maxWidth {
			totalChars += len(words[last]) + 1 // + 1 for space
			last++
		}

		gaps := last - 1 - index // now last is the start point of next iteration

		var b strings.Builder
		// last line or just one word in one line
		if last == len(words) || gaps == 0 {
			for i := index; i < last-1; i++ {
				b.WriteString(words[i])
				b.WriteString(" ")
			}
			b.WriteString(words[last-1])
			for i := 0; i < maxWidth - totalChars; i++ {
				b.WriteString(" ")
			}
		}else{
			space := (maxWidth - totalChars) / gaps
			rest := (maxWidth - totalChars) % gaps

			for i := index; i < last-1; i++ {
				b.WriteString(words[i])
				b.WriteString(" ")
				for j := 0; j < space; j++ {
					b.WriteString(" ")
				}
				if rest != 0 {
					b.WriteString(" ")
					rest--
				}
			}
			b.WriteString(words[last-1])
		}
		output = append(output, b.String())
		index = last
	}

	return output
}
