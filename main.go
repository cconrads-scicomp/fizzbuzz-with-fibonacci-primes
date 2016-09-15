package main

import (
	"fmt"
	"strconv"
	"os"
	"math/big"
)


func IsZero(i *big.Int) bool {
	return i.Sign() == 0
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


	if n > 0 {
		fmt.Printf("1\n")
	}
	if n <= 1 {
		os.Exit(0)
	}


	a := big.NewInt(0)
	b := big.NewInt(1)
	r := big.NewInt(-1)
	i3 := big.NewInt(3)
	i5 := big.NewInt(5)

	for i := uint64(2); i <= n; i++ {
		a.Add(a, b)
		a, b = b, a

		if b.ProbablyPrime(50) {
			fmt.Printf("BuzzFizz\n")
			continue
		}

		r.Mod(b, i3)
		has_factor3 := IsZero(r)

		r.Mod(b, i5)
		has_factor5 := IsZero(r)

		switch {
		case has_factor3 && has_factor5:
			fmt.Printf("FizzBuzz\n")
		case has_factor3:
			fmt.Printf("Buzz\n")
		case has_factor5:
			fmt.Printf("Fizz\n")
		default:
			fmt.Printf( "%s\n", b.String() )
		}
	}
}
