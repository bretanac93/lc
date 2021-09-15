package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func seedData() *LinkedList {
	data := []int{3, 4, 5, 6}
	ll := NewLinkedList()

	ll.FromArray(data)

	// 6 -> 5 -> 4 -> 3
	return ll
}

func TestInsert(t *testing.T) {
	ll := seedData()

	ll.Insert(11)
	ll.Insert(12)

	assert.Equal(t, 12, ll.head.Val, "the head should be 12")
	assert.Equal(t, 11, ll.head.Next.Val, "the following value should be 11")

	assert.ElementsMatch(t, []int{12, 11, 6, 5, 4, 3}, ll.ToArray())
}

func TestSearchFoundSomething(t *testing.T) {
	ll := seedData()

	found := ll.Search(6)

	assert.Equal(t, 6, found.Val, "found element should be labeled equally to the search criteria")
	assert.Equal(t, 5, found.Next.Val)

	found = ll.Search(4)

	assert.Equal(t, 5, found.Previous.Val)
	assert.Equal(t, 3, found.Next.Val)
}

func TestSearchFoundNothing(t *testing.T) {
	ll := seedData()

	notFound := ll.Search(49)

	assert.Nil(t, notFound)
}
