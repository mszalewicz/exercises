cmake_minimum_required(VERSION 3.18.4)

project(Project_Euler VERSION 1.0.0)

if (CMAKE_CXX_COMPILER_ID STREQUAL "MSVC")
    add_compile_options(/std:c++latest /Wall)
else()
    add_compile_options(-std=c++20 -Wall -Wextra -pedantic)
endif()

add_executable(problem_001 problem_001.cpp)
add_executable(problem_002 problem_002.cpp)
add_executable(problem_003 problem_003.cpp)
add_executable(problem_004 problem_004.cpp)
add_executable(problem_005 problem_005.cpp)
add_executable(problem_006 problem_006.cpp)