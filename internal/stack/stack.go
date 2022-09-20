package stack

import (
	"algo/internal/linklist"
	"errors"
	"sync"
)

var (
	ErrEmpty = errors.New("stack has been empty")
)

type Stack struct {
	mutex   sync.RWMutex
	top     int
	topElem any
	length  int
	data    *linklist.LinkList[any]
}

// Init 初始化
func (s *Stack) Init() {
	s.top = -1
	s.topElem = nil
	s.length = 0
	s.data = &linklist.LinkList[any]{}
	s.data.Init()
}

// Length 栈长度
func (s *Stack) Length() int {
	return s.length
}

// Top 返回栈顶元素和栈顶索引
func (s *Stack) Top() (top int, elem any) {
	return s.top, s.topElem
}

// Pop 出栈
func (s *Stack) Pop() (elem any, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.top == -1 || s.length == 0 {
		return nil, ErrEmpty
	}
	elem = s.topElem
	s.top--
	s.length--
	err = s.data.TailDeduct()
	if err != nil {
		return nil, err
	}
	s.topElem, err = s.data.GetElem(s.length)
	if err != nil {
		if errors.Is(err, linklist.ErrInvalidInput) {
			s.topElem = nil
		} else {
			return nil, err
		}
	}
	return elem, nil
}

// Push 入栈
func (s *Stack) Push(elem any) (err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.top++
	s.length++
	s.topElem = elem
	err = s.data.TailAppend(elem)
	if err != nil {
		return err
	}
	return nil
}
