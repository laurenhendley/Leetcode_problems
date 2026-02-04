func convertToTitle(columnNumber int) string {
    if columnNumber == 0 {
        return "A"
    }

    res := ""
    
    for columnNumber > 0 {
        columnNumber--
        remainder := columnNumber % 26
        res = string('A'+remainder)+res
        columnNumber/=26
    }

    return res
}
