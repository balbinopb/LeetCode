
package main

import"fmt"

func lengthOfLastWord(s string) int {
    length := 0
    inWord := false

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            length++
            inWord = true
        } else if inWord {
            break
        }
    }
    
    return length
}
func main(){
	fmt.Println(lengthOfLastWord("helloo world!"))
}