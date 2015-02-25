### Largest palindrome product

#### Problem

https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

#### Solution

I'm sure there's another clever way, but until then...

```
defmodule LargestPalindrome do
    def product(100, 100), do: nil
    def product(i, 100), do: product(i - 1, 999)
    def product(i, j) do
        p = i * j
        strp = Integer.to_string(p)
        if String.reverse(strp) == strp do
            p
        else
            product(i, j - 1)
        end
    end
end

LargestPalindrome.product(999, 999)
```
