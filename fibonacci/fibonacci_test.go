package fibonacci


import (
	"testing"
	"math/big"
)


func testGeneratorInvariant(t *testing.T, gen *generator) {
	a := gen.a
	b := gen.b

	if a.Sign() == -1 {
		t.Error( "expected a > 0, got", a.String() )
	}
	if b.Sign() < 1 {
		t.Error( "expected b > 0, got", b.String() )
	}

	two := big.NewInt(2)
	if a.Cmp(b) != -1 && a.Cmp(two) >= 0 {
		t.Error( "expected a < b, got", a.String(), ">=", b.String() )
	}
}



func Test1(t *testing.T) {
	generator := MakeGenerator()
	testGeneratorInvariant(t, generator)

	var f func(int) *big.Int
	f = func (n int) *big.Int {
		switch {
		case n == 0: return big.NewInt(0)
		case n == 1 || n == 2: return big.NewInt(1)
		default:
			a := f(n-1)
			b := f(n-2)
			return b.Add(a, b)
		}
	}


	for i := 1; i < 20; i++ {
		x := generator.Execute()
		y := f(i)

		testGeneratorInvariant(t, generator)
		if x.Cmp(y) != 0 {
			t.Error("generator/f(n) mismatch", x.String(), y.String())
		}
	}
}
