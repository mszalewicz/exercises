file_path = "input"

# Read file and convert into 2 lists
{lista1, lista2} =
  file_path
  |> File.read!()
  |> String.split("\n", trim: true)
  |> Enum.map(fn line ->
    line
    |> String.split()
    |> Enum.map(&String.to_integer/1)
    |> List.to_tuple()
  end)
  |> Enum.unzip()

# Part 1
lista1
|> Enum.sort()
|> Enum.zip(lista2 |> Enum.sort())
|> Enum.map(fn {x, y} -> abs(x - y) end)
|> Enum.sum()
|> then(fn result -> IO.puts("Part 1: #{result}") end)

# Part 2
lista1
|> Enum.reduce(0, fn x, acc -> (Enum.count(lista2, fn y -> y == x end) * x) + acc end)
|> then(fn result -> IO.puts("Part 2: #{result}") end)
