#include <string>
#include <vector>
#include <fstream>
#include <iostream>

int SumOfTwoNumbers(std::vector<int> values) {

	for(std::vector<std::string>::size_type i = 0; i != values.size(); i++){
		for (std::vector<std::string>::size_type j = 0; j != values.size(); j++) {
			if(values[i]+values[j] == 2020) return values[i]*values[j]; 
		}
	}
 
	throw std::runtime_error("Did not find the correct values for the first task.");
}


int SumOfThreeNumbers(std::vector<int> values) {

	for(std::vector<std::string>::size_type i = 0; i != values.size(); i++){
		for (std::vector<std::string>::size_type j = 0; j != values.size(); j++) {
			for(std::vector<std::string>::size_type z = 0; z != values.size(); z++) {
				if(values[i]+values[j]+values[z] == 2020) return values[i]*values[j]*values[z];
			}
		}
	}
 
	throw std::runtime_error("Did not find the correct values for the first task.");
}

int main() {
	
	std::fstream filein("input.txt");
	std::vector<int> values;

	std::string line;

	while(std::getline(filein, line)) {
		const int entry = std::stoi(line);
		values.push_back(entry);
	}

	std::cout << "First answer: " << SumOfTwoNumbers(values) << std::endl;	
	std::cout << "Second answewr: " << SumOfThreeNumbers(values) << std::endl;

}
