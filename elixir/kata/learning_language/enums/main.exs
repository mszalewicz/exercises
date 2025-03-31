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

IO.puts("-------------------------------------")

numbers |> Enum.count(fn n -> n >= 4 or n <= 2 end) |> IO.inspect()

IO.puts("-------------------------------------")

randomNumbers = Enum.map(1..12, fn _ -> :rand.uniform(400) end)

randomNumbers
|> Enum.chunk_every(2, 1, :discard)
|> Enum.count(fn input ->
  increasing = input |> Enum.all?(fn [a, b] -> a < b end)
  decreasing = input |> Enum.all?(fn [a, b] -> a > b end)

  valid_differences = input |> Enum.all?(fn [a, b] -> abs(a - b) in 1..3 end)

  (increasing or decreasing) and valid_differences

end)
# |> IO.inspect(charlists: :as_lists)
