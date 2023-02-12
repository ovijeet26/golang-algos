package matrix

import (
	"fmt"
	"strings"
)

// Leet code link -> https://leetcode.com/problems/word-search-ii/
func findWords(board [][]byte, words []string) []string {

	validWordsMap := map[string]bool{}

	rowLength := len(board)
	colLength := len(board[0])

	prefixTrie := Constructor()

	for _, word := range words {

		prefixTrie.Insert(word)
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// Make a visited array
	visited := make([][]bool, 0)
	for i := 0; i < rowLength; i++ {

		visitedRow := make([]bool, colLength)
		visited = append(visited, visitedRow)
	}

	validWords := make([]string, 0, len(validWordsMap))

	// Call wordsearch for all the characters in the board.
	// For characters which don't match, we will return in O(1) time.
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {

			wordSearch2Dfs(i, j, board, visited, directions, validWordsMap, prefixTrie.root, []string{})
		}
	}

	for key := range validWordsMap {
		validWords = append(validWords, key)
	}

	return validWords
}

func wordSearch2Dfs(row int, col int, board [][]byte, visited [][]bool, directions [][]int, validWordsMap map[string]bool, trieNode *TrieNode, currentWord []string) {

	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || visited[row][col] {

		return
	}

	currentCharacter := board[row][col]

	nextTrieNode, isPresent := trieNode.keys[currentCharacter]

	if !isPresent {
		return
	}

	visited[row][col] = true
	currentWord = append(currentWord, string(currentCharacter))

	// Check if one word is found
	if nextTrieNode.isEnd == true {
		word := strings.Join(currentWord, "")
		validWordsMap[word] = true
	}

	for _, direction := range directions {

		rowOffset := direction[0]
		colOffset := direction[1]

		wordSearch2Dfs(row+rowOffset, col+colOffset, board, visited, directions, validWordsMap, nextTrieNode, currentWord)
	}

	// Enable backtracking
	visited[row][col] = false
}

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

func WordsSearch2Caller() {

	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	fmt.Println(findWords(board, words))
}
