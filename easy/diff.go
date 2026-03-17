func findTheDifference(s string, t string) byte {
    var res byte

    for _, c := range s {
        res ^= byte(c)
    }

    for _, c := range t {
        res ^= byte(c)
    }

    return res
}
