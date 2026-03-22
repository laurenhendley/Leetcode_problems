import "math"

func findDisappearedNumbers(nums []int) []int {
    for _, v := range nums {
        i := int(math.Abs(float64(v))) - 1
        if nums[i] > 0 {
            nums[i] = -nums[i]
        }
    }

    res := []int{}

    for i, v := range nums {
        if v > 0 {
            res = append(res, i+1)
        }
    }

    return res
}
