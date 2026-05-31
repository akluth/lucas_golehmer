package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func isPrime(n uint64) bool {
	return new(big.Int).SetUint64(n).ProbablyPrime(20)
}

func mersenneNumber(exponent uint64) *big.Int {
	mersenne := new(big.Int).Lsh(big.NewInt(1), uint(exponent))
	return mersenne.Sub(mersenne, big.NewInt(1))
}

func checkMersennePrime(exponent uint64) bool {
	if exponent == 2 {
		return true
	}

	mersenne := mersenneNumber(exponent)
	s := big.NewInt(4)

	for i := uint64(0); i < exponent-2; i++ {
		s.Mul(s, s)
		s.Sub(s, big.NewInt(2))
		s.Mod(s, mersenne)
	}

	return s.Sign() == 0
}

func isMersennePrimeExponent(exponent uint64) bool {
	if exponent == 2 {
		return true
	}

	if exponent < 2 || !isPrime(exponent) {
		return false
	}

	return checkMersennePrime(exponent)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s EXPONENT\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Checks whether 2^EXPONENT - 1 is a Mersenne prime.")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(2)
	}

	exponent, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		os.Exit(2)
	}

	fmt.Printf("2^%d - 1 is ", exponent)

	if isMersennePrimeExponent(exponent) {
		fmt.Println("a Mersenne prime")
	} else {
		fmt.Println("NOT a Mersenne prime")
	}
}
