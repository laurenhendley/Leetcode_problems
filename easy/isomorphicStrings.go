func isIsomorphic(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    sChars := []rune(s)
    tChars := []rune(t)

    if len(sChars) != len(tChars) {
        return false
    }

    sT := make(map[rune]rune)
    tS := make(map[rune]rune)

    for i:=0;i<len(sChars);i++ {
        charS := sChars[i]
        charT := tChars[i]

        if v, ok := sT[charS]; ok {
            if v != charT {
                return false
            }
        } else {
            sT[charS] = charT
        }

        if v, ok := tS[charT]; ok {
            if v != charS {
                return false
            }
        } else {
            tS[charT] = charS
        }
    }
    return true
}
