func countSegments(s string) int {
    cnt := 0

    for i,c := range s {
        if c != ' ' && (i == 0 || s[i-1] == ' '){
            cnt++
        }
    }
    return cnt
}
