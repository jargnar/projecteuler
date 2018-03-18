defmodule Multiple do
  def gcd(a, 0), do: a
  def gcd(a, b), do: gcd(b, rem(a, b))
  def lcm(a, b), do: a * div(b, gcd(a, b))
end

IO.puts Enum.reduce(1..20, &Multiple.lcm/2)