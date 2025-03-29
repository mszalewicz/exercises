file_path = "input"

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

lista1
|> Enum.sort()
|> Enum.zip(lista2 |> Enum.sort())
|> Enum.map(fn {x, y} -> abs(x - y) end)
|> Enum.sum()
|> IO.inspect()
