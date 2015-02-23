### Sum of multiples of 3 or 5 below 1000

Read the problem description at https://projecteuler.net/problem=1

Here are three different approaches that I could think of; which one do you think is the most elegant?

Approach 1: Using Enums

```
sum = Enum.sum Enum.filter(
    Enum.to_list(1..999),
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )
IO.puts sum
```

Approach 2: Using Enums, Pipes and Captures
```
divs3or5? = &(rem(&1, 3) == 0 or rem(&1, 5) == 0)
sum = 1..999
  |> Enum.filter(divs3or5?)
  |> Enum.sum
IO.puts sum
```

Approach 3: Using h|t pattern matching & recursion
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

sum = Multiples.sum(Enum.to_list 1..999)
IO.puts sum
```