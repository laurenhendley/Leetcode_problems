func isSubsequence(s string, t string) bool {
    if len(t) < len(s) {
        return false
    }

    sl := 0
    tl := 0

    for sl < len(s) && tl < len(t) {
        if s[sl] == t[tl] {
            sl++
        }
        tl++
    }

    return sl == len(s)
}
