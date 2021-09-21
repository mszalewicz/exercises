#include <iostream>
#include <string>

bool is_palindrome(int &number)
{
	std::string potential_palindrome = std::to_string(number);

	for(auto i = 0; i < potential_palindrome.length() / 2; ++i)
	{
		if(potential_palindrome[i] != potential_palindrome[potential_palindrome.length() - 1 - i])
			return false;
	}

	return true;
}

int main()
{
	// Problem 4

	// A palindromic number reads the same both ways. 
	// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	// Find the largest palindrome made from the product of two 3-digit numbers.
	
	int number = 0, palindrome = 0, lower_bound = 100;

	for(int i = 999; i > lower_bound; --i)
		for(int j = 999; j > lower_bound; --j)
		{
			number = i*j;

			if(is_palindrome(number) and number > palindrome) 
			{
				palindrome = number; 
				lower_bound = j;
				break;
			}
		}

	std::cout << "\n"
			  << "Largest palindrome made from the product of two 3-digit numbers: " 
			  << palindrome
			  << "\n\n";
	return 0;
}