import "strconv"

func summaryRanges(nums []int) []string {
    res := make([]string, 0)
    n := len(nums)

    for i := 0; i < n; i++ {
        start := nums[i]

        for (i + 1 < len(nums) && nums[i+1] == nums[i] + 1) {
            i++
        }

        if nums[i] == start {
            res = append(res, strconv.Itoa(start))
        } else {
            res = append(res, strconv.Itoa(start) + "->" + strconv.Itoa(nums[i]))
        }
    }
    return res
}
