### Smallest multiple

#### Problem

https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

#### Solution

We'll be using the formula `LCM(a, b) = a * b / GCD(a, b)`, where `LCM` stands for `Lowest Common Multiple` and `GCD` stands for `Greatest Common Divisor`.

We'll use the [Euclidean algorithm](https://en.wikipedia.org/wiki/Euclidean_algorithm#Procedure) for finding `GCD(a, b)`.


```
defmodule Multiple do
    def gcd(a, 0), do: a
    def gcd(a, b), do: gcd(b, rem(a, b))
    def lcm(a, b), do: a * div(b, gcd(a, b))
end

IO.puts Enum.reduce(1..20, &Multiple.lcm/2)
```