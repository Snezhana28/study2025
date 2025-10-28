package containers

type Stack struct {
	Values []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.Values = append(s.Values, value)
}
func (s *Stack) Pop() (int, bool) {
	if len(s.Values) == 0 {
		return 0, false
	}
	top := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *Stack) Size() int {
	return len(s.Values)
}

func (s *Stack) Clear() {
	s.Values = []int{}
}

type Deque struct {
	Values []int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PushFront(value int) {
	d.Values = append([]int{value}, d.Values...)
}
func (d *Deque) PushBack(value int) {
	d.Values = append(d.Values, value)
}
func (d *Deque) PopFront() (int, bool) {
	if len(d.Values) == 0 {
		return 0, false
	}
	frontElement := d.Values[0]
	d.Values = d.Values[1:]
	return frontElement, true
}

func (d *Deque) PopBack() (int, bool) {
	if len(d.Values) == 0 {
		return 0, false
	}
	rearElement := d.Values[len(d.Values)-1]
	d.Values = d.Values[:len(d.Values)-1]
	return rearElement, true
}

func (d *Deque) IsEmpty() bool {
	return len(d.Values) == 0
}
func (d *Deque) Size() int {
	return len(d.Values)
}

func (d *Deque) Clear() {
	d.Values = []int{}
}
