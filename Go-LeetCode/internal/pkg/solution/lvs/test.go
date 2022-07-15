package main

import (
	"fmt"
	"sort"
	"strconv"
)

type HashFunc func(key []byte) uint32

type ConsistentHash struct {
	hash       HashFunc
	hashvals   []int
	hashToKey  map[int]string
	virtualNum int
}

func NewConsistentHash(virtualNum int, fn HashFunc) *ConsistentHash {
	return &ConsistentHash{
		hash:       fn,
		virtualNum: virtualNum,
		hashToKey:  make(map[int]string),
	}
}

func (ch *ConsistentHash) AddNode(keys ...string) {
	for _, k := range keys {
		for i := 0; i < ch.virtualNum; i++ {
			conv := strconv.Itoa(i)
			hashval := int(ch.hash([]byte(conv + k)))
			ch.hashvals = append(ch.hashvals, hashval)
			ch.hashToKey[hashval] = k
		}
	}
	sort.Ints(ch.hashvals)
}

func (ch *ConsistentHash) GetNode(key string) string {
	if len(ch.hashToKey) == 0 {
		return ""
	}
	keyhash := int(ch.hash([]byte(key)))
	id := sort.Search(len(ch.hashToKey), func(i int) bool {
		return ch.hashvals[i] >= keyhash
	})
	return ch.hashToKey[ch.hashvals[id%len(ch.hashvals)]]
}

func main() {
	ch := NewConsistentHash(3, func(key []byte) uint32 {
		ret, _ := strconv.Atoi(string(key))
		return uint32(ret)
	})
	ch.AddNode("1", "3", "5", "7")
	testkeys := []string{"12", "4", "7", "8"}
	for _, k := range testkeys {
		fmt.Printf("k:%s,node:%s\n", k, ch.GetNode(k))
	}
}
