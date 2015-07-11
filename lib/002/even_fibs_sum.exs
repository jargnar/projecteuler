fibs = Stream.unfold({0, 1}, fn {a, b} -> {a, {b, a + b}} end)
sum = fibs
      |> Stream.reject(fn(x) -> rem(x, 2) != 0 end)
      |> Enum.take_while(fn(x) -> x < 4000000 end)
      |> Enum.sum
IO.puts sum