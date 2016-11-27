/*
 *http://yatb.giacomodrago.com/en/post/16/boost_coroutines_instead_of_state_machines_maybe.html
 */
#include "data_source.hpp"

#include <functional>
#include <iostream>

void parse(data_source<char>& s, data_source<char>::yield_type& yield)
{
    char message[255];
    uint8_t message_length;

    do
    {
        s.read(reinterpret_cast<char*>(&message_length), 1, yield);

        if (message_length > 0)
        {
            std::cout << "About to parse a message of length " << (+message_length) << std::endl;
            s.read(message, message_length, yield);

            std::cout << "Parsed message: '";
            std::cout.write(message, message_length);
            std::cout << "'" << std::endl;
        }
    }
    while (message_length > 0);

    std::cout << "Parsed final empty message" << std::endl;
}

template<size_t N>
void write(data_source<char>& s, const char (&data)[N], data_source<char>::call_type& parse_coroutine)
{
    s.write(data, data + N - 1, parse_coroutine);
}

int main()
{
    data_source<char> s;

    {
        data_source<char>::call_type parse_coroutine(std::bind(parse, std::ref(s), std::placeholders::_1));

        write(s, "\x3", parse_coroutine);
        write(s, "ab", parse_coroutine);
        write(s, "c" "\x8", parse_coroutine);
        write(s, "1234567", parse_coroutine);
        write(s, "8" "\x0", parse_coroutine);

        assert(!parse_coroutine);
    }
    std::cout << "-------------------------------------------------------" << std::endl;
    {
        data_source<char>::call_type parse_coroutine(std::bind(parse, std::ref(s), std::placeholders::_1));

        write(s, "\x3" "abc" "\x8" "12345678" "\x0", parse_coroutine);

        assert(!parse_coroutine);
    }
    std::cout << "-------------------------------------------------------" << std::endl;
    {
        data_source<char>::call_type parse_coroutine(std::bind(parse, std::ref(s), std::placeholders::_1));

        const char input[] = "\x3" "abc" "\x8" "12345678" "\x0";

        for (size_t i = 0; i < sizeof(input) - 1; i++)
        {
            s.write(input + i, input + i + 1, parse_coroutine);
        }

        assert(!parse_coroutine);
    }

    return 0;
}
