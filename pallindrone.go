package main
import "fmt"


func isPalindrome(x int) bool {
    // Handle negative numbers and numbers ending with 0 (except 0 itself)
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }

    // Reverse the second half of the integer
    reversed := 0
    for x > reversed {
        // Extract the last digit of the integer
        digit := x % 10
        // Add the digit to the reversed integer
        reversed = reversed*10 + digit
        // Remove the last digit from the original integer
        x /= 10
    }

    // Check if the original integer equals the reversed integer or the reversed integer divided by 10 (for odd-digit integers)
    return x == reversed || x == reversed/10
}
func main (){
	fmt.Println(isPalindrome(4657))
}