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

largest = Factorizer.factorize(600851475143, 2) |> Enum.max

IO.puts largest