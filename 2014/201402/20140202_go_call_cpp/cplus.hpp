#pragma once

class A {
    public:
        virtual void test();
};

class B: public A {
    public:
        void test();
};