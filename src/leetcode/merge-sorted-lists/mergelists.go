package main

/**
other's solution

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := ListNode{}
    mergedList := &dummy
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            mergedList.Next = list1
            mergedList = list1
            list1 = list1.Next
        } else {
            mergedList.Next = list2
            mergedList = list2
            list2 = list2.Next
        }
    }
    if list2 != nil {
        list1 = list2
    }
    mergedList.Next = list1
    return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    for len(lists) > 1 {
        newLists := make([]*ListNode, 0, (len(lists)+1)/2)
        i:=0
        for ; i < len(lists) - 1; i+=2{
            newLists = append(newLists, mergeTwoLists(lists[i], lists[i+1]))
        }
        if i == len(lists) - 1 {
            newLists = append(newLists, lists[i])
        }
        lists = newLists
    }
    return lists[0]
}
*/
import (
	"container/heap"
	"fmt"
	"math/rand"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
type HeapInterface []*ListNode

func (hi HeapInterface) Len() int {
	return len(hi)
}
func (hi HeapInterface) Less(i, j int) bool {
	if hi[i] == nil {
		return true
	}
	if hi[j] == nil {
		return false
	}
	return hi[i].Val < hi[j].Val
}
func (hi HeapInterface) Swap(i, j int) {
	hi[i], hi[j] = hi[j], hi[i]
}
func (h *HeapInterface) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *HeapInterface) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *HeapInterface) Peek() any {
	return (*h)[len(*h)-1]
}

func mergeKLists(lists []*ListNode) *ListNode {
	var myHeap HeapInterface = lists
	var head ListNode
	var current = &head
	heap.Init(&myHeap)

	for len(myHeap) > 0 {
		list := heap.Pop(&myHeap).(*ListNode)
		if list == nil {
			continue
		}
		current.Next = list
		current = current.Next
		heap.Push(&myHeap, list.Next)
	}
	return head.Next
}

func main() {
	var test = []*ListNode{buildList(30), buildList(0), buildList(12), buildList(1), buildList(5)}
	for _, t := range test {
		printList(t)
	}
	list := mergeKLists(test)
	printList(list)
}

func printList(lst *ListNode) {
	for ; lst != nil; lst = lst.Next {
		fmt.Print(lst.Val, ",")
	}
	fmt.Println(nil)
}

func buildList(i int) *ListNode {
	var head *ListNode

	for j := rand.Int() % 100; i > 0; i-- {
		node := ListNode{Val: j}
		node.Next = head
		head = &node
		j -= rand.Int() % 100
	}
	return head
}
