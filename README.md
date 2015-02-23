# Project Euler, with Elixir

[Project Euler](https://projecteuler.net/) is a series of challenging mathematical/computer programming problems that will require more than just mathematical insights to solve.

[Elixir](http://elixir-lang.org/) is a dynamic, functional language designed for building scalable and maintainable applications.

#### [Multiples of 3 or 5 below 1000](https://projecteuler.net/problem=1)

```
sum = Enum.sum Enum.filter(
    Enum.to_list(1..999),
    fn(x) -> rem(x, 3) == 0 or rem(x, 5) == 0 end
  )
IO.puts sum
```

(TODO: Try the same problem with a different approach.)