#include <cstdio>
#include <cstdlib>

struct BinlogEventHeader
{
    int  timestamp;
    unsigned char type_code;
    int  server_id;
    int  event_length;
    int  next_position;
    short flags;
};

int main()
{
    FILE* fp = fopen("/usr/local/var/mysql/master-bin.000001", "rb");
    int magic_number;
    fread(&magic_number, 4, 1, fp);

    printf("%d - %s\n", magic_number, (char*)(&magic_number));

    struct BinlogEventHeader format_description_event_header;
    fread(&format_description_event_header, 19, 1, fp);

    printf("BinlogEventHeader\n{\n");
    printf("    timestamp: %d\n", format_description_event_header.timestamp);
    printf("    type_code: %d\n", format_description_event_header.type_code);
    printf("    server_id: %d\n", format_description_event_header.server_id);
    printf("    event_length: %d\n", format_description_event_header.event_length);
    printf("    next_position: %d\n", format_description_event_header.next_position);
    printf("    flags[]: %d\n}\n", format_description_event_header.flags);

    short binlog_version;
    fread(&binlog_version, 2, 1, fp);
    printf("binlog_version: %d\n", binlog_version);

    char server_version[51];
    fread(server_version, 50, 1, fp);
    server_version[50] = '\0';
    printf("server_version: %s\n", server_version);

    int create_timestamp;
    fread(&create_timestamp, 4, 1, fp);
    printf("create_timestamp: %d\n", create_timestamp);

    char header_length;
    fread(&header_length, 1, 1, fp);
    printf("header_length: %d\n", header_length);

    int type_count = format_description_event_header.event_length - 76;
    unsigned char post_header_length[type_count];
    fread(post_header_length, 1, type_count, fp);
    for(int i = 0; i < type_count; i++) 
    {
        printf("  - type %d: %d\n", i + 1, post_header_length[i]);
    }

    return 0;
}
/*
g++ 编译时候需要加上 -fpack-struct=1 这个参数
*/

/*
一共会输出 40 种类型（从 1 到 40），如官方文档所说，这个数组从 START_EVENT_V3 事件开始（type_code 是 1）。

跳转事件

　　跳转事件即 ROTATE_EVENT，其 type_code 是 4，其 post-header 长度为 8。

　　当一个 Binlog 文件大小已经差不多要分割了，它就会在末尾被写入一个 ROTATE_EVENT——用于指出这个 Binlog 的下一个文件。

　　它的 post-header 是 8 字节的一个东西，内容通常就是一个整数 4，用于表示下一个 Binlog 文件中的第一个事件起始偏移量。我们从上文就能得出在一般情况下这个数字只可能是四，就偏移了一个魔法数字。当然我们讲的是在 v4 这个 Binlog 版本下的情况。

　　然后在 payload 位置是一个字符串，即下一个 Binlog 文件的文件名。

各种不同的事件体

　　由于篇幅原因这里就不详细举例其它普通的不同事件体了，具体的详解在 MySQL 文档中一样有介绍，用到什么类型的事件体就可以自己去查询。
http://dev.mysql.com/doc/internals/en/event-data-for-specific-event-types.html

http://dev.mysql.com/doc/internals/en/binary-log.html
http://dev.mysql.com/doc/internals/en/event-classes-and-types.html
*/