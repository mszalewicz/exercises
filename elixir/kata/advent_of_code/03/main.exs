{status, file} = File.read("input")

if status == :error do
  IO.puts("Could not open file: #{status}")
  System.halt(1)
end

# Use regex capture groups to find:
#     1) mul(digits1, digits2)
#     2) count digits1 * digits2
#     3) sum up all result

Regex.scan(~r/mul\((\d+),(\d+)\)/, file)
|> Enum.map(fn [_, num1, num2] -> String.to_integer(num1) * String.to_integer(num2) end)
|> Enum.reduce(0, fn x, acc -> x + acc end)
|> then(fn result -> IO.puts("Part 1: #{result}") end)

firstMul =
  Regex.run(~r/mul\((\d+),(\d+)\)/, file)
  |> then(fn [_, num1, num2] -> String.to_integer(num1) * String.to_integer(num2) end)

  # |> then(fn result -> IO.puts("Part 2: #{result}") end)

restMuls =
  Regex.scan(~r/do\(\).*?mul\((\d+),(\d+)\)/, file)
  |> IO.inspect()
  |> Enum.map(fn [_, num1, num2] -> String.to_integer(num1) * String.to_integer(num2) end)
  |> Enum.reduce(0, fn x, acc -> x + acc end)

result2 = firstMul + restMuls

IO.puts("Part 2: #{result2}")

IO.inspect(firstMul)

# 82733683
# 8980089