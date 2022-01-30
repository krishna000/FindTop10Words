package main

const (
	ALPHABET_SIZE = 26
)

type Trie struct {
	head *TrieNode
}

type TrieNode struct {
	child [ALPHABET_SIZE]*TrieNode
	isEnd bool
	count int64
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.head = nil
	return trie
}

func (trie *Trie) AddWord(word string) {
	if trie.head == nil {
		trie.head = getTrieNode()
	}
	trie.head.addWord(word)
}

func getTrieNode() *TrieNode {
	node := new(TrieNode)
	node.isEnd = false
	node.count = 0
	return node
}

func (trieNode *TrieNode) addWord(word string) {
	length := len(word)
	headNode := trieNode
	for i := 0; i < length; i++ {
		if headNode.child[word[i]-'a'] == nil {
			headNode.child[word[i]-'a'] = getTrieNode()
		}
		if i == length-1 {
			headNode.child[word[i]-'a'].count += 1
			headNode.child[word[i]-'a'].isEnd = true
		}
		headNode = headNode.child[word[i]-'a']
	}
}
