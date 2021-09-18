#include <cmath>
#include <iostream>

// Problem resolves the prime factorization (https://en.wikipedia.org/wiki/Integer_factorization)

int main()
{
	int biggest_prime = 1;
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
	std::cout << "\nLargest prime factor of the number 600851475143: " << number << "\n\n";

	return 0;
}