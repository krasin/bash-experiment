package gobash

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

func (s *stack) IsEmpty() interface{} {
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
		panic("PopMany(%d): stack contains just %d elements", n, s.cnt)
	}
	return
}

func newStack() *stack {
	return &stack{make([]interface{}, initialCapacity), 0}
}

type int16Stack struct {
	inner *stack
}

func (s *int16Stack) IsEmpty() bool {
	return inner.IsEmpty()
}

func (s *int16Stack) Push(v int16) {
	inner.Push(v)
}

func (s *int16Stack) Pop() int16 {
	return inner.Pop().int16
}

func (s *int16Stack) Peek() int16 {
	return inner.Peek().int16
}

func (s *int16Stack) PopMany(n int) {
	return inner.PopMany(n)
}

func newInt16Stack() *int16Stack {
	return &int16Stack{newStack()}
}

type YYSTYPEStack struct {
	inner *stack
}

func (s *YYSTYPEStack) Push(v YYSTYPE) {
	inner.Push(v)
}

func (s *YYSTYPEStack) Pop() YYSTYPE {
	return inner.Pop().YYSTYPE
}

func (s *YYSTYPEStack) PopMany(n int) {
	return inner.PopMany(n)
}

func newYYSTYPEStack() *YYSTYPEStack {
	return &YYSTYPEStack{newStack()}
}
