func findMaxConsecutiveOnes(nums []int) int {
    maxCnt := 0
    cur := 0

    for _, v := range nums {
        if v == 1 {
            cur++
            if cur > maxCnt{
                maxCnt = cur
            }
        } else{
            cur = 0
        }
    }

    return maxCnt
}
