func firstUniqChar(s string) int {
    var frequency [26]int

    for _, c := range s {
        frequency[c-'a'] += 1
    }

    for i, c := range s {
        if frequency[c-'a'] == 1 {
            return i
        }
    }

    return -1
}
