package main

import (
	"fmt"

	"github.com/raufhm/cmlabs-backend-fulltime-test/fizzbuzz"
	"github.com/raufhm/cmlabs-backend-fulltime-test/palindrome"
)

func main() {
	fizzbuzz.Solution()
	result := palindrome.Solution("")
	fmt.Println(result)
}
