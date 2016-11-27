#include "data_source.hpp"

#include <functional>
#include <iostream>
#include <chrono>

#include <boost/coroutine/all.hpp>
#include <boost/optional.hpp>

void coroutine_parser_func(data_source<char>& s, data_source<char>::yield_type& yield)
{
    char message[255];
    uint8_t message_length;

    do
    {
        s.read((char*)&message_length, 1, yield);
        s.read(message, message_length, yield);
    }
    while (message_length > 0);
}

void coroutine_parser_reusable_func(data_source<char>& s, data_source<char>::yield_type& yield)
{
    char message[255];
    uint8_t message_length;

    while (1)
    {
        s.read((char*)&message_length, 1, yield);
        s.read(message, message_length, yield);
    }
}

class state_machine_parser
{
private:

    char m_message[255];
    uint8_t m_message_length = 0;
    uint8_t m_message_offset = 0;

    enum { MESSAGE_HEADER, MESSAGE_BODY, COMPLETE } m_state = MESSAGE_HEADER;

public:

    void process(char c)
    {
        switch (m_state)
        {
        case MESSAGE_HEADER:
            m_message_length = c;
            if (m_message_length == 0)
            {
                m_state = COMPLETE;
            }
            else
            {
                m_message_offset = 0;
                m_state = MESSAGE_BODY;
            }
            break;

        case MESSAGE_BODY:
            m_message[m_message_offset++] = c;
            if (m_message_offset == m_message_length)
            {
                m_state = MESSAGE_HEADER;
            }
            break;

        case COMPLETE:
            break;
        }
    }

    bool complete() const
    {
        return m_state == COMPLETE;
    }
};

static boost::optional<boost::coroutines::stack_context> cached_context;
static boost::coroutines::standard_stack_allocator allocator;

class cached_stack_allocator
{
public:

    typedef boost::coroutines::standard_stack_allocator::traits_type traits_type;

    void allocate(boost::coroutines::stack_context& ctx, size_t size)
    {
        if (!cached_context)
        {
            allocator.allocate(ctx, size);
            cached_context = ctx;
        }
        else
        {
            assert(size == cached_context->size);
            ctx = *cached_context;
        }
    }

    void deallocate(boost::coroutines::stack_context&)
    {
    }
};

static const size_t ROUNDS = 1000000;

int main()
{
    data_source<char> s;
    const auto coroutine_parser = std::bind(coroutine_parser_func, std::ref(s), std::placeholders::_1);

    const char input[] = "\x3" "abc" "\x8" "12345678" "\x0";

    std::chrono::time_point<std::chrono::high_resolution_clock> start, end;

    //----------------------------------------------------------------------

    std::cout << "Coroutine parser: ";

    {
        start = std::chrono::system_clock::now();
        for (size_t i = 0; i < ROUNDS; i++)
        {
            data_source<char>::call_type parse_coroutine(coroutine_parser);

            for (size_t i = 0; i < sizeof(input) - 1; i++)
            {
                s.write(input + i, input + i + 1, parse_coroutine);
            }

            assert(!parse_coroutine);
        }
        end = std::chrono::system_clock::now();
    }

    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << " ms" << std::endl;

    //----------------------------------------------------------------------

    std::cout << "Coroutine parser (reusable stack): ";

    {
        cached_stack_allocator stack_allocator;

        start = std::chrono::system_clock::now();
        for (size_t i = 0; i < ROUNDS; i++)
        {
            data_source<char>::call_type parse_coroutine(coroutine_parser, boost::coroutines::attributes(), stack_allocator);

            for (size_t i = 0; i < sizeof(input) - 1; i++)
            {
                s.write(input + i, input + i + 1, parse_coroutine);
            }

            assert(!parse_coroutine);
        }
        end = std::chrono::system_clock::now();
    }

    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << " ms" << std::endl;

    //----------------------------------------------------------------------

    std::cout << "Coroutine parser (reusable stack, reusable coroutine): ";

    {
        cached_stack_allocator stack_allocator;

        start = std::chrono::system_clock::now();
        data_source<char>::call_type parse_coroutine(std::bind(coroutine_parser_reusable_func, std::ref(s), std::placeholders::_1),
            boost::coroutines::attributes(), stack_allocator);
        for (size_t i = 0; i < ROUNDS; i++)
        {
            for (size_t i = 0; i < sizeof(input) - 1; i++)
            {
                s.write(input + i, input + i + 1, parse_coroutine);
            }
        }
        end = std::chrono::system_clock::now();
    }

    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << " ms" << std::endl;

    //----------------------------------------------------------------------

    std::cout << "State machine parser: ";

    {
        start = std::chrono::system_clock::now();
        for (size_t i = 0; i < ROUNDS; i++)
        {
            state_machine_parser parser;

            for (size_t i = 0; i < sizeof(input) - 1; i++)
            {
                parser.process(input[i]);
            }

            assert(parser.complete());
        }
        end = std::chrono::system_clock::now();
    }

    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << " ms" << std::endl;

    //----------------------------------------------------------------------

    return 0;
}
