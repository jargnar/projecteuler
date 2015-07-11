defmodule Largest do
  def palindrome_product(100, 100), do: []
  def palindrome_product(i, 100), do: palindrome_product(i - 1, 999)
  def palindrome_product(i, j) do
    p = i * j
    strp = Integer.to_string(p)
    if String.reverse(strp) == strp do
      [p] ++ palindrome_product(i, j - 1)
    else
      palindrome_product(i, j - 1)
    end
  end
end

lpp = Largest.palindrome_product(999, 999) |> Enum.max
IO.puts lpp