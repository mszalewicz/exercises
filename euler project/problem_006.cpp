#include <iostream>

int sum_difference(const int &n)
{
	return std::pow(n*(n+1)/2, 2) - (n*(n+1)*(2*n+1))/6;
}

int main()
{
	// Problem 6

	// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	// ---------------------------------------------------------------------------------------------------------------------

	// Solution uses simplified formulas derived as: 
	//	- https://proofwiki.org/wiki/Sum_of_Sequence_of_Cubes
	//	- https://proofwiki.org/wiki/Sum_of_Sequence_of_Squares

	constexpr int bound = 100;

	std::cout << "\n"
			  << "Difference between the sum of the squares and the square of the sum of the first one hundred natural numbers is: " 
			  << sum_difference(bound)
			  << "\n\n";

	return 0;
}