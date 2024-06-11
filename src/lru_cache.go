package src

// DoubleListNode is a node in a doubly linked list.
type DoubleListNode struct {
	Key, Value int
	Prev, Next *DoubleListNode
}

// LRUCache is a data structure that supports get and put operations.
type LRUCache struct {
	capacity   int
	cache      map[int]*DoubleListNode
	head, tail *DoubleListNode
}

// Cache is an interface for a cache.
type Cache interface {
	Get(key int) int
	Put(key int, value int)
}

// NewLRUCache creates a new LRUCache with the given capacity.
func NewLRUCache(capacity int) LRUCache {
	head, tail := &DoubleListNode{}, &DoubleListNode{}
	head.Next, tail.Prev = tail, head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DoubleListNode),
		head:     head,
		tail:     tail,
	}
}

// addNode adds a node to the end of the list.
func (lru *LRUCache) addNode(node *DoubleListNode) {
	prev := lru.tail.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = lru.tail
	lru.tail.Prev = node
}

// removeNode removes a node from the list.
func (lru *LRUCache) removeNode(node *DoubleListNode) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

// Get returns the value of the key if the key exists, otherwise returns -1.
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.removeNode(node)
		lru.addNode(node)

		return node.Value
	}

	return -1
}

// Put sets or inserts the value if the key is not already present.
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.removeNode(node)
	} else if len(lru.cache) == lru.capacity {
		node := lru.head.Next
		lru.removeNode(node)
		delete(lru.cache, node.Key)
	}

	node := &DoubleListNode{Key: key, Value: value}
	lru.cache[key] = node
	lru.addNode(node)
}
