func majorityElement(nums []int) int {
    el := nums[0]
    counter := 0

    for i:=0; i < len(nums); i++{
        if counter == 0{
            el = nums[i]
            counter = 1
        } else if el == nums[i] {
            counter++
        } else {
            counter--
        }
    }

    return el
}
