import "strings"

func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)
    c2w := map[rune]string{}
    w2c := map[string]rune{}

    if len(pattern) != len(words) {
        return false
    }

    for i:=0; i < len(words); i++ {
        if char, ok := w2c[words[i]]; !ok {
            w2c[words[i]] = rune(pattern[i])
        } else if char != rune(pattern[i]){
            return false
        }

        if word, ok := c2w[rune(pattern[i])]; !ok {
            c2w[rune(pattern[i])] = words[i]
        } else if word != words[i]{
            return false
        }
    }

    return true
}
