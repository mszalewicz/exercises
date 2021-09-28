#include <iostream>


int main()
{
	// Problem 9

	// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a2 + b2 = c2
	// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
	// There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product abc.

	// ---------------------------------------------------------------------------------------------------------------------

	// Additionaly: 
	// a + b > c && a + b + c = 1000, hence:
	// a + b = 1000 - c <=> 1000 - c > c <=> 1000 > 2c <=> 500 > c

	// Used properties:
	// 1) a^2 + b^2 + c^2
	// 2) a < b < c
	// 3) a + b + c = 1000
	// 4) c < 500

	unsigned long long result = 1;

	for(auto c = 499; c > 5; --c)
		for(auto b = c - 1; b > 4; --b)
		{
			auto a = 1000 - c - b; 

			if(a*a + b*b == c*c)
			{
				result = a*b*c;
				goto print_result;
			}			
		}

	print_result:

	std::cout << "\n"
			  << "Product: " 
			  << result
			  << "\n\n";

	return 0;
}