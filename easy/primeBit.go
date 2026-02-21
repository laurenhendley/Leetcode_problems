import "math/bits"

func countPrimeSetBits(left int, right int) int {
    cnt := 0

    for i := left; i <= right; i++ {
        allBits := bits.OnesCount(uint(i))
        if isPrime(allBits){
            cnt++
        }
    }

    return cnt
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }

    return true
}

func optionTwo(left int, right int) int {
    cnt := 0

    primeSet := map[int]bool{
        2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true,
    }

    for i := left; i <= right; i++ {
        if primeSet[bits.OnesCount(uint(i))]{
            cnt++
        }
    }

    return cnt
}
