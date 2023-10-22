package main

import (
	"fmt"
	"os"

	"github.com/o-ga09/samplemonorepo/fizzbuzz"
)

func main() {
	for i := 0; i <= 100; i++ {
		s, err := fizzbuzz.FizzBuzz(i)
		if err != nil {
			fmt.Printf("err")
			os.Exit(1)
		}
		fmt.Printf("%d : %s\n", i, s)
	}
}