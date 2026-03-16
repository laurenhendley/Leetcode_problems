import "sort"

func thirdMax(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    cnt := 1
    for i := n - 2; i >= 0; i-- {
        if nums[i] != nums[i+1] {
            cnt++
            if cnt == 3{
                return nums[i]
            }
        }
    } 

    return nums[n-1]
}
