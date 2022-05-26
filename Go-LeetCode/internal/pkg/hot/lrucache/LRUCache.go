package lrucache

// LRUCache LRU 缓存 https://leetcode.cn/problems/lru-cache/
/**
要点总结：
1. LRUCache结构体字段包括size，capacity，cache（Map），head和tail
2. 内部节点结构体，包含pre和next，k和v
3. 封装三个内部方法，moveToHead、addToHead，removeNode，removeTail
*/
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*node
	head, tail *node
}

// node 内部节点
type node struct {
	k, v      int
	pre, next *node
}

func initNode(key, value int) *node {
	return &node{
		k: key,
		v: value,
	}
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*node{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.pre = lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.v
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.k)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.v = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *node {
	node := this.tail.pre
	this.removeNode(node)
	return node
}
