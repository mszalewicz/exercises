#pragma once

#include <iostream> // for std::cerr
#include <cstdlib>  // for std::abort

#define ASSERT_MSG(cond, msg) \
    do { \
        if (!(cond)) { \
            std::cerr << "Assertion failed: " << #cond \
                      << ", message: " << msg \
                      << ", file: " << __FILE__ \
                      << ", line: " << __LINE__ << "\n"; \
            std::abort(); \
        } \
    } while(0)