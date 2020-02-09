package list

import "testing"

func TestSingleLinkedInsertToHead(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestSingleLinkedInsertToTail(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestSingleLinkedFindByIndex(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(10))
}

func TestSingleLinkedFindMiddleNode(t *testing.T) {
	l := NewSingleLinkedList()
	t.Log(l.FindMiddleNode())

	l = NewSingleLinkedList()
	l.InsertToTail(1)
	t.Log(l.FindMiddleNode())

	l = NewSingleLinkedList()
	for i := 0; i < 2; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindMiddleNode())

	l = NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindMiddleNode())
}

func TestSingleLinkedDeleteNode(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.Delete(l.head.next))
	l.Print()

	t.Log(l.Delete(l.head.next.next))
	l.Print()
}

func TestSingleLinkedReverse(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	l.Reverse()
	l.Print()
}

func TestSortedSingleLinkedListMerge(t *testing.T) {
	a := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		a.InsertToTail(2*i + 1)
	}
	a.Print()

	b := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		b.InsertToTail(2 * i)
	}
	b.Print()

	SortedSingleLinkedListMerge(a, b, func(x, y interface{}) bool { return x.(int) < y.(int) }).Print()
}
