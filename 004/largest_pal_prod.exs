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

IO.puts LargestPalindrome.product(999, 999)