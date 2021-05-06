package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

func check_mersenne(prime int64) bool {
	mersenne := big.NewInt(int64(math.Pow(2, float64(prime))) - 1)

	if !mersenne.ProbablyPrime(0) {
		return false
	}

	return true
}

func is_mersenne(prime int64) bool {
	if prime == 2 {
		return true
	}

	if !big.NewInt(prime).ProbablyPrime(0) {
		return false
	}

	return check_mersenne(prime)
}

func main() {
	number_to_check, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Print(number_to_check, " is ")

	if check_mersenne(int64(number_to_check)) {
		fmt.Println("a mersenne prime")
	} else {
		fmt.Println("NOT a mersenne prime")
	}
}
