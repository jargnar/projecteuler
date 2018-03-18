x = Enum.sum(1..100) |> :math.pow(2)
y = Enum.reduce(1..100, fn(x, acc) -> acc + :math.pow(x, 2) end)

IO.puts x - y