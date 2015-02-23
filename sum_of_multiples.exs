# Using Enums
sum = Enum.sum Enum.filter(
    Enum.to_list(1..999),
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )

# Using some elixir h|t pattern matching
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

# if method 1 == method 2
# print the result (either one)
if sum == Multiples.sum(Enum.to_list 1..999) do
  IO.puts sum
end