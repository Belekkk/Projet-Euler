package main

import "fmt"

func inverse(val int) int {
    reversed := 0
    for val > 0 {
        reversed = 10*reversed + val%10
        val = val / 10
    }
    return reversed
}

func isPalindrome(val int) bool {
    if val == inverse(val) {
        return true
    }
    return false
}

func main() {
    max := 0
    a := 999
    for a >= 100 {
        b := 999
        for b >= a {
            if a*b <= max {
                break
            }
            if isPalindrome(a * b) {
                 max = a * b
            }
            b --
        }
        a --
    }
    fmt.Println(max)
}