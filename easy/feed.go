func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    cookie, children, cnt := 0, 0, 0

    for children < len(g) && cookie < len(s) {
        if s[cookie] >= g[children] {
            cnt++
            children++
        }
        cookie++
    }

    return cnt
}
