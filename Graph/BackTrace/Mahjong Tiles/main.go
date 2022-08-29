package main

import "fmt"

/*
You've decided to make an advanced version of a variant of the game Mahjong, incorporating runs.

Players have a number of tiles, each marked 0-9. The tiles can be grouped into pairs or triples of the same tile or runs.

A run is a series of three consecutive tiles, like 123, 678, or 456. 0 counts as the lowest tile, so 012 is a valid run, but 890 is not.

A complete hand consists of exactly one pair, and any number (including zero) of triples and/or three-tile runs. Each tile is used in exactly one triple, pair, or run.

Write a function that returns true or false depending on whether or not the collection represents a complete hand under these rules.

hand1 = "11123" # True. 11 123
[0, 3, 1, 1 .... ]
hand2 = "12131" # True. Also 11 123. Tiles are not necessarily sorted.
hand3 = "11123455" # True. 111 234 55
[3, 1, 1, 1, 2]
hand4 = "11122334" # True. 11 123 234
hand5 = "11234" # True. 11 234
hand6 = "123456" # False. Needs a pair
hand7 = "11133355577" # True. 111 333 555 77
hand8 = "11223344556677" # True. 11 234 234 567 567 among others

[0,2,2,2,2,2,2,2,0,0]
hand9 = "12233444556677" # True. 123 234 44 567 567
hand10 = "11234567899" # False.
hand11 = "00123457" # False.
hand12 = "0012345" # False. A run is only three tiles
hand13 = "11890" # False. 890 is not a valid run
hand14 = "99" # True.
hand15 = "111223344" # False.

All Test Cases
advanced(hand1) => True
advanced(hand2) => True
advanced(hand3) => True
advanced(hand4) => True
advanced(hand5) => True
advanced(hand6) => False
advanced(hand7) => True
advanced(hand8) => True
advanced(hand9) => True
advanced(hand10) => False
advanced(hand11) => False
advanced(hand12) => False
advanced(hand13) => False
advanced(hand14) => True
advanced(hand15) => False
*/
func main() {
	tests := []string{
		"11123", "12131", "11123455", "11122334", "11234", "123456", "11133355577", "11223344556677", "12233444556677",
		"11234567899", "00123457", "0012345", "11890", "99", "111223344",
	}

	results := []bool{
		true,true,true,true,true,false,true,true,true,false,false,false,false,true,false,
	}

	for i, test := range tests {
		if canComplete(test) != results[i]{
			fmt.Println(i, test, results[i])
			return
		}
	}

	fmt.Println("all right")
}

func canComplete(s string) bool {
	// build counter map
	counter := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	// for 0-9 select pair
	for c := '0'; c <= '9'; c++ {
		if counter[byte(c)] < 2 { continue }
		counter[byte(c)] -= 2 // withdraw a pair
		if isTripleCompletable(copyMap(counter)) {
			return true
		}
		counter[byte(c)] += 2 // back
	}
	return false
}

func isTripleCompletable(counter map[byte]int) bool{
	// remove all triple
	for c := '0'; c <= '9'; c++ {
		counter[byte(c)] %= 3
	}
	// remove increase sequence until 7
	for c := '0'; c <= '7'; c++ {
		if counter[byte(c)] < 0 {
			return false
		}
		if counter[byte(c)] == 0 {
			continue
		}
		second := c + 1
		third := c + 2
		counter[byte(second)] -= counter[byte(c)]
		counter[byte(third)] -= counter[byte(c)]
	}
	return counter['8'] == 0 && counter['9'] == 0
}

func copyMap(counter map[byte]int) map[byte]int {
	duplicate := make(map[byte]int)
	for char, count := range counter {
		duplicate[char] = count
	}
	return duplicate
}

