package hash

const ArraySize = 100

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(key string) {
	if b.search(key) {
		// Already exist: skip duplicates
		return
	}

	newNode := &bucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	current := b.head

	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}

	return false
}

func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

type bucketNode struct {
	key  interface{}
	next *bucketNode
}

type HashTable struct {
	data [ArraySize]*bucket
}

func NewDataStorage(size ...int) *HashTable {
	newHT := &HashTable{}

	for i := range newHT.data {
		newHT.data[i] = &bucket{}
	}

	return newHT
}

func hash(value string) int {
	sum := 0

	for _, c := range value {
		sum += int(c)
	}

	return sum % ArraySize
}

func (d *HashTable) Insert(value string) {
	index := hash(value)
	d.data[index].insert(value)
}

func (d *HashTable) Delete(value string) {
	index := hash(value)
	d.data[index].delete(value)
}

func (d *HashTable) Search(value string) bool {
	index := hash(value)
	return d.data[index].search(value)
}
