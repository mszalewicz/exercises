#include <iostream>
#include <vector>

void sieve(std::vector<long long>& primes) noexcept
{
	unsigned int current_position = 0, local_upper_bound = primes.size();

	do
	{
		for(auto i = current_position + 1; i < local_upper_bound; ++i)
		{
			if(primes[i] % primes[current_position] == 0)
			{
				primes.erase(primes.cbegin() + i);
				i -= 1;
				local_upper_bound -= 1;
			}
		}
		current_position += 1;

	} while(current_position < primes.size());
}

long long find_nth_prime(const int& nth) noexcept
{
	int lower_bound = 2, upper_bound = 2*nth;
	std::vector<long long> primes;
	
	do
	{
		for(auto i = lower_bound; i < upper_bound; ++i)
			primes.push_back(i);

		sieve(primes);

		lower_bound = std::move(upper_bound);
		upper_bound = lower_bound + nth;

	} while(primes.size() < nth);
 
	return primes[nth - 1];
}

int main()
{
	// Problem 7

	// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
	// What is the 10 001st prime number?

	// ---------------------------------------------------------------------------------------------------------------------

	// Solution uses Sieve of Eratosthenes algorithm

	constexpr int bound = 10001;

	std::cout << "\n"
			  << "The 10001st prime number: " 
			  << find_nth_prime(bound)
			  << "\n\n";

	return 0;
}