#include "isogram.h"

namespace isogram {
	bool is_isogram(std::string to_test){
		int string_length = to_test.length();

		for(int i=0; i<string_length; i++) {to_test[i]=std::tolower(to_test[i]);}
		
		for(int i=0; i<string_length-1; i++) {
			if(to_test[i]=='-' or to_test[i]==' ') {continue;}
			
			for(int j=i+1; j<string_length; j++) {
				if(to_test[i]==to_test[j]) {
					return false;
				}
			}
		}
		return true;
	}
}  // namespace isogram
