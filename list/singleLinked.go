package list

import (
	"errors"
	"fmt"
)

type SingleLinkedListNode struct {
	next  *SingleLinkedListNode
	value interface{}
}

func NewSingleLinkedListNode(v interface{}) *SingleLinkedListNode {
	s := new(SingleLinkedListNode)
	s.value = v
	return s
}

func (s *SingleLinkedListNode) Next() *SingleLinkedListNode { return s.next }
func (s *SingleLinkedListNode) Value() interface{}          { return s.value }

type SingleLinkedList struct {
	head   *SingleLinkedListNode
	length uint
}

func NewSingleLinkedList() *SingleLinkedList {
	s := new(SingleLinkedList)
	// sentinel
	s.head = NewSingleLinkedListNode(0)
	return s
}

func (s *SingleLinkedList) Head() *SingleLinkedListNode { return s.head.next }
func (s *SingleLinkedList) Len() uint                   { return s.length }

func (s *SingleLinkedList) Print() {
	cur := s.Head()
	var out string
	for cur != nil {
		out += fmt.Sprintf("%+v", cur.Value())
		cur = cur.next
		if cur != nil {
			out += "->"
		}
	}
	fmt.Println(out)
}

func (s *SingleLinkedList) InsertAfter(p *SingleLinkedListNode, v interface{}) error {
	if p == nil {
		return errors.New("invalid position")
	}

	n := NewSingleLinkedListNode(v)
	oldNext := p.next
	p.next = n
	n.next = oldNext

	s.length++

	return nil
}

func (s *SingleLinkedList) InsertBefore(p *SingleLinkedListNode, v interface{}) error {
	if p == nil || p == s.head {
		return errors.New("invalid position")
	}

	cur := s.Head()
	pre := s.head
	for cur.next != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return errors.New("position not found")
	}

	n := NewSingleLinkedListNode(v)
	pre.next = n
	n.next = cur

	s.length++

	return nil
}

func (s *SingleLinkedList) InsertToHead(v interface{}) error {
	return s.InsertAfter(s.head, v)
}

func (s *SingleLinkedList) InsertToTail(v interface{}) error {
	cur := s.head
	for cur.next != nil {
		cur = cur.next
	}
	return s.InsertAfter(cur, v)
}

func (s *SingleLinkedList) FindByIndex(i uint) (*SingleLinkedListNode, error) {
	if i >= s.length {
		return nil, errors.New("index out of range")
	}

	cur := s.Head()
	for c := uint(0); c < i; c++ {
		cur = cur.next
	}

	return cur, nil
}

func (s *SingleLinkedList) FindMiddleNode() *SingleLinkedListNode {
	if s.head == nil || s.Head() == nil {
		return nil
	}

	if s.Head().next == nil {
		return s.Head()
	}

	slow, fast := s.head, s.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func (s *SingleLinkedList) Delete(p *SingleLinkedListNode) error {
	if p == nil || p == s.head {
		return errors.New("invalid position")
	}

	cur := s.Head()
	pre := s.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return errors.New("position not found")
	}

	pre.next = p.next

	s.length--

	return nil
}

func (s *SingleLinkedList) Reverse() {
	if s.head == nil || s.Head() == nil || s.Head().next == nil {
		return
	}

	var newHead *SingleLinkedListNode
	cur := s.Head()
	for cur != nil {
		tmp := cur.next
		cur.next = newHead
		newHead = cur
		cur = tmp
	}

	s.head.next = newHead
}

func SortedSingleLinkedListMerge(a, b *SingleLinkedList, less func(x, y interface{}) bool) *SingleLinkedList {
	if a == nil || a.head == nil || a.Head() == nil {
		return b
	}
	if b == nil || b.head == nil || b.Head() == nil {
		return a
	}

	r := NewSingleLinkedList()
	cur := r.head
	cura := a.Head()
	curb := b.Head()

	for cura != nil && curb != nil {
		if less(cura.Value(), curb.Value()) {
			cur.next = cura
			cura = cura.next
		} else {
			cur.next = curb
			curb = curb.next
		}
		cur = cur.next
	}

	if cura != nil {
		cur.next = cura
	} else {
		cur.next = curb
	}

	return r
}
