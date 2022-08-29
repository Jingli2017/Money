package main

import (
	"fmt"
	"strings"
)
/*
You decide to create a substitution cipher. The cipher alphabet is based on a key shared amongst those of your friends who don't mind spoilers.

Suppose the key is:
"The quick onyx goblin, Grabbing his sword ==}-------- jumps over the 1st lazy dwarf!".

We use only the unique letters in this key to set the order of the characters in the substitution table.

T H E Q U I C K O N Y X G B L R A S W D J M P V Z F

(spaces added for readability)

We then align it with the regular alphabet:
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
T H E Q U I C K O N Y X G B L R A S W D J M P V Z F

Which gives us the substitution mapping: A becomes T, B becomes H, C becomes E, etc.

Write a function that takes a key and a string and encrypts the string with the key.

Example:
key = "The quick onyx goblin, Grabbing his sword ==}-------- jumps over the 1st lazy dwarf!"
encrypt("It was all a dream.", key) -> "Od ptw txx t qsutg."
encrypt("Would you kindly?", key) -> "Pljxq zlj yobqxz?"
 */
func main() {
	key := "The quick onyx goblin, Grabbing his sword ==}-------- jumps over the 1st lazy dwarf!"
	messages := []string{"It was all a dream.", "Would you kindly?"}
	ans := []string{"Od ptw txx t qsutg.", "Pljxq zlj yobqxz?"}
	for i, message := range messages {
		temp := encrypt(key, message)
		if temp != ans[i] {
			fmt.Println(temp, ans[i])
		}
	}
}

func encrypt(key string, message string) string{
	set := make(map[byte]bool)
	list := make([]byte, 0)
	for i := 0; i < len(key); i++ {
		if !isAlphabet(key[i]) {
			continue
		}
		char := toLowerCase(key[i])
		if _, ok := set[char]; !ok {
			set[char] = true
			list = append(list, char)
		}
	}

	fmt.Println(len(list))

	mapping := make(map[byte]byte)
	k := 0
	// mapping lower
	for c := 'a'; c <= 'z'; c++ {
		mapping[byte(c)] = list[k]
		k++
	}
	k = 0
	// mapping upper
	for c := 'A'; c <= 'Z'; c++ {
		mapping[byte(c)] = toUpperCase(list[k])
		k++
	}
	// encrypt
	var b strings.Builder
	for i := 0; i < len(message); i++ {
		if !isAlphabet(message[i]){
			b.WriteByte(message[i])
		}else {
			b.WriteByte(mapping[message[i]])
		}
	}

	return b.String()
}

func isAlphabet(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func toLowerCase(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char
	}

	char += 'a' - 'A'
	return char
}

func toUpperCase(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char
	}

	char -= 'a' - 'A'
	return char
}
