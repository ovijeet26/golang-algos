package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	keys  map[byte]*TrieNode
	isEnd bool
}

func Constructor() Trie {

	return Trie{root: &TrieNode{keys: make(map[byte]*TrieNode, 0), isEnd: false}}
}

func (this *Trie) Insert(word string) {

	wordLength := len(word)
	currentNode := this.root

	// Iterate over the length of the word
	for i := 0; i < wordLength; i++ {

		currentCharacter := word[i]

		// If the current character is present, simply traverse the node
		if nextNode, ok := currentNode.keys[currentCharacter]; ok {

			currentNode = nextNode
		} else {
			// Create a new node and add to the chain
			newNode := TrieNode{keys: make(map[byte]*TrieNode), isEnd: false}

			currentNode.keys[currentCharacter] = &newNode

			currentNode = &newNode
		}
	}

	currentNode.isEnd = true
}

func (this *Trie) Search(word string) bool {

	currentNode := this.root

	wordLength := len(word)

	for i := 0; i < wordLength; i++ {

		currentCharacter := word[i]

		if nextNode, ok := currentNode.keys[currentCharacter]; ok {

			currentNode = nextNode
		} else {
			return false
		}
	}

	// We will return true only if the last node is marked as the end node for a word.
	return currentNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {

	currentNode := this.root
	prefixLength := len(prefix)

	for i := 0; i < prefixLength; i++ {

		currentCharacter := prefix[i]

		if nextNode, ok := currentNode.keys[currentCharacter]; ok {

			currentNode = nextNode
		} else {

			return false
		}
	}

	return true
}
