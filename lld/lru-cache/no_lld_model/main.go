package main

import "fmt"

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	cache    map[int]*Node
	capacity int
	recentD  *Node
	leastD   *Node
}

func Constructor(capacity int) LRUCache {
	recentD := &Node{Val: 0}
	leastD := &Node{Val: 0}

	recentD.Next, leastD.Prev = leastD, recentD

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		recentD:  recentD,
		leastD:   leastD,
	}
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.remove(node)
	l.insert(node)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; ok {
		node := l.cache[key]
		node.Val = value
		l.remove(node)
		l.insert(node)
	} else {
		if len(l.cache) == l.capacity {
			delete(l.cache, l.leastD.Prev.Key)
			l.remove(l.leastD.Prev)
		}

		newNode := &Node{
			Key: key,
			Val: value,
		}
		l.cache[key] = newNode
		l.insert(newNode)
	}
}

func (l *LRUCache) insert(node *Node) {
	node.Next = l.recentD.Next
	node.Prev = l.recentD
	l.recentD.Next.Prev = node
	l.recentD.Next = node
}

func (l *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func main() {
	obj := Constructor(2)
	// param_1 := obj.Get(key)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	// obj.Get(2)
	// obj.Put(1, 1)

	fmt.Println(*obj.recentD.Next)
	fmt.Println(*obj.leastD.Prev)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
