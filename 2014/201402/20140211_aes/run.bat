@echo off
rem 参数：第一个是obd路径，第二个是该obd文件的编码用于解决乱码，默认gbk，第三个是否debug，默认开启，传nodebug关闭
java -cp .;commons-codec-1.9.jar com.test.T 123 1234567890123456
PAUSE