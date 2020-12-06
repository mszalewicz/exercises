#include <fstream>
#include <iostream>
#include <string>
#include <vector>

class Password {
	private:
		int min_occurences;
		int max_occurences;
		char policy_character;
		std::string password_value;

		std::vector<std::string> split_input(const std::string& input, const std::string& delim) {
			std::vector<std::string> tokens;
			std::size_t prev = 0, pos = 0;
			do {
				pos = input.find(delim, prev);
				if(pos == std::string::npos) pos = input.length();
				std::string token = input.substr(prev, pos-prev);
				if(!token.empty()) tokens.push_back(token);
				prev = pos + delim.length();
			} while (pos < input.length() && prev < input.length());


			return tokens;
		}
		
	public:
		Password(std::string input) {
			//Split object string
			std::vector<std::string> sections = split_input(input, " ");
			//Split rule (min, max) boundaries
			std::vector<std::string> boundaries = split_input(sections[0], "-");
			
			// Populate boundaries
			min_occurences = std::stoi(boundaries[0]);
			max_occurences = std::stoi(boundaries[1]);

			//Get policy character
			policy_character = sections[1].at(0);
			
			//Get password value
		 	password_value = sections[2];	
		}

		int is_correct_password_old_policy() {

			int occurences_counter = 0;
			
			for(unsigned int i = 0; i != password_value.length(); i++) {
				if(password_value.at(i) == policy_character) occurences_counter += 1;
			}

			return (min_occurences <= occurences_counter && occurences_counter <= max_occurences)?1:0;
		}

		int is_correct_password_new_policy() {
			return ((password_value.at(min_occurences-1) == policy_character) != (password_value.at(max_occurences-1) == policy_character))?1:0;
		}


};

std::vector<int> count_correct_passwords(std::vector<Password> collection) {

	int counterOld = 0, counterNew = 0;
	
	for(auto item:collection) {
		counterOld += item.is_correct_password_old_policy();	
		counterNew += item.is_correct_password_new_policy();
	}

	return {counterOld, counterNew};
}

int main() {

	std::fstream filein("input.txt");
        std::string line;
	std::vector<Password> passwords;

        while(std::getline(filein, line)) {
                passwords.push_back(Password(line));
        }

	std::vector<int> answers = count_correct_passwords(passwords);

	std::cout << "First answer: " << answers[0] << std::endl;
	std::cout << "Second answer: " << answers[1] << std::endl;
}	
