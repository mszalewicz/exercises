#include <iostream>

int main()
{
	// Problem 1

	// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
	// The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.

	int sum_of_multiple_of_3_and_5 = 0;

	for(auto i = 3; i < 1000; ++i)
	{
		if(i % 3 == 0 || i % 5 == 0) sum_of_multiple_of_3_and_5 += i;
	}

	std::cout << "\nSum of all the multiples of 3 or 5 below 1000 equals: " << sum_of_multiple_of_3_and_5 << "\n\n"; 

	return 0;
}