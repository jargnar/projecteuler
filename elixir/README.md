# Project Euler, with Elixir

### [Multiples of 3 or 5 below 1000](https://projecteuler.net/problem=1)

```
sum = Enum.sum Enum.filter(
    Enum.to_list(1..999),
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )
IO.puts sum
```

(TODO: Try the same problem with a different approach.)