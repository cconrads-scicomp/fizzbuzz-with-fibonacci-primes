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

package main

import (
	"github.com/cconrads-scicomp/fizzbuzz-with-fibonacci-primes/fibonacci"
	"testing"
)

func Example() {
	// f1, f2, ... =
	//    1,    1,    2,    3,    5,    8,   13,   21,   34,   55,   89,  144,
	//  233,  377,  610,  987, 1597, 2584, 4181, 6765,
	fizzbuzz(20)
	// Output:
	// 1
	// 1
	// BuzzFizz
	// BuzzFizz
	// BuzzFizz
	// 8
	// BuzzFizz
	// Buzz
	// 34
	// Fizz
	// BuzzFizz
	// Buzz
	// BuzzFizz
	// 377
	// Fizz
	// Buzz
	// BuzzFizz
	// 2584
	// 4181
	// FizzBuzz
}

const n0 = uint64(1000)
const n1 = uint64(2000)

func BenchmarkFibonacciPrimeTest(b *testing.B) {
	b.StopTimer()
	m := 50

	for i := 0; i < b.N; i++ {
		gen := fibonacci.MakeGenerator()
		for n := uint64(0); n < n0; n++ {
			gen.Execute()
		}

		b.StartTimer()
		for n, f := gen.Execute(); n <= n1; n, f = gen.Execute() {
			f.ProbablyPrime(m)
		}
		b.StopTimer()
	}
}

func BenchmarkFibonacciFastPrimeTest(b *testing.B) {
	b.StopTimer()
	n0 := uint64(1000)

	for i := 0; i < b.N; i++ {
		gen := fibonacci.MakeGenerator()
		for n := uint64(0); n < n0; n++ {
			gen.Execute()
		}

		b.StartTimer()
		for n, f := gen.Execute(); n <= n1; n, f = gen.Execute() {
			fibonacci.IsPrime(n, f)
		}
		b.StopTimer()
	}
}
