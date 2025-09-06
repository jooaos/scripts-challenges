package main

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head, tail}
}

func (dll *DoublyLinkedList) AddToFront(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
}

func (dll *DoublyLinkedList) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *DoublyLinkedList) MoveToFront(node *Node) {
	dll.RemoveNode(node)
	dll.AddToFront(node)
}

func (dll *DoublyLinkedList) RemoveTail() *Node {
	if dll.tail.prev == dll.head {
		return nil
	}
	lru := dll.tail.prev
	dll.RemoveNode(lru)
	return lru
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	list     *DoublyLinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		list:     NewDoublyLinkedList(),
	}
}

func (c *LRUCache) Get(key int) int {
	if node, exists := c.cache[key]; exists {
		c.list.MoveToFront(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		node.value = value
		c.list.MoveToFront(node)
	} else {
		if len(c.cache) >= c.capacity {
			lru := c.list.RemoveTail()
			if lru != nil {
				delete(c.cache, lru.key)
			}
		}
		newNode := &Node{key: key, value: value}
		c.list.AddToFront(newNode)
		c.cache[key] = newNode
	}
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // evicts key 2
	fmt.Println(lru.Get(2)) // -1 (not found)
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}
