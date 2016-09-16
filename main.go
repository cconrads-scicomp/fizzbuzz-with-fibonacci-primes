package main

import (
	"fmt"
	"strconv"
	"os"
	"math/big"
	"fizzbuzz_prime/fibonacci"
)


func IsZero(i *big.Int) bool {
	return i.Sign() == 0
}



func fizzbuzz(n uint64) {
	fibGen := fibonacci.MakeGenerator()
	r := big.NewInt(-1)
	i3 := big.NewInt(3)
	i5 := big.NewInt(5)

	for i := uint64(1); i <= n; i++ {
		f := fibGen.Execute()

		if f.ProbablyPrime(50) {
			fmt.Printf("BuzzFizz\n")
			continue
		}

		r.Mod(f, i3)
		has_factor3 := IsZero(r)

		r.Mod(f, i5)
		has_factor5 := IsZero(r)

		switch {
		case has_factor3 && has_factor5:
			fmt.Printf("FizzBuzz\n")
		case has_factor3:
			fmt.Printf("Buzz\n")
		case has_factor5:
			fmt.Printf("Fizz\n")
		default:
			fmt.Printf( "%s\n", f.String() )
		}
	}
}



func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: fizzbuzz_prime <n>\n")
		os.Exit(1)
	}

	input, err := strconv.Atoi(args[1])
	if err != nil || input < 0 {
		fmt.Fprintf(os.Stderr, "n must be non-negative\n")
		os.Exit(2)
	}

	n := uint64(input)
	fizzbuzz(n)

}
