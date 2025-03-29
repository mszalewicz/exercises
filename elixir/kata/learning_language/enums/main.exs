numbers = [1, 2, 3, 4, 5]

numbers
|> Enum.map(fn x -> x * x end)
# anonymous function
|> Enum.each(&IO.puts(&1))

IO.puts("-------------------------------------")

numbers
|> Enum.filter(fn x -> rem(x, 2) == 0 end)
|> Enum.each(&IO.puts(&1))

IO.puts("-------------------------------------")

numbers
|> Enum.filter(&(rem(&1, 2) == 0))
|> IO.inspect()

# IO.puts("\n")
IO.puts("-------------------------------------")

numbers
|> Enum.reduce(&(&1 * &2))
|> IO.puts()

IO.puts("-------------------------------------")

test =
  numbers
  |> Enum.reduce(fn x, acc -> x * acc end)

IO.inspect(test)
