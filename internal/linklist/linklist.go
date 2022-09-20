/*
	linklist 实现双向链表（带头节点和尾结点）
*/
package linklist

import (
	"errors"
	"reflect"
	"sync"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("elem not found")
	ErrEmpty        = errors.New("list has been empty")
)

// ListNode 链表节点
type ListNode struct {
	data  any
	prior *ListNode
	next  *ListNode
}

// LinkList 链表，双向链表，带头节点和尾结点
type LinkList[T any] struct {
	mutex  sync.RWMutex
	head   *ListNode
	tail   *ListNode
	length int
}

// Init 初始化，初始化header和tail, head <-> tail
func (l *LinkList[T]) Init() {
	l.head = &ListNode{}
	l.tail = &ListNode{
		prior: l.head,
	}
	l.head.next = l.tail
	l.length = 0
}

// Length 链表长度
func (l *LinkList[T]) Length() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.length
}

// HeadAppend 头插
func (l *LinkList[T]) HeadAppend(elem T) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var p = &ListNode{
		data:  elem,
		prior: l.head,
		next:  l.head.next,
	}
	l.head.next.prior = p
	l.head.next = p
	l.length++
	return nil
}

// TailAppend 尾插
func (l *LinkList[T]) TailAppend(elem T) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var p = &ListNode{
		data:  elem,
		prior: l.tail.prior,
		next:  l.tail,
	}
	l.tail.prior.next = p
	l.tail.prior = p
	l.length++
	return nil
}

// Insert 在指定位置插入元素
func (l *LinkList[T]) Insert(i int, elem T) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if i < 1 || i > l.length+1 { // 有length+1一个位置可插入，而可删除的元素只有length个
		return ErrInvalidInput
	}
	var p = l.head.next
	var j = 1
	for j < i {
		j++
		p = p.next
	}
	var q = &ListNode{
		data:  elem,
		prior: p.prior,
		next:  p,
	}
	p.prior.next = q
	p.prior = q
	l.length++
	return nil
}

// HeadDeduct 头删
func (l *LinkList[T]) HeadDeduct() (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.length == 0 {
		return ErrEmpty
	}
	var p = l.head.next
	l.head.next = p.next
	p.next.prior = l.head
	l.length--
	return nil
}

// TailDeduct 尾删
func (l *LinkList[T]) TailDeduct() (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.length == 0 {
		return ErrEmpty
	}
	var p = l.tail.prior
	l.tail.prior = p.prior
	p.prior.next = l.tail
	l.length--
	return nil
}

// Delete 删除指定位置元素
func (l *LinkList[T]) Delete(i int) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if i < 1 || i > l.length {
		return ErrInvalidInput
	}
	var p = l.head.next
	var j = 1
	for j < i {
		j++
		p = p.next
	}
	p.prior.next = p.next
	p.next.prior = p.prior
	l.length--
	return nil
}

// LocateElem 查询指定数据节点的位置
func (l *LinkList[T]) LocateElem(elem T) (i int, err error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	var p = l.head.next
	var j = 1
	for !reflect.DeepEqual(p.data, elem) && j < l.length {
		j++
		p = p.next
	}
	if p == nil {
		return -1, ErrNotFound
	}
	return j, nil
}

// GetElem 获取指定位置的数据域
func (l *LinkList[T]) GetElem(i int) (elem T, err error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if i < 1 || i > l.length {
		return elem, ErrInvalidInput
	}
	var p = l.head.next
	var j = 1
	for j < i {
		j++
		p = p.next
	}
	elem = p.data.(T)
	return elem, nil
}
