package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func InitDLinkedNode(k, v int) *DLinkedNode {
	node := &DLinkedNode{
		key:   k,
		value: v,
		pre:   nil,
		next:  nil,
	}
	return node
}

// 初始化 LRUCache
func Constructor(cap int) *LRUCache {
	l := &LRUCache{
		size:     0,
		capacity: cap,
		cache:    make(map[int]*DLinkedNode),
		head:     InitDLinkedNode(0, 0),
		tail:     InitDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	// 移动链表
	lru.UpdateToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key, value int) {
	// 如果不存在，则插入新的值，否则将对应的值移动到链头
	node, ok := lru.cache[key]
	if !ok {
		newNode := InitDLinkedNode(key, value)
		if lru.size >= lru.capacity {
			// 删掉最少使用的那个
			lru.DeleteLeast()
		}
		lru.cache[key] = newNode
		lru.InsertNewHead(newNode)
	} else {
		node.value = value
		lru.UpdateToHead(node)
	}
}

func (lru *LRUCache) InsertNewHead(node *DLinkedNode) {
	tmp := lru.head.next
	lru.head.next = node
	node.pre = lru.head
	node.next = tmp
	tmp.pre = node
	lru.size++
}

func (lru *LRUCache) DeleteLeast() {
	node := lru.tail.pre
	lru.tail.pre = node.pre
	node.pre.next = node.next
	node.pre = nil
	node.next = nil
	lru.size--
	delete(lru.cache, node.key)
}

func (lru *LRUCache) UpdateToHead(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	tmp := lru.head.next
	lru.head.next = node
	node.pre = lru.head
	node.next = tmp
	tmp.pre = node
}
