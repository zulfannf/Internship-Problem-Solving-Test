package main

import "fmt"


func isValid(s string) bool {
        stack := []rune{}
        pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

        for _, char := range s {
                if _, ok := pairs[char]; ok { 
                        if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
                                return false
                        }
                        stack = stack[:len(stack)-1]
                } else {
                        stack = append(stack, char)
                }
        }

        return len(stack) == 0
}

func main() {
        s := "{[(())]}"
        fmt.Println(isValid(s)) 
}
