import (
    "strings"
    "unicode"
)

func isPalindrome(s string) bool {
    s = removeButLetter(s)
    s = strings.ToLower(s)
    return s == reverse(s)
}

func reverse(s string) string {
    runes := []rune(s)
    for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
        runes[i],runes[j]=runes[j],runes[i]
    }
    return string(runes)
}

func removeButLetter(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r){
            return r
        }
        return -1
    }, s)
}
