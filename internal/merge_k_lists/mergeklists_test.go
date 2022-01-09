package merge_k_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	expected := []int{11, 21, 23, 25, 41, 42, 51, 56, 66, 72}
	a := createLinkedList([]int{11, 41, 51})
	b := createLinkedList([]int{21, 23, 42})
	c := createLinkedList([]int{25, 56, 66, 72})
	list1 := []*LinkedListNode{a, b, c}

	res := toList(mergeKLists(list1))

	assert.Equal(t, expected, res)
}
