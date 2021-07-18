// :date 2021/7/15

package SyncSortedList

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	unMarked = int32(0)
	isMarked = int32(1)
)

type IntList struct {
	head   *intNode
	length int64
}

type intNode struct {
	value   int
	next    *intNode
	nLocker sync.Mutex
	marked  int32
}

func NewInt() *IntList {
	return &IntList{head: NewNode(0)}
}

func NewNode(value int) *intNode {
	return &intNode{value: value}
}

func (n *intNode) getNextNode() *intNode {
	return (*intNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&n.next))))
}

func (n *intNode) getValue() int {
	return int(atomic.LoadInt64((*int64)(unsafe.Pointer(&n.value))))
}

func (n *intNode) setValue(value int) {
	atomic.StoreInt64((*int64)(unsafe.Pointer(&n.value)), int64(value))
}

func (n *intNode) setNext(next *intNode) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&n.next)), unsafe.Pointer(next))
}

func (n *intNode) getMark() int32 {
	return atomic.LoadInt32(&n.marked)
}

func (l *IntList) Insert(value int) bool {
	var a, b *intNode
	for {
		a = l.head
		b = a.getNextNode()
		for b != nil && b.getValue() < value {
			a = b
			b = a.getNextNode()
		}

		if b != nil && b.getValue() == value {
			return false
		}

		a.nLocker.Lock()
		// 1. why need recheck a.next == b in here?
		if b == a.next && a.marked == unMarked {
			break
		} else {
			a.nLocker.Unlock()
			continue
		}
	}

	newNode := NewNode(value)
	newNode.next = b
	a.setNext(newNode)
	atomic.AddInt64(&l.length, 1)
	a.nLocker.Unlock()

	return true
}

func (l *IntList) Delete(value int) bool {
	var a, b *intNode
	for {
		a = l.head
		b = a.getNextNode()
		for b != nil && b.getValue() < value {
			a = b
			b = b.getNextNode()
		}
		if b == nil || b.getValue() != value {
			return false
		}

		b.nLocker.Lock()
		if b.marked == isMarked {
			b.nLocker.Unlock()
			return false
		}

		a.nLocker.Lock()
		// 2. why need a.next == b?
		if a.next != b || a.marked == isMarked {
			b.nLocker.Unlock()
			a.nLocker.Unlock()
			continue
		}
		break
	}

	a.setNext(b.next)
	atomic.AddInt32(&b.marked, isMarked)
	atomic.AddInt64(&l.length, -1)
	a.nLocker.Unlock()
	b.nLocker.Unlock()

	return true
}

func (l *IntList) Len() int {
	return int(atomic.LoadInt64(&l.length))
}

func (l *IntList) Range(f func(value int) bool) {
	n := l.head.getNextNode()
	for n != nil {
		if !f(n.getValue()) {
			break
		}
		n = n.getNextNode()
	}
}

func (l *IntList) Contains(value int) bool {
	a := l.head.getNextNode()
	for a != nil {
		if a.getValue() == value {
			// need atomic get marked
			return a.getMark() == unMarked
		} else {
			a = a.getNextNode()
		}
	}
	return false
}
