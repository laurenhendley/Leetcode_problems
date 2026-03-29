func findComplement(num int) int {
    binaryStr := strconv.FormatInt(int64(num), 2)
    length := len(binaryStr)

    mask := (1 << length) - 1

    return num ^ mask
}
