require("bz2")
local compressor = bz2.initCompress()
local decompressor = bz2.initDecompress()

cmstr = compressor:update('test')
--cmstr = compressor:flush()
cmstr = compressor:update()
compressor:close()

print(string.len(cmstr))

resultString, remaining = decompressor:update(cmstr)
decompressor:close()
print(resultString)

file=io.open("data.dat","w")
file:write(cmstr)
file:close()

--https://github.com/harningt/lua-bz2