package main

import (
	"fmt"
)

func main() {
	//按我自己的理解 interface{} 就是其它语言中的 * 类型（任意类型），更深含义，以后再想
	var i interface{}
	var b []byte

	i = " NBA "
	fmt.Print(i.(string)) 		//interface{}转换成别的类型的写法	I.(T)

	
	b = []byte(" CBA ")
	fmt.Print(string(b)) 		//非interface{}类型转换的写法		A(B)
	
	/*上面最后一行的写法，如果换成 b.(string) 的写法，则显示
	 *报错信息: non-interface type []byte on left，这个着实让我这个菜鸟走了一天的弯路，入门教程里面并没有强调这种怪异的写法 
	 *似乎.(type)这种转换方法只适用于interface{}类型
	 */
}
