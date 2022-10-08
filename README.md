# Golang Study Note

- 采用 <a href="https://github.com/FOS-Lover/Golang-Study-Notes/blob/master/LICENSE">MIT</a> 协议
- 更新于 : 2022年10月8日

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