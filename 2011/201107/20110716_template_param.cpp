/*
have several kinds of template parameters

Type Parameters.
Types
Templates (only classes, no functions)
Non-type Parameters
Pointers
References
Integral constant expressions
What you have there is of the last kind. It's a compile time constant (so-called constant expression) and is of type integer or enumeration. After looking it up in the standard, i had to move class templates up into the types section - even though templates are not types. But they are called type-parameters for the purpose of describing those kinds nonetheless. You can have pointers (and also member pointers) and references to objects/functions that have external linkage (those that can be linked to from other object files and whose address is unique in the entire program). Examples:

Template type parameter:

template<typename T>
struct Container {
    T t;
};

// pass type "long" as argument.
Container<long> test;
Template integer parameter:

template<unsigned int S>
struct Vector {
    unsigned char bytes[S];
};

// pass 3 as argument.
Vector<3> test;
Template pointer parameter (passing a pointer to a function)

template<void (*F)()>
struct FunctionWrapper {
    static void call_it() { F(); }
};

// pass address of function do_it as argument.
void do_it() { }
FunctionWrapper<&do_it> test;
Template reference parameter (passing an integer)

template<int &A>
struct SillyExample {
    static void do_it() { A = 10; }
};

// pass flag as argument
int flag;
SillyExample<flag> test;
Template template parameter.

template<template<typename T> class AllocatePolicy>
struct Pool {
    void allocate(size_t n) {
        int *p = AllocatePolicy<int>::allocate(n);
    }
};

// pass the template "allocator" as argument. 
template<typename T>
struct allocator { static T * allocate(size_t n) { return 0; } };
Pool<allocator> test;
A template without any parameters is not possible. But a template without any explicit argument is possible - it has default arguments:

template<unsigned int SIZE = 3>
struct Vector {
    unsigned char buffer[SIZE];
};

Vector<> test;
Syntactically, template<> is reserved to mark an explicit template specialization, instead of a template without parameters:

template<>
struct Vector<3> {
    // alternative definition for SIZE == 3
};

*/