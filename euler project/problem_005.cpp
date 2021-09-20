#include <iostream>
int main()
{
	// Problem 5

	// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20

	int number = 2520;

	while(true)
	{
		continue_checking:
		number += 20;

		for(auto i = 2; i <= 20; ++i)
		{
			if(number % i != 0) goto continue_checking;
		}

		break;
	}

	std::cout << "\n"
			  << "Smallest positive number that is evenly divisible by all of the numbers from 1 to 20: " 
			  << number
			  << "\n\n";
	return 0;
}