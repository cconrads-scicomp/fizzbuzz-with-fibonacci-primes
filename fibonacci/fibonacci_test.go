// Copyright 2016 Christoph Conrads (https://christoph-conrads.name)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package fibonacci

import (
	"math/big"
	"testing"
)

func testGeneratorInvariant(t *testing.T, gen *generator) {
	a := gen.a
	b := gen.b

	if a.Sign() == -1 {
		t.Error("expected a > 0, got", a.String())
	}
	if b.Sign() < 1 {
		t.Error("expected b > 0, got", b.String())
	}

	two := big.NewInt(2)
	if a.Cmp(b) != -1 && a.Cmp(two) >= 0 {
		t.Error("expected a < b, got", a.String(), ">=", b.String())
	}
}

func TestGenerator(t *testing.T) {
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

func TestCompute(t *testing.T) {
	fibs := [...]*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(5),
		big.NewInt(8),
	}

	l := uint64(len(fibs))
	for i := uint64(0); i < l; i++ {
		f := Compute(i)
		g := fibs[i]

		if f.Cmp(g) != 0 {
			t.Error("expected: f=g, got:", f.String(), "!=", g.String())
		}
	}
}

func TestIsPrime(t *testing.T) {
	// indices n of prime Fibonacci numbers F(n)
	// https://oeis.org/A001605
	primeIndices := [...]uint64{
		3, 4, 5, 7, 11, 13, 17, 23, 29, 43, 47, 83, 131, 137, 359,
	}

	isPrimeIndex := make(map[uint64]bool)
	for i := 0; i < len(primeIndices); i++ {
		isPrimeIndex[primeIndices[i]] = true
	}

	generator := MakeGenerator()
	for n, f := generator.Execute(); n <= 370; n, f = generator.Execute() {
		if IsPrime(n, f) && !isPrimeIndex[n] {
			t.Error("expected: non-primal F(n), got: n =", n, "F(n) =", f)
		}
		if !IsPrime(n, f) && isPrimeIndex[n] {
			t.Error("expected: Fibonacci prime F(n), got: n =", n, "F(n) =", f)
		}
	}
}
