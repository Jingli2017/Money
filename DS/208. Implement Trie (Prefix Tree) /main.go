package main

import "fmt"

/*
   List<String> words = new ArrayList<>();
   words.add("ab");
   words.add("a");
   words.add("de");
   words.add("abde");
*/
func main() {
	root := &Trie{}
	root.Insert("ab")
	root.Insert("a")
	root.Insert("de")
	root.Insert("abde")

	fmt.Println(root.StartsWith("abb"))

	prefixNode := root.StartsWith("ab")
	fmt.Printf("%#v", prefixNode)

	output := make([]string, 0)
	dfs(&output, "ab", prefixNode)
	fmt.Println(output)
}

func dfs(output *[]string, curr string, prefixNode *Trie) {
	if prefixNode == nil {
		return
	}

	if prefixNode.isEnd {
		*output = append(*output, curr)
	}

	for i := 0; i < 26; i++ {
		dfs(output, curr + string(rune('a'+i)), prefixNode.children[i])
	}
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	if len(word) == 1 {
		if this.children[word[0]-'a'] != nil {
			this.children[word[0]-'a'].isEnd = true
		} else {
			newTrie := &Trie{}
			newTrie.isEnd = true
			this.children[word[0]-'a'] = newTrie
		}
	} else {
		if this.children[word[0]-'a'] != nil {
			this.children[word[0]-'a'].Insert(word[1:])
		} else {
			newTrie := &Trie{}
			this.children[word[0]-'a'] = newTrie
			newTrie.Insert(word[1:])
		}
	}
}

func (this *Trie) Search(word string) bool {
	if this.children[word[0]-'a'] == nil {
		return false
	}
	if len(word) == 1 {
		return this.children[word[0]-'a'].isEnd
	} else {
		return this.children[word[0]-'a'].Search(word[1:])
	}
}

func (this *Trie) StartsWith(word string) *Trie {
	if this.children[word[0]-'a'] == nil {
		return nil
	}
	if len(word) == 1 {
		return this.children[word[0]-'a']
	} else {
		return this.children[word[0]-'a'].StartsWith(word[1:])
	}
}
