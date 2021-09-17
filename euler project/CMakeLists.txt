cmake_minimum_required(VERSION 3.18.4)

project(Project_Euler VERSION 1.0.0)

if (CMAKE_CXX_COMPILER_ID STREQUAL "MSVC")
    add_compile_options(/std:c++latest /Wall)
else()
    add_compile_options(-std=c++20 -Wall -Wextra -pedantic)
endif()

add_executable(problem_1 problem_1.cpp)
add_executable(problem_2 problem_2.cpp)