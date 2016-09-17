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

	for n, f := generator.Execute(); n <= 20; n, f = generator.Execute() {
		g := Compute(n)

		testGeneratorInvariant(t, generator)
		if f.Cmp(g) != 0 {
			t.Error("generator/f(n) mismatch", f.String(), g.String())
		}
	}
}
