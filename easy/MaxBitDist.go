func binaryGap(n int) int {
    max_dist := 0
    curr_dist := 0
    found_first := false

    for n > 0 {
        curr_bit := n % 2

        if curr_bit == 1 {
            if found_first {
                max_dist = max(max_dist, curr_dist)
            }
            curr_dist = 1
            found_first = true
        } else {
            if found_first {
                curr_dist++
            }
        }
        n /= 2
    }
    
    return max_dist
}
