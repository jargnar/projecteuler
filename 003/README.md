### Largest prime factor

#### Problem

https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

#### Solution

Good old recursion. I wonder if there's a better way...

```
defmodule Factorizer do
    def factorize(1, i) do
        []
    end
    def factorize(n, i) do
        if rem(n, i) == 0 do
            [i] ++ factorize(div(n, i), i)
        else
            factorize(n, i + 1)
        end
    end
end

Factorizer.factorize(600851475143, 2) |> Enum.max
```