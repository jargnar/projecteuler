# Project Euler, with Elixir

[Project Euler](https://projecteuler.net/) is a series of challenging mathematical/computer programming problems that will require more than just mathematical insights to solve.

[Elixir](http://elixir-lang.org/) is a dynamic, functional language designed for building scalable and maintainable applications.

#### [Multiples of 3 or 5 below 1000](https://projecteuler.net/problem=1)

Approach 1: Using Enums

```
sum = Enum.sum Enum.filter(
    Enum.to_list(1..999),
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )
IO.puts sum
```

Approach 2: Using h|t pattern matching & recursion
```
defmodule Multiples do
  def sum([]) do
    0
  end
  def sum([h|t]) do
    if rem(h, 3) == 0 or rem(h, 5) == 0 do
      h + sum(t)
    else
      sum(t)
    end
  end
end
```
