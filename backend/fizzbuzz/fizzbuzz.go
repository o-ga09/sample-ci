package fizzbuzz

import "strconv"

func  FizzBuzz(arg int) (string, error) {
	if arg == 0 {
		return strconv.Itoa(arg),nil
	}

	if arg % 15 == 0 {
		return "fizzbuzz", nil
	} else if arg % 3 == 0 {
		return "fizz", nil
	} else if arg % 5 == 0 {
		return "buzz", nil
	} else {
		return strconv.Itoa(arg), nil
	} 
}