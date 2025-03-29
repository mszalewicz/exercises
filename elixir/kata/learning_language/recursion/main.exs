# -- With onel line function notation:
defmodule MathTest do
  def power(_base, 0), do: 1

  def power(base, exponent) do
    power(base, exponent - 1) * base
  end
end

# -- With if statement:
# defmodule MathTest do
#   def power(base, exponent) do
#     if exponent == 0 do
#       1
#     else
#       power(base, exponent - 1) * base
#     end
#   end
# end

IO.puts(MathTest.power(2, 3))
