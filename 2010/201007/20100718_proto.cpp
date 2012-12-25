/**
 * @file 20100718_proto.cpp
 * @brief 
 * @author bonly
 * @date 2012-11-19 bonly Created
 */

#include "message.pb.h"
#include <pb_encode.h>
#include <pb_encode.c>
int main()
{
    Example mymessage = {42};
    uint8_t buffer[10];
    pb_ostream_t stream = pb_ostream_from_buffer (buffer, sizeof(buffer));
    pb_encode (&stream, Example_fields, &mymessage);

    return 0;
}


