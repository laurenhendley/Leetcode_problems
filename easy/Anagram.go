func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    cnt := make([]int, 26)

    for i := 0; i < len(s); i++{
        cnt[s[i] - 'a']++
        cnt[t[i] - 'a']--
    }

    for _, j := range cnt {
        if j != 0 {
            return false
        }
    }

    return true
}
