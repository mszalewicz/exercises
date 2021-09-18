#include <iostream>

int main()
{
	int sum_of_multiple_of_3_and_5 = 0;

	for(auto i = 3; i < 1000; ++i)
	{
		if(i % 3 == 0 || i % 5 == 0) sum_of_multiple_of_3_and_5 += i;
	}

	std::cout << "\nSum of all the multiples of 3 or 5 below 1000 equals: " << sum_of_multiple_of_3_and_5 << "\n\n"; 

	return 0;
}