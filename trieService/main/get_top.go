package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func getTop10MostUsedWordsFromContent(content string) *MinHeap {
	trie := NewTrie()
	addContentToTrie(content, trie)

	h := &MinHeap{}

	trie.traverseAndAddToHeap(h, 10)
	return h
}

func addContentToTrie(content string, trie *Trie) {
	wordSize := len(content)
	for i := 0; i < wordSize; {
		word := ""
		for i < len(content) {
			ch := content[i]
			if (ch <= 'Z' && ch >= 'A') || (ch <= 'z' && ch >= 'a') {
				word = fmt.Sprintf("%s%c", word, ch)
			} else {
				break
			}
			i++
		}
		if len(word) != 0 {
			trie.AddWord(strings.ToLower(word))
		}
		i++
	}
}

func (trie *Trie) traverseAndAddToHeap(h *MinHeap, cap int) {
	if trie.head == nil {
		fmt.Printf("Empty trie")
	} else {
		trie.head.traverseAndAddToHeap("", h, cap)
	}
}

func (node *TrieNode) traverseAndAddToHeap(s string, h *MinHeap, cap int) {
	if node.isEnd == true {
		addToHeap(h, cap, &WordCount{Word: s, Count: node.count})
	}
	for i := 0; i < len(node.child); i++ {
		if node.child[i] != nil {
			node.child[i].traverseAndAddToHeap(fmt.Sprintf("%s%c", s, 'a'+i), h, cap)
		}
	}
}

func addToHeap(h *MinHeap, capacity int, node *WordCount) {
	size := h.Len()
	if size < capacity {
		heap.Push(h, node)
	} else {
		top := heap.Pop(h)
		if top.(*WordCount).Count < node.Count {
			heap.Push(h, node)
		} else {
			heap.Push(h, top)
		}
	}
}

type WordCount struct {
	Word  string
	Count int64
}

type MinHeap []*WordCount

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*WordCount))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
