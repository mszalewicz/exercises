{status, result} = File.read("input")

if status == :error do
  IO.puts("Could not open file: #{inspect(status)}")
  System.halt(1)
end

defmodule AOC do
  def isSafe?(list) do # does list meet the conditions ?
    consequtive_tuples = Enum.chunk_every(list, 2, 1, :discard)
    increasing = consequtive_tuples |> Enum.all?(fn [a, b] -> a < b end)
    decreasing = consequtive_tuples |> Enum.all?(fn [a, b] -> a > b end)

    valid_differences = consequtive_tuples |> Enum.all?(fn [a, b] -> abs(a - b) in 1..3 end)

    (increasing or decreasing) and valid_differences
  end

  def isSafeWithOneBad?(list) do # does list meet the conditions if we delete one value ?
    if isSafe?(list) do
      true
    else
      Enum.any?(0..(length(list) - 1), fn i ->
        list |> List.delete_at(i) |> isSafe?()
      end)
    end
  end
end

# Prepare lists of integers
integers =
  result
  |> String.split("\n", trim: true)
  |> Enum.map(fn line ->
    line
    |> String.split()
    |> Enum.map(&String.to_integer/1)
  end)

#  Part 1
integers
|> Enum.count(&AOC.isSafe?/1)
|> then(fn result -> IO.puts("Part 1: #{result}") end)

# Part 2
integers
|> Enum.count(&AOC.isSafeWithOneBad?/1)
|> then(fn result -> IO.puts("Part 2: #{result}") end)
