func isPerfectSquare(num int) bool {
    if num < 2{
        return true
    }
    l := 0
    r := int(num / 2)

    for l <= r {
        mid := int((l+r) / 2)
        square := mid * mid

        if square == num {
            return true
        } else if square > num {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return false
}
