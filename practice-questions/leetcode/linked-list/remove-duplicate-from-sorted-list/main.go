package main

import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil {
    return head
  }
	
  sentinel := &ListNode{Val: 0, Next: head}
  curr, prev := head, sentinel
	
  for curr != nil {
		fmt.Println("Begining Curr: ",linkedListToSlice(curr))
		fmt.Println("Begining Prev: ",linkedListToSlice(prev))
		fmt.Println("=======================================")

		fmt.Printf("prev.Next: %d, curr: %d\n", prev.Next.Val, curr.Val)
    if curr.Next != nil && curr.Val == curr.Next.Val {
      curr = curr.Next
      continue
    }

    if prev.Next != curr {
      prev.Next = curr.Next
    } else {
			fmt.Printf("Prev: %d, Prev.Nex: %d\n", prev.Val, prev.Next.Val)
			prev = prev.Next
			// fmt.Println("Else prev: ", linkedListToSlice(prev))
    }
		fmt.Println("End Curr: ",linkedListToSlice(curr))
		fmt.Println("End Prev: ",linkedListToSlice(prev))
		fmt.Println("=======================================")

    curr = curr.Next
  }

  return sentinel.Next
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func linkedListToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	nums := []int{1,2,3,3,4,4,5}
	head := sliceToLinkedList(nums)
	linkListResult := deleteDuplicates(head)
	result := linkedListToSlice(linkListResult)
	fmt.Println(result)
}