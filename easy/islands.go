func islandPerimeter(grid [][]int) int {
    r := len(grid)
    c := len(grid[0])
    p := 0

    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if(grid[i][j] == 1){
                p += 4
                if i > 0 && grid[i-1][j] == 1 {
                    p -= 2
                }
                if j > 0 && grid[i][j-1] == 1 {
                    p -= 2
                }
            }
        }
    }
    return p
}
