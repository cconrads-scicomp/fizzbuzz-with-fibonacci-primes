package fibonacci


import (
	"math/big"
)



type generator struct {
	n uint64
	a *big.Int
	b *big.Int
}



func MakeGenerator() *generator {
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)

	return &generator{1, f0, f1}
}



// Execute computes the next number in the sequence of Fibonacci numbers F(n)
// starting with F(1).
func (gen *generator) Execute() (uint64, *big.Int) {
	n := gen.n
	a := gen.a
	b := gen.b

	a.Add(a, b)

	gen.n = n+1
	gen.a, gen.b = b, a

	return n, b
}



func Compute(n uint64) *big.Int {
	switch {
	case n == 0: return big.NewInt(0)
	case n == 1 || n == 2: return big.NewInt(1)
	default:
		a := Compute(n-1)
		b := Compute(n-2)
		return b.Add(a, b)
	}
}



func IsPrime(n uint64, f *big.Int) bool {
	if n == 4 {
		return true
	}

	b := big.NewInt(0)
	b.SetUint64(n)

	return b.ProbablyPrime(1) && f.ProbablyPrime(64)
}
