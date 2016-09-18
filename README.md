# FizzBuzz with Fibonacci Primes

This program is a [FizzBuzz][fizzbuzz] variant based on [Fibonacci
numbers][fibonacci-number] F(n). Given a natural number n, the program prints
one of the following for every Fibonacci number F(n) from 1 to n:
* *BuzzFizz* if F(n) is prime (a [Fibonacci prime][fibonacci-prime]),
* *FizzBuzz* if F(n) is divisible by 15,
* *Buzz* if F(n) is divisible by 3,
* *Fizz* if F(n) is divisible by 5,
* F(n) otherwise.

The program is implemented in [Go][golang] (golang) and uses the [Miller-Rabin
primality test][miller-rabin] with 64 iterations (see Appendix C.3 of [FIPS PUB
186-4][fips186-4] for a rationale) to detect composite numbers.

[fibonacci-number]: <https://en.wikipedia.org/wiki/Fibonacci_number>
[fibonacci-prime]: <https://en.wikipedia.org/wiki/Fibonacci_prime>
[fips186-4]: <http://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-4.pdf>
[fizzbuzz]: <https://www.rosettacode.org/wiki/FizzBuzz>
[golang]:   <https://golang.org/>
[miller-rabin]: <https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test>
