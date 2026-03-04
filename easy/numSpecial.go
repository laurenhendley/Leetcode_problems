func numSpecial(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    col := make([]int, n)
    row := make([]int, m)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                row[i]++
                col[j]++
            }
        }
    }

    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
                res++
            }
        }
    }

    return res
}
