package main

import (
	"fmt"
	"strings"
)

/*
   public static void main(String[] args){
       List<String> inputs = new ArrayList<>();
       inputs.add("def");
       inputs.add("abc:");
       inputs.add("  bcc");
       inputs.add("  abc:");
       inputs.add("    def");
       inputs.add("    def");
       inputs.add("  bcc");

       System.out.println(valid_python_indentation(inputs));
   }
*/
func main() {
	inputs := []string{"def", "abc:", "  bcc", "  abc:", "    def", "    def", "  bcc"}
	fmt.Println(isValid(inputs))
}

func isValid(inputs []string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(inputs); i++ {
		if len(stack) == 0 {
			if getLevel(inputs[i]) != 0 {
				return false
			}
		} else if stack[len(stack)-1][len(stack[len(stack)-1])-1] == ':' {
			if getLevel(stack[len(stack)-1]) >= getLevel(inputs[i]) {
				return false
			}
		} else {
			for len(stack) != 0 && getLevel(stack[len(stack)-1]) > getLevel(inputs[i]) {
				stack = stack[:len(stack)-1]
			}
			if len(stack) != 0 && getLevel(stack[len(stack)-1]) != getLevel(inputs[i]) {
				return false
			}
		}
		stack = append(stack, inputs[i])
	}
	return true
}

func getLevel(str string) int {
	return len(str) - len(strings.Trim(str, " "))
}
