/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    mid_point := head
    end_point := head

    for end_point != nil && end_point.Next != nil {
        end_point = end_point.Next.Next
        mid_point = mid_point.Next
    }

    var prev *ListNode
    curr := mid_point
    for curr != nil{
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    for prev != nil {
        if prev.Val != head.Val{
            return false
        }
        prev, head = prev.Next, head.Next
    }

    return true
}
