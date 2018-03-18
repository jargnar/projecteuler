defmodule Multiples do
  def sum([]), do: 0
  def sum([h|t]) do
    if rem(h, 3) == 0 or rem(h, 5) == 0, do: h + sum(t), else: sum(t)
  end
end

IO.puts Multiples.sum Enum.to_list 1..999