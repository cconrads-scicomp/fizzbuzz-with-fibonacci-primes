package fibonacci


import (
	"math/big"
)



type generator struct {
	a *big.Int
	b *big.Int
}



func MakeGenerator() *generator {
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)

	return &generator{f0, f1}
}



// Execute computes the next number in the sequence of Fibonacci numbers F(n)
// starting with F(1).
func (gen *generator) Execute() *big.Int {
	a := gen.a
	b := gen.b

	a.Add(a, b)

	gen.a, gen.b = b, a

	return b
}
