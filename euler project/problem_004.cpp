#include <algorithm>
#include <iostream>
#include <string>

bool is_palindrome1(int &number)
{
	std::string potential_palindrome = std::to_string(number);

	for(auto i = 0; i < potential_palindrome.length() / 2; ++i)
	{
		if(potential_palindrome[i] != potential_palindrome[potential_palindrome.length() - 1 - i])
			return false;
	}

	return true;
}

// bool is_palindrome2(int number)
// {
// 	std::string potential_palindrome = std::to_string(number);

// 	for(auto i = 0; i < potential_palindrome.length() / 2; ++i)
// 	{
// 		if(potential_palindrome[i] != potential_palindrome[potential_palindrome.length() - 1 - i])
// 			return false;
// 	}

// 	return true;
// }

// number - 998001
// answers - 906609

int construct_next_number(int &first_half)
{
	std::string first_half_string = std::to_string(first_half);
	std::string second_half_string = first_half_string;
	std::reverse(second_half_string.begin(), second_half_string.end());

	return std::stoi(first_half_string + second_half_string);
}

bool correct_factorization(int number_to_factor)
{
	
	return false;
}

int main()
{

	int first_half = 998;

	while(first_half > 990)
	{
		// if(is_palindrome2(construct_next_number(first_half))) break
		std::cout << construct_next_number(first_half) << "\n";
		--first_half;
	}



	// Problem 4

	// A palindromic number reads the same both ways. 
	// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	// Find the largest palindrome made from the product of two 3-digit numbers.
	
	// int number = 0, palindrome = 0, lower_bound = 100;

	// for(int i = 999; i > lower_bound; --i)
	// 	for(int j = 999; j > lower_bound; --j)
	// 	{
	// 		number = i*j;

	// 		if(is_palindrome1(number) and number > palindrome) 
	// 		{
	// 			palindrome = number; 
	// 			lower_bound = j;
	// 			break;
	// 		}
	// 	}

	// std::cout << "\n"
	// 		  << "Largest palindrome made from the product of two 3-digit numbers: " 
	// 		  << palindrome
	// 		  << "\n\n";
	return 0;
}