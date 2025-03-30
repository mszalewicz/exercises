{status, result} = File.read("input")

if status == :error do
  IO.puts("Could not open file: #{inspect(status)}")
  System.halt(1)
end

result
|> String.split("\n", trim: true)
|> Enum.map(fn line ->
  line
  |> String.split()
  |> Enum.map(&String.to_integer/1)
end)
|> Enum.count(fn numbers ->
  consequtive_tuples = Enum.chunk_every(numbers, 2, 1, :discard)
  increasing = consequtive_tuples |> Enum.all?(fn [a, b] -> a < b end)
  decreasing = consequtive_tuples |> Enum.all?(fn [a, b] -> a > b end)

  valid_differences = consequtive_tuples |> Enum.all?(fn [a, b] -> abs(a - b) in 1..3 end)

  (increasing or decreasing) and valid_differences
end)
|> then(fn result -> IO.puts("Part 1: #{result}") end)

