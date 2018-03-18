divs3or5? = &(rem(&1, 3) == 0 or rem(&1, 5) == 0)
sum = 1..999
  |> Enum.filter(divs3or5?)
  |> Enum.sum

IO.puts sum