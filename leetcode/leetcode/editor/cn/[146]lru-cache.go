package cn

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode // 双向链表
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0), // 添加头部有哑节点
		tail:     initDLinkedNode(0, 0), // 添加尾部哑节点
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok { // 不存在则创建新的node，并将其加入链表头部与hash table中
		node = initDLinkedNode(key, value)
		this.addNode(node)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity { // 如果超出容量限制，则移除链表尾部元素
			node = this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}
	} else { // 存在则更新，并将元素提前到链表头部
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addNode(node *DLinkedNode) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
