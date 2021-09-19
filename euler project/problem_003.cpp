#include <cmath>
#include <iostream>

int main()
{
	// Problem 3

	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?

	unsigned long long number = 600851475143;

	while(number % 2 == 0)
	{
		number = number / 2;
	}

	for(auto i = 3; i <= std::sqrt(number); i += 2)
	{			
		while(number % i == 0)
		{
			number = number / i;
		}
	}

	// At the end we end up with two primes: 'i' and 'number', where i is smaller, hence
	std::cout << "\n"
			  << "Largest prime factor of the number 600851475143: "
			  << number
			  << "\n\n";

	return 0;
}