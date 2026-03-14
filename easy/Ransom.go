func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }

    var chars [26]int

    for _, c := range magazine {
        chars[c - 'a']++
    }

    for _, c := range ransomNote{
        index := c - 'a'
        if chars[index] <= 0 {
            return false
        }
        chars[index] -= 1
    }

    return true
}
