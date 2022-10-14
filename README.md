# Golang Study Note

- 采用 <a href="https://github.com/FOS-Lover/Golang-Study-Notes/blob/master/LICENSE">MIT</a> 协议
- 更新于 : 2022年10月14日
- 不足地方或错误地方欢迎fork提交

### 常用命令

- `go mod init note` # 初始化包
- `go get github.com/go-sql-driver/mysql` # 下载并安装包和依赖
- `go install` # 编译并安装包和依赖
- `go list` # 列出包
- `go env` # 打印go的环境信息
- `go run build main.go`  # 编译包和依赖
- `go run main.go` # 编译并运行go程序
- `go clean` # 移除对象文件
- `go tool` # 查看go tool工具
- `go version` # 查看go版本



### 初始化项目

- `go mod init projectName`

### 模块的导入

- #### 创建模块
```go
package user

func User() string {
	return "hello user"
}
```

- #### 导入模块
```go
package main

import (
	"fmt"
	"note/user"
)

func main() {
	data := user.User()
	fmt.Println(data)
}
```

### 标识符, 关键字, 命名规则

- #### 标识符的组成和命名规则
  - 标识符由数字，字母和下划线(_)组成
  - 只能以字母和下划线(_)开头
  - 标识符区分大小写


- #### 关键字

| break    | default     | func   | interface | select |
|----------|-------------|--------|-----------|--------|
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

- #### 命名规范

  - ##### 包名称
    - 保持和package的名字和目录名一直，尽量采取有意义的包名，简短，有意义，尽量不要和标准库冲突，包名应该为小写单词，不要混合下划线和大小写
    - ```go
      package service
      package dao
      
  - ##### 文件命名
    - 尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词
    - ```customer_service.go```

  - ##### 结构体命名
    - 采用驼峰命名法，首字母根据访问控制大写和小写
    - `struct` 声明和初始化格式采用多行
    - ```go
      type CustomerOrder struct {
        Name string
        Address string
      }
      order := CustomerOrder{"tom", "China"}
      
  - ##### 接口命名
    - 单个函数的结构名以`er`作为后缀, 例如`Reader`,`Writer`
    - ```go
      type Reader interface {
        Read(p []byte) (n int, err error)
      }
  
  - ##### 变量命名
    - 遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循一下规则
    - 如果变量为私有，且特有名词为首个单词，则使用小写，如 appService 若变量类型为bool类型，则名称应以Has, Is, Can 或 Allow 开头
    - ```go
      var isExist bool
      var hasConflict bool
      var canManage bool
      var allowGitHook bool

  - ##### 常量命名
    - 常量均需使用全部大写字母组成，并使用下划线分词
    - ```const APP_URL = "https://baidu.com"```
    - 如果时枚举类型的常量，需要创建相应类型
    - ```go
      type Scheme string
      const {
        HTTP Scheme = "http"
        HTTPS Scheme = "https"
      }

  - ##### 错误处理
    - 错误处理的原则就是不能丢弃任何有返回err的调用，不要使用_丢弃，必须全部处理。接收到错误，要么返回err，或者使用log记录下来尽早return：一旦有错误发生，马上返回，尽量不要使用panic，除非知道你在做什么，错误描述如果时英文必须为小写，不需要标点结尾，采用独立的错误流进行处理
    - ```go
      if err != nil {
        // 错误处理
        return // 或者继续 
      }
      // 正常代码
    
  - 单元测试
    - 单元测试文件名命名规范为`example_test.go`测试用例的函数名称必须以Test开头，例如`TestExample`每个重要的函数都要首先编写测试用例，测试应用和正规代码一起提交方便进行回归测试.

### 变量

- #### 变量语法
  - `var` 声明变量的关键字
  - `identifier` 变量名称
  - `type` 变量类型
  - ```go
    var identifier type

- #### 例如

```go
package main

func main() {
    var name string
    var age int
    var email string
}
```

- #### 批量声明
  - ```go
    var (
        name string
        age int
        win false
    )

- #### 批量初始化声明
  - ```go
    var name, age, win = "tom", 20, true

- #### 类型推断声明
  - ```go
    var name = "tom"
    var age = 20
    var win = false
    
- #### 短变量声明
  - 只能在函数内部声明和初始化
  - ```go
    name := "tom"
    age := 20
    win := true

- #### 匿名变量
  - ```go
    func getNameAndAge() (string, int){
        return "tom", 20
    }
    // _ 也为匿名变量
    func main(){
        _, age := getNameAndAge()
    }
    
### 常量

- #### 常量语法
  - `const` 定义常量关键字
  - `constantName` 常量名称
  - `type` 常量类型
  - `value` 常量的值
  - ```go 
    const constantName [type] = value
    
- #### 实例
```go
package main

func main()  {
    const PI float64 = 3.14
	const PI2 = 3.14159
	
	const (
		width = 100
		height = 100
    )
	
	const i, j = 1, 2
	const a, b, c = 10, 20, "tom"
}
```

- #### iota 关键字
  - 可修改的常量关键字
  - ```go
    const (
      a1 = iota // 0
      a2 = iota // 1
      a3 = iota // 2
    )

- #### _ 跳值
  - ```go
    const (
        a1 = iota // 0
        _ // 1
        a2 = iota // 2
    )

- #### itoa 中间插队
  - 与跳值作用一样，但中间值改变了
  - ```go
    const (
       a1 = iota // 0
       a3 = 100 // 100
       a2 = iota // 2
    )
    
### 数据类型

| 序号  | 类型    | 描述                                                                                                        |
|-----|-------|-----------------------------------------------------------------------------------------------------------|
| 1   | 布尔型   | 布尔型的值只可以是常量true或false。`var a bool = true`                                                                 |
| 2   | 数字类型  | 整型int和浮点型float32, float64, Go语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码                                               |
| 3   | 字符串类型 | 字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本                                 |
| 4   | 派生类型  | 包括: (a)指针类型 (Pointer) (b)数组类型 (c)结构化类型 (struct) (d)Channel类型 (e)函数类型 (f)切片类型 (g)接口类型 (interface) (h)map类型 |

- #### 数字类型

| 序号  | 类型      | 描述                                                  |
|-----|---------|-----------------------------------------------------|
| 1   | uint8   | 无符号8位整型(0到255)                                      |
| 2   | uint16  | 无符号16位整型(0到65535)                                   |
| 3   | uint32  | 无符号32位整型(0到4294967295)                              |
| 4   | uint64  | 无符号64位整型(0到18446744073709221615                     |
| 5   | int8    | 有符号8位整型(-128到127)                                   |
| 6   | int16   | 有符号16位整型(-32768到-32767)                             |
| 7   | int32   | 有符号32位整型(-2147483648到-2147483647)                   |
| 8   | int64   | 有符号64位整型(-9223372036854775808到-9223372036854775807) |

- #### 浮点型

| 序号  | 类型         | 描述               |
|-----|------------|------------------|
| 1   | float32    | IEEE-754 32位浮点型数 |
| 2   | float64    | IEEE-754 64位浮点型数 |
| 3   | complex64  | 32位实数和虚数         |
| 4   | complex128 | 64位实数和虚数         |

- #### 其他数字类型

| 序号  | 类型          描述        |
|-----|:----------------------|
| 1   | byte类似uint8           |
| 2   | rune类似int32           |
| 3   | unit 32或64 位          |
| 4   | int与unit一样大小          |
| 5   | uintptr无符号整型，用于存放一个指针 |


### 布尔类型

- 不能使用0和非0表示true和false

### 数字类型

### 字符串类型

- #### 字符串声明
```go
package main

import "fmt"

func main() {
  str := `
  <h1>Hello World!</h1>
  `
  var str2 string = "Hello Go!"
  var str3 = "Hello Go"
  str4 := "Hello Go!"
  fmt.Println(str, str2, str3, str4)
}
```
- #### 字符串连接

```go
package main

import (
  "fmt"
  "strings"
  "bytes"
)

func main() {
  str := "tom"
  age := "20"
  fmt.Println(str + age)

  msg := fmt.Sprintf("%s%s", str, age)
  fmt.Println(msg)

  msg2 := strings.Join([]string{str, age}, ",")
  fmt.Println(msg2)

  var buffer bytes.Buffer
  buffer.WriteString("tom")
  buffer.WriteString(",")
  buffer.WriteString("20")
  fmt.Println(buffer.String())
}
```

- #### 字符串转义字符

| 转义符 | 含义        |
|-----|-----------|
| \r  | 回车符(返回行首) |
| \n  | 换行符       |
| \t  | 制表符       |
| \'  | 单引号       |
| \"  | 双引号       |
| \\  | 反斜杆       |

```go
package main

import "fmt"

func main() {
	str := "hello\n"
	fmt.Print(str, str)

	str2 := "hello\tworld"
	fmt.Print(str2)

	str3 := "C:\\user\\project"
	fmt.Print(str3)
}
```

- #### 字符串切片

```go
package main

import "fmt"

func main() {
	str := "HelloWorld"
	n := 2
	m := 5
	fmt.Printf("%c\n", str[n])   // 截取n的字符串
	fmt.Printf("%v\n", str[n:m]) // 截取n到m的字符串
	fmt.Printf("%v\n", str[n:])  // 截取n到最后的字符串
	fmt.Printf("%v\n", str[:n])  // 截取从0到n的字符串
}
```

- #### 字符串函数

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	fmt.Printf("%v\n", len(str))                        // 字符串长度
	fmt.Printf("%v\n", strings.Split(str, " "))         // 分割
	fmt.Printf("%v\n", strings.Contains(str, "hello"))  // 是否包含某个字符串
	fmt.Printf("%v\n", strings.ToLower(str))            // 小写转换
	fmt.Printf("%v\n", strings.ToUpper(str))            // 大写转换
	fmt.Printf("%v\n", strings.HasPrefix(str, "Hello")) // 是否以"Hello"开头
	fmt.Printf("%v\n", strings.HasSuffix(str, "World")) // 是否以"World"结尾
	fmt.Printf("%v\n", strings.Index(str, "He"))        // 查找索引位置
	fmt.Printf("%v\n", strings.LastIndex(str, "He"))    // 查找索引位置
}
```

### 格式化输出

- #### 普通占位符

| 占位符 | 说明         | 举例                          |
|-----|------------|-----------------------------|
| %v  | 相应值的默认格式。  | `fmt.Printf("%v", site)`    |
| %+v | 会添加字段名     | `fmt.Printf("%#v", site)`   |
| %#v | 相应值的Go语法表示 | `fmt.Printf("%#v", site)`   |
| %T  | 输出类型       | `	fmt.Printf("%T", site)`   |
| %%  | %输出        | `	fmt.Printf("%%")`         |

- #### 布尔占位符

| 占位符 | 说明         | 举例                       |
|-----|------------|--------------------------|
| %t  | true或false | `fmt.Printf("%t", true)` |

- #### 整型占位符

| 占位符 | 说明                             | 举例                         |
|-----|--------------------------------|----------------------------|
| %b  | 二进制表示                          | `fmt.Printf("%b", 5)`      |
| %c  | 相应的Unicode码点所表示的字符             | `fmt.Printf("%c", 0x4E2D)` |
| %d  | 十进制表示                          | `fmt.Printf("%d", 0x12)`   |
| %o  | 八进制表示                          | `fmt.Printf("%o", 10)`     |
| %q  | 单引号围绕的字符字面值                    | `fmt.Printf("%q", 13)`     |
| %x  | 十六进制表示，字母形式为小写a-f              | `fmt.Printf("%x", 13)`     |
| %X  | 十六进制表示，字母形式为大写A-F              | `fmt.Printf("%X", 13)`     |
| %U  | Unicode格式: U+1234,等同于 "U+%04X" | `fmt.Printf("%U", 0x4E2D)` |

### 运算符

- `++`(自增)和`--`(自减) 在Golang是单独的语句，不是运算符
```go
package main

import "fmt"

func main() {
	a := 10
	a--
	fmt.Println(a) // 9
	b := 10
	b++
	fmt.Println(b) // 11
}
```
- #### 关系运算符
```go
package main

import "fmt"

func main() {
	fmt.Println(100 == 100) // true
	fmt.Println(100 > 99) // true
	fmt.Println(99 < 100) // true
	fmt.Println(100 >= 100) // true
	fmt.Println(100 <= 101) // true
	fmt.Println(100 != 101) // true
}
```

- #### 逻辑运算符

| 运算符          | 描述                                      |
|--------------|-----------------------------------------|
| &&           | 逻辑AND运算符，如果两边的操作数都是true，则为true，否则为false |
| &#124;&#124; | 逻辑OR运算符，如果两边的操作数有一个true，则为true，否则为false |
| !            | 逻辑NOT运算符，如果条件为true，则为false，否则为true      |

```go
package main

import "fmt"

func main() {
	a := true
	b := false
	r := a && b
	fmt.Println(r) // false
	r2 := a || b
	fmt.Println(r2) // true
	r3 := !a
	fmt.Println(r3)	// false
}
```

- #### 位运算符

| 运算符      | 描述                                            |
|----------|-----------------------------------------------|
| &        | 参与运算的两数各对应的二进位相与 (两位均为1才为1)                   |
| &#124;   | 参与运算的两数各对应的二进位相或 (两位有一个为1就为1)                 |
| ^        | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1 (两位不一样则为1) |
| <<       | 左移n位就是乘以2的n次方. "a<<b"是把a的各二进位全部左移b位，高位丢弃，低位补0 |
| &gt;&gt; | 右移n位就是除以2的n次方，"a>>b"是把a的各二进位全部右移b位            |

```go
package main

import "fmt"

func main() {
	a := 4 // 二进制 0000 0100
	fmt.Printf("%b\n", a)
	b := 8 // 二进制 0000 1000
	fmt.Printf("%b\n", b)

	fmt.Printf("(a & b) : %v, %b\n", (a & b), (a & b))
	fmt.Printf("(a | b) : %v, %b\n", (a | b), (a | b))
	fmt.Printf("(a ^ b) : %v, %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << 2) : %v, %b\n", (a << 2), (a << 2))
	fmt.Printf("(b << 2) : %v, %b\n", (b << 2), (b << 2))
}
```

- #### 赋值运算符

| 运算符       | 描述            | 实例                            |
|-----------|---------------|-------------------------------|
| =         | 将表达式的值赋给另一个左值 | `c = a + b`                   |
| +=        | 相加后再赋值        | `c += a等于c = c + a`           |
| -=        | 相减后再赋值        | `c -= a等于c = c - a`           |
| *=        | 相乘后再赋值        | `c *= a等于c = c * a`           |
| /=        | 相除后再赋值        | `c /= a等于c = c / a`           |
| %=        | 求余后再赋值        | `c %= a等于c = c % a`           |
| <<=       | 左移后赋值         | `c <<= 2等于c = c << 2`         |
| &gt;&gt;= | 右移后赋值         | `c >>= 2等于c = c >> 2`         |
| &=        | 按位与后赋值        | `c &= 2等于c = c & 2`           |
| &#124;=   | 按位或后赋值        | `c &#124;= 2等于c = c &#124; 2` |
| ^=        | 按位异或后赋值       | `c ^= 2等于c = c ^ 2`           |

### 流程控制

```go
package main

import "fmt"

func main() {
	fmt.Println("aaa")
	fmt.Println("bbb")

	a := 100
	if a > 20 {
		fmt.Println(">")
	} else {
		fmt.Println("<")
	}

	switch a {
	case 100:
		fmt.Println(100)
	default:
		fmt.Println("default")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := [...]int{1, 2, 3} // 定义数组
	for i, v := range x {
		fmt.Println(v, i)

	}
}
```

- #### if语句
  - 初始变量可以声明在布尔表达式里面，注意作用域
  - 不能使用0或非0表示真或假
  - if 和 { 必须在同一行
  - { 必须在 if 或 else 同一行
  - ```go
    package main
  
    import "fmt"
  
    func main() {
        a := 1
        b := 2
        if a > b {
        fmt.Println(a)
        } else {
        fmt.Println(b)
        }
  
      flag1 := true
      if flag1 {
          fmt.Println(true)
      }
  
      if age := 20; age > 18 {
          fmt.Println("成年了", age)
      } else {
          fmt.Println("未成年", age)
      }
    }
    
  - ##### if嵌套
  - ```go
    package main
  
    import "fmt"
  
    func main() {
      a, b, c := 1, 2, 3
      if a > b {
          if a > c {
              fmt.Println("a最大")
          }
      } else {
          if b > c {
              fmt.Println("b最大")
          } else {
              fmt.Println("c最大")
          }
      }
    }
    
- #### switch语句

```go
package main

import "fmt"

func main() {
  // 单值条件
  a := 100
  switch a {
  case 10:
      fmt.Println(false)
  case 50:
      fmt.Println(false)
  case 100:
      fmt.Println(true)
  default:
      fmt.Println("Error")
  }
	
	
  // 多值条件
  day := 2
  switch day {
  case 1, 2, 3, 4, 5:
    fmt.Println("工作日")
  case 6, 7:
    fmt.Println("休息日")
  default:
    fmt.Println("非法输入")
  }
  
  // fallthrough穿透
  num := 100
  switch num {
  case 100:
    fmt.Println("100")
    fallthrough
  case 200:
    fmt.Println("200")
  case 300:
    fmt.Println("300")
  }
}
```

- #### for循环
  - Go语言中的for循环，只有for关键字，去除了其他语言的while和do while
  - ```go
    package main
    import "fmt"
  
    func f1() {
      // for在内的初始值
      for i := 0; i <= 10; i++ {
          fmt.Println(i)
      }
    }
  
    func f2() {
      // for在外的初始值
      i := 0
      for ; i <= 10; i++ {
          fmt.Println(i)
      }
    }
  
    func f3() {
      // for在外的初始值在内的改变条件
      i := 0
      for i <= 10 {
          fmt.Println(i)
          i++
      }
    }
  
    func f4() {
      // 永真循环
      i := 1
      for {
          i++
          fmt.Println(i)
      }
    }
  
    func main() {
      f1()
      f2()
      f3()
      f4()
    }

  - for range
```go
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5} // 数组
	for index, value := range a {
		fmt.Println("索引: ", index, "值: ", value)
	}

	// _ 匿名变量
	for _, value := range a {
		fmt.Println(value)
	}

	var b = []int{1, 2, 3} // 切片
	for _, v := range b {
		fmt.Println(v)
	}

	m := make(map[string]string, 0) // map
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "xxx@gmail.com"
	for key, value := range m {
		fmt.Println(key, value)
	}

	s := "hello"
	for _, i2 := range s {
		fmt.Printf("%c\n", i2)
	}
}
```

- #### `break`关键字
  - 语句可以结束`for`,`switch`,`select`的代码块
  - 单独在`select`中使用`break`和不使用`break`没有区别
  - 单独在表达式`switch`语句，并且没有`fallthrough`，使用`break`和不使用`break`没有区别
  - 单独在表达式`switch`语句，并且有`fallthrough`，使用`break`能够终止`fallthrough`后面的`case`语句的执行
  - 带标签的`break`，可以跳出多层`select/switch`作用域。让`break`更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带`break`的只能跳出当前语句块
```go
package main

import "fmt"

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func f2() {
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
		break
	case 2:
		fmt.Println(2)
		break
		fallthrough
	case 3:
		fmt.Println(3)

	}
}

func f3() {
MYLABEL:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			break MYLABEL
		}
	}
	fmt.Println("exit")
}
func main() {
	f1()
	f2()
	f3()
}
```

- #### `continue`关键字
  - `continue`只能用在循环中，在Go中只能用在`for`循环中
  - 终止本次循环，进行下一次循环
  - 在`continue`语句后添加标签时，表示开始标签对应的循环
```go
package main

import "fmt"

func f1() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v\n", i)
		} else {
			continue
		}
	}
}

func f2() {
	for i := 0; i < 10; i++ {
	MYLABEL:
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				continue MYLABEL
			}
			fmt.Printf("i:%v,j:%v\n", i, j)
		}
	}
}

func main() {
	f1()
	f2()
}
```

- #### `goto`关键字
  - `goto`语句通过标签进行代码间的无条件跳转
  - `goto`语句可以在快速跳出循环，避免重复退出上有一定的帮助
```go
package main

import "fmt"

func f1() {
	a := 0
	if a == 1 {
		fmt.Println(1)
	} else {
		goto END
	}
END:
	fmt.Println("exit")
}

func f2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				goto END
			}
			fmt.Println(i, j)
		}
	}
END:
	fmt.Println("exit")
}

func main() {
	f2()
}
```

### 数组

- #### 数组的定义
  - ```var variable_name [SIZE] variable_type```
  - `variable_name` 数组名称
  - `SIZE` 数组长度，必须是常量
  - `variable_type` 数组保存元素的类型
```go
package main

import "fmt"

func f1() {
	var a [3]int
	var b [5]string
	fmt.Printf("%T\n", a) // [3]int
	fmt.Printf("%T\n", b) // [5]string
	fmt.Printf("%v\n", a) // [0 0 0]
	fmt.Printf("%v\n", b) // [    ]
}

func main() {
	f1()
}
```

- #### 数组的初始化
  - 初始化就是给数组的元素赋值，没有初始化的数组，默认元素值都是零值，布尔类型是`false`，字符串是空字符串
```go
package main

import "fmt"

func f1() {
  // 初始化数组，固定长度
  var a = [2]int{1, 2}
  fmt.Printf("%v\n", a)
  var b = [2]string{"tom", "jack"}
  fmt.Printf("%v\n", b)
  var c = [2]bool{true, false}
  fmt.Printf("%v\n", c)
}

func f2() {
  // 初始化数组，长度不固定
  var a = [...]int{1, 2, 3, 4, 5}
  fmt.Printf("%v\n", a)
  var b = [...]string{"tome", "jack", "som"}
  fmt.Printf("%v\n", b)
}

func f3() {
  // 初始化数组，索引值初始化
  var a = [...]int{0: 1, 1: 2, 2: 5, 3: 4}
  fmt.Printf("%v\n", a)
  fmt.Printf("%v\n", a[3])
}

func main() {
  f1()
  f2()
  f3()
}
```

- #### 访问数组元素

```go
package main

import "fmt"

func f1() {
	var a [2]int
	// 添加元素
	a[0] = 100
	a[1] = 200
	fmt.Println(a) // [100 200]

	// 修改元素
	a[0] = 300
	a[1] = 500
	fmt.Println(a) // [300 500]
}

func f2() {
	// 数组长度不能越界
	//fmt.Println(a[3])

	// 数组长度
	var b = [3]int{1, 2, 3}
	fmt.Println(len(b)) // 3

	var c = [...]int{1, 3, 2, 3, 3}
	fmt.Println(len(c)) // 5
}

func f3() {
	// 数组for遍历
	a := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println("索引:", i, "值:", a[i])
	}
	fmt.Println()
	// 数组for range遍历
	for index, value := range a {
		fmt.Println("索引:", index, "值:", value)
	}
}
func main() {
	f1()
	f2()
	f3()
}
```

- #### 数组的切分

```go
package main

import "fmt"

func main() {
  var a = [...]int{1, 2, 3, 4, 5, 6}
  b := a[0:3]
  fmt.Println(b) // [1 2 3]
  c := a[3:]
  fmt.Println(c) // [4 5 6]
  d := a[1:3]
  fmt.Println(d) // [2 3]
  e := a[:]
  fmt.Println(e) // [1 2 3 4 5 6]
}
```

### 切片

- #### 定义切片
  - `var identifier []type`
  - 切片是引用类型
```go
package main

import "fmt"

func main() {
	// 直接声明
	var a []int
	var b []string
	fmt.Println(a, b)

	// make声明
	var c = make([]int, 0)
	fmt.Println(c)
}
```
- #### 切片的长度和容量

```go
package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(len(a)) // 长度
	fmt.Println(cap(a)) // 容量

	// 切片可以越界 但会抛出错误
	fmt.Println(a[9])
	fmt.Println(a[2])
}
```

- #### 切片的初始化

```go
package main

import "fmt"

func main() {
	// 数组初始化切片
	arr := [...]int{1, 2, 3}
	a := arr[:] // : 所有数组元素
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// 数组部分初始化切片
	arr2 := [...]int{1, 2, 3}
	b := arr2[0:2]
	fmt.Println(b)
	c := arr2[2:]
	fmt.Println(c)
	d := arr2[:2]
	fmt.Println(d)
}
```

- #### 切片的切分

```go
package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6}
	b := a[0:3]
	fmt.Println(b) // [1 2 3]
	c := a[3:]
	fmt.Println(c) // [4 5 6]
	d := a[1:3]
	fmt.Println(d) // [2 3]
	e := a[:]
	fmt.Println(e) // [1 2 3 4 5 6]
}
```

- #### 空(nil)切片

```go
package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a == nil) // true
	fmt.Println("长度", len(a), "容量", cap(a)) // 0 0
}
```

- #### 切片的遍历

```go
package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}

	// for遍历
	for i := 0; i < len(a); i++ {
		fmt.Printf("索引 : %v,值 :  %v\n", i, a[i])
	}

	// for range遍历
	for index, value := range a {
		fmt.Printf("index: %v value: %v\n", index, value)
	}
}
```

- #### 切片的增删改查和复制

```go
package main

import "fmt"

func add() {
	// 添加
	var a = []int{}
	a = append(a, 100)
	a = append(a, 200)
	fmt.Println(a)
}

func del() {
	// 语法：a = append(a[:index], a[index+1:]...)
	// 删除
	var a = []int{1, 2, 3, 4}
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
}

func upd() {
	// 修改
	var a = []int{1, 2, 3, 4}
	a[1] = 100
	fmt.Println(a)
}

func sel() {
	// 查询
	var a = []int{1, 2, 3, 4}
	for i, i2 := range a {
		if i2 == 2 {
			fmt.Printf("a[%v]\n", i)
		}
	}
}

func cpy() {
	// 深复制
	var a = []int{1, 2, 3, 4}
	var b = a
	b[0] = 100
	fmt.Println(a) // [100 2 3 4]
	fmt.Println(b) // [100 2 3 4]
	// 浅复制
	var c = []int{1, 2, 3, 4}
	var d = make([]int, 4)
	copy(d, c)
	d[0] = 100
	fmt.Println(c) // [1 2 3 4]
	fmt.Println(d) // [100 2 3 4]
}

func main() {
	add()
	del()
	upd()
	sel()
	cpy()
}
```

### map
  - map是一种`key:value`键值对的数据结构容器。map内部实现是哈希表`hash`
  - map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值
  - map是引用类型
    - map语法格式
    - ```go
      // 声明变量，默认 map 是nil
      var map_variable map[key_data_type]value_data_type
      // 使用 make 函数
      map_variable = make(map[key_data_type]value_data_type)
    - `map_variable` map变量名称
    - `key_data_type` key的数据类型
    - `value_data_type` 值的数据类型

- #### map声明和初始化

```go
package main

import "fmt"

func main() {
	// 两种声明和初始化map
	var mapData = map[string]string{
		"name": "tom",
		"age":  "18",
	}
	fmt.Println(mapData) // map[age:18 name:tom]

	var mapData2 = make(map[string]string)
	mapData2["name"] = "tom"
	mapData2["age"] = "20"
	fmt.Println(mapData2) // map[age:20 name:tom]

	fmt.Printf("%T %T\n", mapData, mapData2)

	// 判断key是否存在
	var mapData3 = map[string]string{
		"name": "tom",
	}
	v, ok := mapData3["name"]
	fmt.Println(mapData3, v, ok) // map[name:tom] tom true
}
```

- #### 遍历map

```go
package main

import "fmt"

func main() {
	var data = map[string]string{
		"name": "tom",
		"age":  "22",
	}
	for key, value := range data {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
```

### 函数

- #### Go的函数特性
  - Go有三种函数:普通函数，匿名函数(沒有名称的函数)，方法(定义在`struct`上的函数)
  - Go中不允许函数重载(overload), 也就是不允许函数同名
  - Go中的函数不能嵌套函数，但可以嵌套匿名函数
  - 函数是一个值，可以将函数赋值给变量，使这个变量也成为函数
  - 函数可以作为参数传递给另一个函数
  - 函数的返回值可以是一个函数
  - 函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数
  - 函数参数可以没有名称

- #### 函数的定义
  - ```go
    func function_name( [parameter list] ) [return_types]{
      // 函数体
    }
  - `func`  函数由`func`声明
  - `function_name` 函数名称，函数名和参数列表一起构成了函数签名
  - `[parameter list]` 参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际函数。参数列表指定的是参数类型，顺序及参数个数。参数是可选的，也就是说函数也可以不包含参数
  - `return_types` 返回类型，函数返回一列值。`return_types` 是该列值的数据类型。有些功能不需要返回函数值，这种情况下`return_types`不是必须的
  - 函数体: 函数定义的代码集合
```go
package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func comp(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	r := sum(1, 2)
	fmt.Println(r) // 3
	
	r2 := comp(3, 5)
	fmt.Println(r2) // 5
}
```

- #### 函数的返回值

```go
package main

import "fmt"

func f1() {
	fmt.Println("无参数和返回值函数")
}

func f2(a int, b int) {
	fmt.Println("有参数无返回值", a, b)
}

// 有参数有返回值
func f3(a int, b int) int {
	return a + b
}

// 有参数有返回并且有返回值变量
func f4(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 无参数有多个返回值
func f5() (name string, age int) {
	name = "tom"
	age = 20
	//return name, age
	return
}

func main() {
	f1()

	f2(1, 2)

	a := f3(1, 2)
	fmt.Println(a)

	b := f4(1, 2)
	fmt.Println(b)

	name, age := f5()
	fmt.Println(name, age)
}
```

- #### 函数的参数
  - 可以是0个或多个参数，参数需要指定数据类型
  - 声明函数时的参数列表叫做形参，调用时传递的参数叫做实参
  - Go是通过传值的方式传参的，意味着传递给函数的是拷贝后的副本，所以函数的内部访问，修改的也是这个副本
  - Go可以使用变长参数，有时候并不能确定参数的个数，可以使用变长参数，可以在函数定义语句的参数部分使用`ARGS...TYPE`的方式。这时会将`...`代表的参数全部保存到一个名为ARGS的slice中，注意这些参数的数据类型都是TYPE
```go
package main

import "fmt"

// 参数变长
func f1(args ...int) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}
// 固定参数和参数变长
func f2(name string, age int, args ...int) {
	fmt.Println(name, age, args)
}

// 验证浅拷贝和深拷贝
// 浅拷贝会生成副本不会影响原来的实参参数值
// 深拷贝会拷贝值的指针，会影响双向数据
// map slice interface channel

// 浅拷贝
func f3(a int) {
	a = 100
	fmt.Println(a) // 100
}

// 深拷贝
func f4(a []int) {
	a[0] = 10000
	fmt.Println(a) // [10000 2 3]
}

func main() {
	f1(1, 2, 3, 4, 5)
	f2("tom", 20, 123, 12313212, 123123123)

	test := 10
	f3(test)
	fmt.Println(test) // 10

	test2 := []int{1, 2, 3}
	f4(test2)
	fmt.Println(test2) // [10000 2 3]
}
```

- #### 函数类型与函数变量
  - 可以使用`type`关键字来定义一个函数类型
  - ```go
    type fun func(int, int) int

```go
package main

import "fmt"

// 函数类型
type fun func(int, int) int

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var ff fun
	
	ff = sum
	r := ff(1, 2)
	fmt.Println(r)

	ff = max
	r2 := ff(1, 2)
	fmt.Println(r2)
}
```

- #### 高阶函数

1. 函数参数

```go
package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello %s", name)
}
func test(name string, f func(string)) {
	f(name)
}

func main() {
	test("tom", sayHello)
}
```

2. 函数返回值

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	a := cal("+")
	fmt.Println(a(1, 2))
	
	s := cal("-")
	fmt.Println(s(2, 1))
}
```

3. 匿名函数

```go
package main

import "fmt"

func test() {
	name := "tom"
	age := "20"

	f := func() string {
		return name + age
	}
	msg := f()
	fmt.Println(msg)
}

func main() {
	// 匿名函数
	add := func() string {
		return "Hello World"
	}
	a := add()
	fmt.Println(a)

	// 自调用函数
	b := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(b)

	test()
}
```

- #### 闭包
  - 定义在一个函数内部的函数
  - 本质上是将函数内部和函数外包连接起来的桥梁或是函数和其引用环境的组合体
  - `闭包=函数+引用环境`
```go
package main

import "fmt"

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	var f = add()
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3

	f2 := add()
	fmt.Println(f2(40)) // 40
	fmt.Println(f2(50)) // 90
}
```

```go
package main

import "fmt"

// 闭包理解

func sayHello(name string) func(string) string {
	fmt.Println(name)
	return func(n string) string {
		name = n + name
		return name
	}
}

func sayHello2() func(string) string {
	var name = "Hello,"
	return func(n string) string {
		name = name + n
		return name
	}
}

func main() {
	str := sayHello("Tom")
	fmt.Println(str)

	fmt.Println("---------")

	str2 := sayHello2()
	s := str2("tom")
	fmt.Println(s)
}
```

- #### 递归
  - 函数内部调用函数自身的函数称为递归
  - 必须先定义函数的退出条件，没有退出条件，递归将成为死循环
  - 递归函数很可能会产生一大堆的goroutine，也可能会出现栈空间内存溢出的问题
```go
package main

import "fmt"

func f1() {
	// 阶乘
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Println(s)
}

func f2(n int) int {
	// 阶乘
	if n == 1 {
		// 退出条件
		return 1
	} else {
		// 自调用
		return n * f2(n-1)
	}
}

func f3(n int) int {
	// 斐波那契数列
	// f(n)=f(n-1)+f(n-2)且f(2)=f(1)=1
	if n == 1 || n == 2 {
		return 1
	}
	return f3(n-1) + f3(n-2)
}

func main() {
	f1()
	fmt.Println(f2(10))
	fmt.Println(f3(10))
}
```

### defer语句
  - Go中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行
  - #### defer特性
    - 关键字`defer`用于注册延迟调用
    - 这些调用直到`return`前才被执。因此，可以用来做资源清理
    - 多个`defer`语句，按先进后出的方式执行
    - `defer`语句中的变量，在`defer`声明时就决定了 
  - #### defer用途
    - 关闭文件句柄
    - 锁资源释放
    - 数据库连接释放
```go
package main

import "fmt"

func main() {
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	fmt.Println("d")
	// a b c d

	fmt.Println("-----")

	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("d")
	// a d c b
}
```

### init函数

  - Go有个特殊的函数`init`函数，先于`main`函数执行，实现包级别的一些初始化操作
  - #### init函数的主要特点
    - init函数先于`main`函数自动执行，不能被其他函数调用
    - init函数没有输入参数，返回值
    - 每个包可以有多个init函数
    - 包的每个源文件也可以有多个init函数
    - 同一个包的init执行顺序，Go没有明确定义
    - 不同包的init函数按照包导入的依赖关系决定执行顺序
  - 初始化顺序 : 变量的初始化->init()->main()
```go
package main

import "fmt"

var i int = initVar()

func main() {
	fmt.Println("main")
}

func initVar() int {
	fmt.Println("initVar")
	return 100
}

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
```

### 指针
  - Go中的函数传参都是值拷贝，当我们想要修改某个变量的时候，可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据
  - 类型指针不能进行偏移和运算
  - Go中的指针操作非常简单，只需要记住两个符号：`&`(取地址) 和`*`(根据地址取值)
  - #### 指针地址和指针类型
    - 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go中使用`&`字符放在变量前面对变量进行取地址操作。Go中的值类型`(int,float,bool,string,array,struct)`都有对应的指针类型,如`*int,*int64,*string`
  - #### 指针语法
    - 一个变量指向了一个值的内存地址。(也就是我们声明了一个指针后，可以像变量赋值一样，把一个值的内存地址放入到指针当中)
    - 类似于变量和常量，在使用指针前你需要声明指针
    - ```go
      var var_name *var-type
    - `var-type` 指针类型
    - `var_name` 指针变量名
    - `*` 指定变量是作为一个指针
```go
package main

import "fmt"

func main() {
	var ip *int
	fmt.Println(ip)        // <nil>
	fmt.Printf("%T\n", ip) // *int

	var i int = 100
	ip = &i
	fmt.Println(ip)  // 0xc00001a0e0
	fmt.Println(*ip) // 100

	var sp *string
	var s string = "hello"
	sp = &s
	fmt.Printf("%T\n", sp) // *string
	fmt.Println(sp)        // 0xc00004e250
	fmt.Println(*sp)       // hello

	var bl *bool
	var b bool = true
	bl = &b
	fmt.Printf("%T\n", bl) // *bool
	fmt.Println(bl)        // 0xc0000a60dc
	fmt.Println(*bl)       // true
}
```

- #### 指针数组
  - ```go
    var ptr [MAX]*int // 表示数组里面的元素的类型是指针类型
```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var ptr [3]*int
	fmt.Println(ptr) // [<nil> <nil> <nil>]

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i] // 值赋值给指针数组
	}
	fmt.Println(ptr) // [0xc000010138 0xc000010140 0xc000010148]

	for i := 0; i < len(ptr); i++ {
		fmt.Println(*ptr[i]) // 1 2 3
	}
}
```

### 类型定义和类型别名
  - #### 区别
    - 类型定义相当于定义了一个全新的类型，与之前的类型不同；但是类型别名并没有定义一个全新的类型，而是使用一个别名来替换之前的类型
    - 类型别名只会在代码中存在，在编译完成之后并不会存在该别名
    - 因为类型别名和原来的类型是一致的，所以原来的类型所拥有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前的任何方法
  - 类型定义语法 : ```type NewType Type```
```go
package main

import "fmt"

func main() {
	type MyInt int
	var i MyInt
	i = 100
	fmt.Println(i)  // 100
	fmt.Printf("%T", i) // main.MyInt
}
```
  - 类型别名语法 : `type NewType = Type`
```go
package main

import "fmt"

func main() {
	type MyInt = int
	var i MyInt
	i = 100
	fmt.Println(i)  // 100
	fmt.Printf("%T", i) // int
}
```

### 结构体
  - Go中没有面向对象的概念了，但是可以使用结构体来实现，面向对象编程的一些特性，例如：继承，组合等
- #### 结构体的定义
  - 结构体与类型定义类似，多了一个`struct`关键字
  - ```go
    type struct_variable_type struct {
      member definition;
      member definition;
      ...
      member definition;
    }
  - `type` 结构体类型定义关键字
  - `struct_variable_type` 结构体类型名称
  - `struct` 结构体关键字
  - `member definition` 成员定义
```go
package main

import "fmt"

func main() {
	// 声明结构体
	type Person struct {
		id   int
		age  int
		name string
	}

	// 声明结构体变量
	var tom Person
	fmt.Printf("%T, %v\n", tom, tom)
	jack := Person{}
	fmt.Printf("%T, %v\n", jack, jack)

	// 结构体赋值-1
	lave := Person{
		id:   1,
		age:  10,
		name: "lave",
	}
	fmt.Printf("%v", lave)
	fmt.Println(lave.name)

	// 结构体赋值-2
	leso := Person{}
	leso.id = 1
	leso.age = 20
	leso.name = "leso"
	fmt.Printf("%v", leso)
	fmt.Println(leso.name)

	// 匿名结构体
	var kono struct {
		id   int
		age  int
		name string
	}
	kono.id = 10
	kono.age = 20
	kono.name = "kono"
	fmt.Println(kono)
}
```

- #### 结构体的初始化
  - 未初始化的结构体成员都是零值
  - `int=0` `float=0.0` `bool=false` `string=nil`
```go
package main

import "fmt"

func main() {
	type Person struct {
		id      int
		name    string
		isAdmin bool
		price   float64
	}
	var tom Person
	fmt.Println(tom)

	// 键值对初始化
	a := Person{
		id:      1,
		name:    "a",
		isAdmin: true,
		price:   2.1,
	}
	fmt.Println(a)

	// 列表初始化
	b := Person{
		1,
		"b",
		false,
		2.2,
	}
	fmt.Println(b)

	// 部分初始化
	c := Person{
		id:    2,
		name:  "c",
		price: 2.3,
	}
	fmt.Println(c)
}
```

- #### 结构体指针
```go
package main

import "fmt"

func f1() {
	// 声明式创建结构体指针
	type Person struct {
		id   int
		name string
		age  int
	}
	tom := Person{
		id:   101,
		name: "tom",
		age:  10,
	}
	var p_person *Person
	p_person = &tom
	fmt.Println(p_person)  // &{101 tom 10}
	fmt.Println(tom)       // {101 tom 10}
	fmt.Println(&p_person) // 0xc00000a028
}

func f2() {
	// 通过new关键字创建结构体指针
	type Person struct {
		id   int
		name string
		age  int
	}
	tom := new(Person)
	fmt.Println(&tom) // 0xc00000a038
	fmt.Println(*tom) // {0 0}
}

func main() {
	f1()
	f2()
}
```

- #### 函数的结构体参数
  - 结构体可以像普通变量一样，作为函数的参数
  - 直接传递结构体，这是一个副本(拷贝)，在函数内部不会改变外面结构体内容
  - 传递结构体指针，这时在函数内部，能够改变外部结构体内容
```go
package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per Person) {
	per.id = 100
	fmt.Println("per:", per)
}

func showPerson2(per *Person) {
	per.id = 101
	fmt.Println(per)
}

func main() {
	// 值传递
	tom := Person{id: 1, name: "tom"}
	fmt.Println("tom:", tom)
	showPerson(tom)
	fmt.Println("tom:", tom)

	fmt.Println("----")

	// 地址拷贝
	jack := Person{
		id:   1,
		name: "jack",
	}
	fmt.Println(jack)
	per := &jack
	showPerson2(per)
	fmt.Println(*per)
}
```

- #### 结构体嵌套
  - Go中没有面向对象编程思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果
```go
package main

import "fmt"

func main() {
	type Dog struct {
		name  string
		age   int
		color string
	}
	type Person struct {
		dog  Dog
		name string
		age  int
	}
	dog := Dog{
		name:  "losa",
		age:   2,
		color: "black",
	}
	per := Person{
		name: "Tom",
		age:  20,
		dog:  dog,
	}
	fmt.Println(per)
	fmt.Println(per.dog)
	fmt.Println(per.dog.name)
}
```

### 方法
  - Go没有面向对象的特性，也没有类的概念，但是可以通过结构体来模拟这些特性，也可以声明一些方法，属于某个结构体
  - Go中的方法，是一种特殊的函数，定义于`struct`之上(与`struct`关联，绑定)，被称为`struct`的接收者(receiver)
  - 方法就是有接受者的函数
  - 语法格式
  - ```go
    type mytype struct {}
    func (recv mytype) my_method(para) return_type {}
    func (recv *mytype) my_method(para) return_type {}
  - `mytype` 结构体名称
  - `recv` 接收该方法的结构体(receiver)
  - `my_method` 方法名称
  - `para` 参数列表
  - `return_type` 返回值类型
```go
package main

import "fmt"

type Person struct {
	name string
}

func (receiver Person) eat() {
	fmt.Println("eat:", receiver)
}

func (receiver Person) sleep() {
	fmt.Println("sleep:", receiver)
}

type Customer struct {
	name string
}

func (c Customer) login(name string, pwd string) bool {
	fmt.Println(c.name)
	if name == "tom" && pwd == "123" {
		return true
	}
	return false
}

func main() {
	per := Person{
		name: "tom",
	}
	per.sleep()
	per.eat()

	cus := Customer{
		name: "tom",
	}
	resp := cus.login("tom", "123")
	fmt.Println(resp)
}
```

- #### 方法接收者指针类型
```go
package main

import "fmt"

type Person struct {
	name string
}

func f1() {

	p := Person{
		name: "tom",
	}
	p2 := &Person{
		name: "tom",
	}
	fmt.Println(p, p2)
	fmt.Printf("%T %T\n", p, p2)
}

func (receiver Person) showPerson1() {
	receiver.name = "tom..."
}
func (receiver *Person) showPerson2() {
	receiver.name = "tom..."
}
func main() {
	//f1()
	p1 := Person{
		name: "tom",
	}
	p2 := Person{
		name: "tom",
	}
	p1.showPerson1()
	fmt.Println(p1)
	fmt.Println("-----")
	p2.showPerson2()
	fmt.Println(p2)
}
```

- #### 方法结构体和嵌套结构体的实例
```go
package main

import "fmt"

type Box struct {
	title   string
	url     string
	setting Setting
}

type Setting struct {
	account  string
	password string
	token    string
	info     Info
}

type Info struct {
	avatar string
	name   string
	age    int
}

func (receiver *Setting) Login(account string, password string) {
	if account == receiver.account && password == receiver.password {
		setToken(receiver)
	}
}

func setToken(setting *Setting) {
	setting.token = "token-true"
}

func main() {
	box := Box{
		title: "test",
		url:   "127.0.0.1",
		setting: Setting{
			account:  "admin",
			password: "admin",
		},
	}

	box.setting.Login("admin", "admin")
	if box.setting.token == "token-true" {
		box.setting.info.avatar = "https://"
		box.setting.info.age = 20
		box.setting.info.name = "tom"
	} else {
		fmt.Println("login error")
	}

	fmt.Println(box)
	fmt.Println(box.setting)
	fmt.Println(box.setting.info)
}
```

### 接口
  - 接口类似公司的领导，他会定义一些通用规范，只设计规范，而不实现规范
  - Go中的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
```go
package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (receiver Computer) read() {
	fmt.Printf("%v\n", receiver.name)
	fmt.Print("read...\n")
}
func (receiver Computer) write() {
	fmt.Printf("%v\n", receiver.name)
	fmt.Print("write...\n")
}

func (m Mobile) read() {
	fmt.Printf("%v\n", m.model)
	fmt.Print("read...\n")
}

func (m Mobile) write() {
	fmt.Printf("%v\n", m.model)
	fmt.Print("write...\n")
}

func main() {
	var c USB
	c = Computer{name: "tom"}
	c.read()
	c.write()

	fmt.Print("---------\n")

	var m USB
	m = Mobile{
		model: "RedMi",
	}
	m.read()
	m.write()
}
```

- #### 接口值类型接收者和指针类型接收者
  - 本质上和方法的值类型接收者和指针类型接收者的思考方式是一样的
```go
package main

import "fmt"

type Pet interface {
	eat(string) string
	read(string) string
}

type Dog struct {
	name string
}

func (receiver Dog) eat(name string) string {
	receiver.name = name
	fmt.Println(name)
	return "eat very good!"
}

func (receiver *Dog) read(name string) string {
	receiver.name = name
	fmt.Println(name)
	return "read ok"
}

func main() {
	var dog Pet = new(Dog)

	res := dog.eat("food")
	fmt.Println(res)

	res2 := dog.read("food")
	fmt.Println(res2)

	fmt.Println(dog)
}
```

- #### 接口和类型的关系
  - 一个类型可以实现多个接口
  - 多个类型可以实现同一个接口(多态)
```go
package main

import "fmt"

// 一个类型可以实现多个接口
type Music interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct{}

func (receiver Mobile) playMusic() {
	fmt.Println("play music")
}

func (receiver Mobile) playVideo() {
	fmt.Println("play video")
}

type Person interface {
	read()
}

// 多个类型可以实现同一个接口(多态)
type A struct{}
type B struct{}

func (receiver A) read() {
	fmt.Println("A")
}
func (receiver B) read() {
	fmt.Println("B")
}

func main() {
	var m Music = new(Mobile)
	m.playMusic()
	var v Video = new(Mobile)
	v.playVideo()

	var perA Person = new(A)
	var perB Person = new(B)
	perA.read()
	perB.read()
}
```

- #### 接口嵌套
  - 接口可以通过嵌套，创建新的接口
```go
package main

import "fmt"

type AA interface {
	readAA()
}

type BB interface {
	readBB()
}

type CC interface {
	AA
	BB
}

type ABC struct{}

func (receiver ABC) readAA() {
	fmt.Println("AAAA")
}

func (receiver ABC) readBB() {
	fmt.Println("BBBB")
}

func main() {
	var c CC = new(ABC)
	c.readAA()
	c.readBB()
}
```

- #### 通过接口实现OCP设计原则
  - 面向对象的可复用设计的第一块基石，便是所谓的“开-闭”原则(Open-Closed Principle，缩写为OCP)
```go
package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}
type Dog struct{}
type Cat struct{}

// Dog 实现Pet接口
func (receiver Dog) eat() {
	fmt.Println("dog eat")
}
func (receiver Dog) sleep() {
	fmt.Println("dog sleep")
}

// Cat 实现Pet接口
func (receiver Cat) eat() {
	fmt.Println("cat eat")
}
func (receiver Cat) sleep() {
	fmt.Println("cat sleep")
}

type Person struct{}

// pet 既可以传递Dog也可以传递Cat
func (receiver Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
func main() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	person.care(dog)
	person.care(cat)
}
```

- #### 使用OOP思想的属性和方法

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (receiver Person) eat() {
	fmt.Println("eat")
}
func (receiver Person) sleep() {
	fmt.Println("sleep")
}
func (receiver Person) work() {
	fmt.Println("work")
}
func main() {
	p := Person{
		name: "tom",
		age:  20,
	}
	p.eat()
	p.sleep()
	p.work()
}
```

### 继承
  - Go中是没有OOP的概念，也没有继承的概念，可以通过结构体嵌套实现这个特性
```go
package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (receiver Animal) eat() {
	fmt.Println("eat", receiver.name, receiver.age)
}

func (receiver Animal) sleep() {
	fmt.Println("sleep", receiver.name, receiver.age)
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func main() {
	d := Dog{
		Animal{
			name: "Losa",
			age:  2,
		},
	}
	c := Cat{
		Animal{
			name: "Cosa",
			age:  3,
		},
	}
	d.eat()
	d.sleep()

	c.eat()
	c.sleep()
}
```

### 构造函数
  - Go中没有构造函数的概念，但可以使用函数来模拟构造函数的功能
```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 不能小于0")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	per, err := NewPerson("tom", 0)
	if err == nil {
		fmt.Println(per)
	} else {
		fmt.Println(err)
	}
}
```

### 包
  - 包可以区分命名空间(一个文件夹中不能有两个同名文件)，也可以更好的管理项目。
  - Go创建一个包，一般是创建一个文件夹，在该文件夹里面的go文件中，使用package关键字声明包名称
  - 通常，文件夹名称和包名相同，并且同一个文件下面只有一个包

- #### 创建包
  - 创建一个名为service的文件夹
  - 创建一个service.go文件
  - 在该文件中声明包
  - ```go
    package service
    
    func test(){
      print("test")
    }

- #### 导入包
  - 要使用某个包下面的变量或者方法，需要导入该包，要导入从`GOPATH`开始的包路径，
  - ```go
    package main
    
    import "service"
    
    func main(){
      service.test()
    }
- #### 包管理工具
  - go modules 是 Golang 1.11新加的特性，用来管理模块中包的依赖关系
  - ##### 使用方法
    - 初始化模块
    - `go mod init <项目模块名称>`
    - 依赖关系处理，根据go.mod文件
    - `go mod tidy`
    - 将依赖包复制到项目下的vendor目录(如果包被屏蔽(墙),可以使用这个命令，然后再使用`go build -mod=vendor`编译)
    - `go mod vendor`
    - 显示依赖关系
    - `go list -m all`
    - 显示详细依赖关系
    - `go list -m -json all`
    - 下载依赖 [path@version] 是非必写的
    - `go mod download [path@version]`
```go
package main

import (
	"fmt"
	"github.com/FOS-Lover/Golang-Study-Notes/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	service.TestUserService()
	service.TestCustomerService()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

}
```

### 并发编程

- #### 协程
  - Go中的并发是函数相互独立运行的能力。Goroutines是并发运行的函数。
  - Go提供了Goroutines作为并发处理操作的一种方式
  - 创建一个协程非常简单，在任务函数前面加一个`go`关键字
```go
package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("test") // 启动一个协程来运行
	go showMsg("Golang")

	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end") // 主函数退出 程序就结束了

}
```

- #### 通道channel
  - Go中提供了一种为通道的机制，用于在goroutine之间共享数据。当您作为goroutine执行并发活动时，需要在goroutine之间共享资源或数据，通道充当goroutine之间的管道(通道)并提供了一种机制来保证同步交换
  - 需要在声明通道时指定数据类型。可以共享内置，命名，结构和引用类型的值和指针。数据在通道上传递：在任何给定时间只有一个goroutine可以访问数据项：因此按照设计不会发生数据竞争
  - 根据数据交换的行为，有两种类型的通道：无缓冲通道和缓冲通道。无缓冲通道用于执行goroutine之间的同步通信，而缓冲通道用于执行异步通信。无缓冲通道保证在发送和接收发送的瞬间执行两个goroutine之间的交换。缓冲通道没有这样的保证
  - 通道由make函数创建，该函数指定chan关键字和通道的元素类型
  - ##### 创建无缓冲和缓冲通道的语法
  - ```go
    Unbuffered := make(chan int) // 整型无缓冲通道
    buffered := make(chan int, 10) // 整型有缓冲通道
  - 使用内置函数`make`创建无缓冲和缓冲通道，`make`第一个参数需要关键字`chan`,然后是通道允许交换的数据类型
  - ##### 将值发送到通道的代码块需要`<-`运算符
  - ```go
    goroutine1 := make(chan string, 5)
    goroutine1 <- "test"
  - ##### 从通道接收值的代码块
  - ```go
    data := <-goroutine1
  - 通道发送和接收的特性
    - 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
    - 发送操作和接收操作中对元素值的处理都是不可分割的
    - 发送操作在完全完成之前会被阻塞。接收操作也是如此
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("send : ", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Println("receiver : ", value)
	fmt.Println("end...")
}
```

- #### WaitGroup实现同步
  - 查看添加`WaitGroup`和不添加`WaitGroup`的区别
```go
package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(i int) {
	//defer wp.Add(-1)
	defer wp.Done() // goroutine结束就登记-1
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wp.Add(1) // 启动一个goroutine就登记+1
	}
	wp.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("--end--")
}
```

- #### runtime包
  - runtime包里面定义了一些协程管理相关的API

- ##### runtime.Gosched()
  - 让出CPU时间片，重新等待安排任务
```go
  package main
  
  import (
      "fmt"
      "runtime"
  )
  
  func show(s string) {
      for i := 0; i < 2; i++ {
          fmt.Println(s)
      }
  }
  func main() {
      go show("java") // 子协程运行
      for i := 0; i < 2; i++ {
          runtime.Gosched() // 调度cpu给其他子协程
          fmt.Println("Golang")
      }
      fmt.Println("end...")
  }
```

- ##### runtime.Goexit()
  - 退出当前协程
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func main() {
	go show()
	time.Sleep(time.Second)
}
```
- ##### runtime.GOMAXPROCS()
  - 设置CPU核心数
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	fmt.Println(runtime.NumCPU()) // CPU核心数
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end...")
}
```

- #### Mutex互斥锁实现同步
  - 除了使用channel实现同步之外，还可以使用Mutex互斥锁的方式实现同步
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Println("i++", i)
	time.Sleep(time.Millisecond)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	defer wg.Done()
	i -= 1
	fmt.Println("i--", i)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Println("end i:", i)
}
```

- #### channel的遍历

```go
package main

import "fmt"

var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 3; i++ {
		r := <-c
		fmt.Println(r)
	}

	for i := range c {
		fmt.Println(i)
	}

	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	
}
```

- #### select
  - select是Go中的一个控制结构，类似于`switch`语句，用于处理异步IO操作。`select`会监听case语句中的channel的读写操作，当case中channel读写操作为非阻塞状态(既能读写)时，将会触发相应的动作
    - select中的case语句必须时一个channel操作
    - select中的default子句是可运行的
  - 如果有多个`case`都可以运行，`select`会随机公平地选出一个执行，其他不会执行
  - 如果没有可运行的`case`语句，且有`default`语句，那么就会执行`default`的动作
  - 如果没有可运行的`case`语句，且没有`default`语句，`select`将阻塞，直到某个`case`通信可以运行
```go
package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "Hello"
		defer close(chanInt)
		defer close(chanStr)
	}()
	for {
		select {
		case r := <-chanInt:
			fmt.Println("chanInt:", r)
		case r := <-chanStr:
			fmt.Println("chanStr:", r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
```

- #### Timer
  - Timer顾名思义，就是定时器的意思，可以实现一些定时操作，内部也是通过channel来实现的
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器 4种用法
	
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println(time.Now())
	t1 := <-timer1.C // 阻塞的，直到时间到了
	fmt.Println(t1)
	
	fmt.Println(time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println(time.Now())
	
	fmt.Println(time.Now())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now())
	
	
	fmt.Println(time.Now())
	<-time.After(time.Second * 2)
	fmt.Println(time.Now())
	

	// 停止定时器

	timerStop := time.NewTimer(time.Second)
	go func() {
		<-timerStop.C
		fmt.Println("func...")
	}()
	s := timerStop.Stop()
	if s {
		fmt.Println("stop...")
	}
	time.Sleep(time.Second)
	fmt.Println("main end...")

	// 重设定时器
	fmt.Println(time.Now())
	timerReset := time.NewTimer(time.Second * 5)
	timerReset.Reset(time.Second * 2)
	<-timerReset.C
	fmt.Println(time.Now())
}
```

- #### Ticker
  - Timer只执行一次，Ticker可以周期的执行
```go
package main

import (
	"fmt"
	"time"
)

func f1() {
	ticker := time.NewTicker(time.Second)

	count := 1
	for _ = range ticker.C {
		fmt.Println("ticker..")
		count++
		if count >= 5 {
			ticker.Stop()
			break
		}
	}
}

func f2() {
	ticker := time.NewTicker(time.Second)

	chanInt := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()
	sum := 0
	for i := range chanInt {
		fmt.Println("收到: ", i)
		sum += i
		if sum >= 10 {
			break
		}
	}
}

func main() {
	f1()
	f2()
}
```

- #### 原子变量的引入
```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}
func sub() {
	atomic.AddInt32(&i, -1)
}

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("exit...", i)
}
```

- #### 原子操作的详解
  - atomic提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作，善用atomic能够避免程序中出现大量的锁操作
  - atomic常见操作有
    - 增减
    - 载入 read
    - 比较并交换 cas
    - 交换
    - 存储 write
```go
package main

import (
  "fmt"
  "sync/atomic"
)

func test_add() {
  var i int32 = 100
  atomic.AddInt32(&i, 1)
  fmt.Println(i)	// 101
  atomic.AddInt32(&i, -1)
  fmt.Println(i)	// 100

  var j int64 = 200
  atomic.AddInt64(&j, 1)
  fmt.Println(j)	// 201
  atomic.AddInt64(&j, -1)
  fmt.Println(j)	// 200
}
func test_rw() {
  var i int32 = 100
  atomic.LoadInt32(&i) // read
  fmt.Println(i)	// 100
  atomic.StoreInt32(&i, 200) // write
  fmt.Println(i)	// 200
}

func test_cas() {
  var i int32 = 100
  b := atomic.CompareAndSwapInt32(&i, 100, 200)
  fmt.Println(b)	// true
  fmt.Println(i)	// 200
}

func main() {
  test_add()
  test_rw()
  test_cas()
}
```

### os模块
  - os标准库实现了平台(操作系统)无关的编程接口

- #### 文件目录相关
  - 增删改查写文件
  - 增删改目录

```go
package main

import (
	"fmt"
	"os"
)

func createFile() {
	// 创建文件
	file, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name())
	}
}

func createDir() {
	// 创建单个目录
	err := os.Mkdir("test2", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
}
func deleteFileOrDir() {
	// 删除文件
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 删除目录
	err2 := os.RemoveAll("test2")
	if err2 != nil {
		fmt.Println(err2)
	}
}
func getWd() {
	// 获取当前的工作目录
	dir, _ := os.Getwd()
	fmt.Println(dir)

	// 修改目录
	os.Chdir("d:/")
	dir, _ = os.Getwd()
	fmt.Println(dir)

	// 临时目录
	tmp := os.TempDir()
	fmt.Println(tmp)
}

func rename() {
	// 重命名
	err := os.Rename("text2.txt", "test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func readFile() {
	// 读取文件，返回字节数组
	bt, _ := os.ReadFile("test.txt")
	fmt.Println(string(bt[:]))
}
func writeFile() {
	// 写入文件
	str := "test"
	err := os.WriteFile("test.txt", []byte(str), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//createFile()
	//createDir()
	//deleteFileOrDir()
	//getWd()
	//rename()
	//readFile()
	writeFile()
}
```

- #### File文件读操作
```go
package main

import (
	"fmt"
	"os"
)

// 打开关闭文件
func openCloseFile() {
	// 只能读
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name())
	}
	errClose := file.Close()
	if errClose != nil {
		fmt.Println(err)
	} else {
		fmt.Println(errClose)
	}
	// 可以加权和加建等
	file2, err2 := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 755)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(file2)
	}
	errClose2 := file2.Close()
	if errClose2 != nil {
		fmt.Println(errClose2)
	}
}

// 创建文件
func createFile() {
	// 创建文件: 等价于os.OpenFile(name, os.O_RDWR|os.O_CREATE, 755)
	f, _ := os.Create("a.txt")
	fmt.Println(f.Name())
	// 创建临时文件: 第一个参数 目录默认，Temp 第二个参数 文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Println(f2.Name())
}

// 读取文件
func readFile() {
	// 固定字节读取
	/*
		f, _ := os.Open("test.txt")
		for {
			buf := make([]byte, 10)
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			} else {
				fmt.Println("n: ", n)
			}
			fmt.Println("Content:", string(buf))
		}
		f.Close()
	*/

	// 从偏移量读取
	/*
		f, _ := os.Open("test.txt")
		buf := make([]byte, 10)
		n, _ := f.ReadAt(buf, 3)
		fmt.Println(n)
		fmt.Println(string(buf))
	*/

	// 遍历读取目录
	/*
		de, _ := os.ReadDir("testDir")
		for _, v := range de {
			fmt.Println(v.IsDir()) // 判断是否为目录
			fmt.Println(v.Name())  // 目录名
		}
	*/
	
	// 定位
	f, _ := os.Open("test.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf))
	f.Close()
}
func main() {
	//openCloseFile()
	//createFile()
	readFile()
}
```

- #### File文件写操作

```go
package main

import "os"

func write() {
	// 字节写入
	// os.RDWR 读写  os.O_APPEND 追加  os.O_TRUNC 覆盖
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_TRUNC, 777)
	f.Write([]byte("test"))
	f.Close()
}

func writeString() {
	// 字符串写入
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 777)
	f.WriteString("Hello World")
	f.Close()
}

func writeAt() {
	// 字节偏数组写入
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 777)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}

func main() {
	//write()
	//writeString()
	writeAt()
}
```

- #### 进程相关操作