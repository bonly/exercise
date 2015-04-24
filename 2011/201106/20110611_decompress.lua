zlib = require("zlib");
--ucompress = zlib.inflate(); --解压
--local inflated, eof, bytes_in, bytes_out = uncompress(compressed_string);
--参数 compressed_string 是压缩的数据， 返回的 inflated 是解压后的数据，bytes_in 是压缩数据的长度，bytes_out 是解压后数据的长度
--compress = zlib.deflate(); --压缩
--local deflated, eof, bytes_in, bytes_out = compress(inflated, "full")


file = io.open("data.dat","r")
hp = file:read()
print(hp)
ucompress = zlib.inflate(); --解压
local inflated, eof, bytes_in, bytes_out = ucompress(hp); --解压包体
print("recv body: ");print(inflated);

