func toHex(num int) string {
    if num == 0{
        return "0"
    }

    res := ""
    n := num & 0xffffffff

    for n > 0 {
        rem := n & 15
        if rem < 10 {
            res = string('0'+rem) + res
        } else {
            res = string('a'+(rem-10)) + res
        }
        n = n >> 4
    }
    
    return res
}
