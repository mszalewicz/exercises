#include <iostream>

int next_fibonacci (int &previous_value1, int &previous_value2)
{
	return previous_value1 + previous_value2;
}

int main()
{
	constexpr int upper_bound = 4000000;
	int sum_of_even_fibonacci_numbers_up_to_upper_bound = 2, previous_value1 = 1, previous_value2 = 2, new_fibonacci_number;

	while(true)
	{
		new_fibonacci_number = next_fibonacci(previous_value1, previous_value2);

		previous_value1 = previous_value2;
		previous_value2 = new_fibonacci_number;

		if(new_fibonacci_number > upper_bound) break;
		if(new_fibonacci_number % 2 == 0) sum_of_even_fibonacci_numbers_up_to_upper_bound += previous_value2;
	}

	std::cout << "\nSum of the even-valued terms in the Fibonacci sequence, whose values do not exceed four million: " << sum_of_even_fibonacci_numbers_up_to_upper_bound << "\n\n"; 

	return 0;
}