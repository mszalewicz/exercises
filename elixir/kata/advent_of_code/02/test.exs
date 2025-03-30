{result, content} = File.read("input")

if result == :error do
  IO.puts("Could not open file: #{inspect(result)}")
  System.halt(1)
end

  content
  |> String.split("\n", trim: true)
  |> Enum.map(fn line ->
    line
    |> String.split()
    |> Enum.map(&String.to_integer/1)
  end)
|> Enum.count(fn set ->
  number_tuples = set |> Enum.chunk_every(2, 1, :discard)
  increasing = number_tuples |> Enum.all?(fn [a, b] -> a < b end)
  decreasing = number_tuples |> Enum.all?(fn [a, b] -> a > b end)
  valid_diffrences = number_tuples |> Enum.all?(fn [a, b] -> abs(a - b) in 1..3 end)

  (increasing or decreasing) and valid_diffrences
end)
|> then(fn result -> IO.puts("Part 1: #{result}") end)



# |> IO.inspect(charlists: :as_lists)