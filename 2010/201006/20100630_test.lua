nGlobal = 10 --一个全局的整形变量
strGlobal = "hello i am in lua" --一个全局的字符串变量
--一个返回值为int类型的函数
function add(a, b)
   return a+b
end
--一个返回值为string类型的函数
function strEcho(a)
   print(a .. 10) 
   return 'haha i have print your input param'
end
cppapi.testFunc() --调用c++暴露的一个测试函数

t={name='ettan', age=23, desc='正值花季年龄'}
