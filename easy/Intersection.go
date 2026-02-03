/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    trackerA := headA
    trackerB := headB

    for trackerA != trackerB {
       if trackerA != nil{
        trackerA = trackerA.Next
       }else{
        trackerA = headB
       }

       if trackerB != nil{
        trackerB = trackerB.Next
       }else{
        trackerB=headA
       }

    }
    return trackerB
}
