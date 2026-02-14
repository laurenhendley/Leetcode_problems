func containsNearbyDuplicate(nums []int, k int) bool {
    set := make(map[int]struct{})

    for i := 0; i < len(nums); i++{
        _, found := set[nums[i]]
        if found {
            return true
        }

        set[nums[i]] = struct{}{}

        if len(set) > k {
            delete(set, nums[i-k])
        }
    }

    return false
}
