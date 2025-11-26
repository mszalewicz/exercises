#include "../assert.h"
#include <algorithm>
#include <cassert>
#include <cstdlib>
#include <sstream>
#include <string>
#include <vector>

/*
 *
 * TODO:
 *
 * You are given two strings word1 and word2. Merge the strings by adding letters
 * in alternating order, starting with word1. If a string is longer than the other,
 * append the additional letters onto the end of the merged string.
 *
 * Return the merged string.
 *
 * Constraints:
 *
 *  1 <= word1.length, word2.length <= 100
 *  word1 and word2 consist of lowercase English letters.
 *
 */


/*
 *
 * NOTES:
 *  - using string builder as optimization
 *  - iterating over size of bigger word, as we can just stop iteration of short word with cheap comparison
 *  - unfortunately did not find solution without branch predictions (confirmed with godbolt)
 *
 */

class Solution {
public:
  static std::string mergeAlternately(std::string word1, std::string word2) {
    assert(0 <= word1.size() && word1.size() <= 100);
    assert(0 <= word2.size() && word2.size() <= 100);

    int maxSize = std::max(word1.size(), word2.size());

    std::ostringstream builder;

    for (auto i = 0; i < maxSize; ++i) {
        if (i < word1.size())
            builder << word1[i];
        if (i < word2.size())
            builder << word2[i];
    }

    return builder.str();
  }
};

// TESTING

struct Test {
  std::string word1;
  std::string word2;
  std::string result;
};

int main() {
  std::vector<Test> tests = {
      {"abc", "pqr", "apbqcr"},
      {"ab", "pqrs", "apbqrs"},
      {"abcd", "pq", "apbqcd"},
  };

  for (const auto &test : tests) {
    ASSERT_MSG(
        Solution::mergeAlternately(test.word1, test.word2) == test.result,
        "word1: " + test.word1 + ", word2: " + test.word2 + ", merge result: " + test.result
    );
  }

  std::cout << "Tests pass correctly.\n";
}
