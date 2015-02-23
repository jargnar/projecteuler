############################################
# Using Enums                              #
############################################
sum1 = Enum.sum Enum.filter(
    1..999,
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )

############################################
# Using h|t pattern matching and recursion #
############################################
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

sum2 = Multiples.sum Enum.to_list 1..999

###########################################
# Using Enums, Fn Capturing and Pipes     #
###########################################
divs3or5? = &(rem(&1, 3) == 0 or rem(&1, 5) == 0)
sum3 = 1..999
  |> Enum.filter(divs3or5?)
  |> Enum.sum


# How do I better check if all the sums are equal?
if (sum1 == sum2) and (sum2 == sum3) and (sum3 == sum1) do
  IO.puts sum1
end