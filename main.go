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



func fizzbuzz(m uint64) {
	fibGen := fibonacci.MakeGenerator()
	r := big.NewInt(-1)
	i3 := big.NewInt(3)
	i5 := big.NewInt(5)

	for n, f := fibGen.Execute(); n <= m; n, f = fibGen.Execute() {
		if fibonacci.IsPrime(n, f) {
			fmt.Printf("BuzzFizz\n")
			continue
		}

		r.Mod(f, i3)
		hasFactor3 := IsZero(r)

		r.Mod(f, i5)
		hasFactor5 := IsZero(r)

		switch {
		case hasFactor3 && hasFactor5:
			fmt.Printf("FizzBuzz\n")
		case hasFactor3:
			fmt.Printf("Buzz\n")
		case hasFactor5:
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
