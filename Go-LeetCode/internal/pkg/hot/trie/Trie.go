package trie

type Trie struct {
	next  []*Trie
	isEnd bool
}

func Constructor() Trie {
	root := Trie{
		next:  make([]*Trie, 26),
		isEnd: false,
	}
	return root
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.next[idx] == nil {
			trie := Constructor()
			node.next[idx] = &trie
		}
		node = node.next[idx]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		idx := c - 'a'
		if node.next[idx] == nil {
			return false
		}
		node = node.next[idx]
	}
	return true
}
