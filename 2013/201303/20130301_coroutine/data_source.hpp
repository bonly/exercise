#ifndef DATA_SOURCE_H
#define DATA_SOURCE_H

#include <boost/coroutine/coroutine.hpp>

template<typename DataType>
class data_source final
{
public:

    typedef DataType data_type;
    typedef boost::coroutines::symmetric_coroutine<void>::call_type call_type;
    typedef boost::coroutines::symmetric_coroutine<void>::yield_type yield_type;

private:

    data_type* m_dest_buffer;
    size_t m_pending = 0;

public:

    data_source() = default;
    data_source(const data_source&) = delete;
    data_source& operator=(const data_source&) = delete;
    data_source(data_source&&) = delete;
    data_source& operator=(data_source&&) = delete;

    template<typename InputIt>
    void write(InputIt begin, InputIt end, call_type& parse_coroutine)
    {
        while (parse_coroutine)
        {
            if (m_pending == 0)
            {
                parse_coroutine();
            }

            if (begin == end)
            {
                break;
            }

            while ((begin != end) && (m_pending != 0))
            {
                *m_dest_buffer++ = *begin++;
                m_pending--;
            }
        }
    }

    void read(data_type* dest_buffer, size_t len, yield_type& yield)
    {
        if (len != 0)
        {
            m_dest_buffer = dest_buffer;
            m_pending = len;
            yield();
        }
    }
};

#endif // DATA_SOURCE_H
