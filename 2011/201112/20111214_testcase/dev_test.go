/*
文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
必须import testing这个包
所有的测试用例函数必须是Test开头
测试用例会按照源代码中写的顺序依次执行
测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态
测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。
*/

package gotest

import (
    "testing"
)

func Test_Division_1(t *testing.T) {
    if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
        t.Error("除法函数测试没通过"); // 如果不是如预期的那么就报错
    } else {
        t.Log("第一个测试通过了"); //记录一些你期望记录的信息
    }
}

/*
func Test_Division_2(t *testing.T) {
    t.Error("就是不通过");
}
*/

func Test_Division_3(t *testing.T) {
    if _, e := Division(6, 0); e == nil { //try a unit test on function
        t.Error("Division did not work as expected."); // 如果不是如预期的那么就报错
    } else {
        t.Log("one test passed.", e); //记录一些你期望记录的信息
    }
}  

/*
压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

func BenchmarkXXX(b *testing.B) { ... }
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
文件名也必须以_test.go结尾
*/

func Benchmark_Division1(b *testing.B) {
    for i := 0; i < b.N; i++ { //use b.N for looping 
        Division(4, 5)
    }
}

func Benchmark_TimeConsumingFunction1(b *testing.B) {
    b.StopTimer() //调用该函数停止压力测试的时间计数

    //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
    //这样这些时间不影响我们测试函数本身的性能

    b.StartTimer() //重新开始时间
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}