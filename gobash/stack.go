package gobash

import (
	"fmt"
)

const (
	initialCapacity = 100
)

type stack struct {
	a   []interface{}
	cnt int
}

func (s *stack) Push(v interface{}) {
	if s.cnt >= len(s.a) {
		b := make([]interface{}, 2*len(s.a))
		for i, v := range s.a {
			b[i] = v
		}
		s.a = b
	}
	s.a[s.cnt] = v
	s.cnt++
}

func (s *stack) Pop() interface{} {
	if s.cnt > 0 {
		s.cnt--
		return s.a[s.cnt]
		s.a[s.cnt] = nil
	}
	panic("Stack is empty")
}

func (s *stack) Peek() interface{} {
	if s.cnt > 0 {
		return s.a[s.cnt-1]
	}
	panic("Stack is empty")
}

func (s *stack) PeekN(n int) interface{} {
	if n > 0 {
		panic(fmt.Sprintf("PeekN(n=%d) - n must be non-positive", n))
	}
	if s.cnt < -n+1 {
		panic(fmt.Sprintf("PeekN(n=%d) failed: stack has just %d elements", n, s.cnt))
	}
	return s.a[s.cnt + n - 1]
}

func (s *stack) IsEmpty() bool {
	return s.cnt == 0
}

func (s *stack) PopMany(n int) {
	if n < 0 {
		panic(fmt.Sprintf("n: %d, must be non-negative", n))
	}
	if s.cnt >= n {
		old := s.cnt
		s.cnt -= n
		for i := s.cnt; i < old; i++ {
			s.a[i] = nil
		}
	} else {
		panic(fmt.Sprintf("PopMany(%d): stack contains just %d elements", n, s.cnt))
	}
	return
}

func newStack() *stack {
	return &stack{make([]interface{}, initialCapacity), 0}
}

type intStack struct {
	inner *stack
}

func (s *intStack) IsEmpty() bool {
	return s.inner.IsEmpty()
}

func (s *intStack) Push(v int) {
	s.inner.Push(v)
}

func (s *intStack) Pop() int {
	return s.inner.Pop().(int)
}

func (s *intStack) Peek() int {
	return s.inner.Peek().(int)
}

func (s *intStack) PopMany(n int) {
	s.inner.PopMany(n)
}

func newIntStack() *intStack {
	return &intStack{newStack()}
}

type YYSTYPEStack struct {
	inner *stack
}

func (s *YYSTYPEStack) IsEmpty() bool {
	return s.inner.IsEmpty()
}

func (s *YYSTYPEStack) Push(v YYSTYPE) {
	s.inner.Push(v)
}

func (s *YYSTYPEStack) Pop() YYSTYPE {
	return s.inner.Pop().(YYSTYPE)
}

func (s *YYSTYPEStack) PopMany(n int) {
	s.inner.PopMany(n)
}

func (s *YYSTYPEStack) PeekN(n int) YYSTYPE {
	return s.inner.PeekN(n).(YYSTYPE)
}

func newYYSTYPEStack() *YYSTYPEStack {
	return &YYSTYPEStack{newStack()}
}
