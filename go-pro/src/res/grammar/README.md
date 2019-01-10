# Go
Golang 语法总结


## 目录
 * [几个熟识的概念](#几个熟识的概念)
 * [程序结构](#程序结构)
 * [基本数据类型](#基本数据类型)
 * [函数](#函数)
 * [接口](#接口)
 * [Goroutines Channels](#goroutines-channels)
 * [并发与共享变量](#并发与共享变量)
 * [包](#包)
 * [测试](#测试)
 * [反射](#反射)
 * [unsafe包](#unsafe包)
 * [忠告](#忠告)
 * [Golang 开发中遇到的各种奇葩问题](https://github.com/AaronWharton/collections/blob/master/go-pro/src/res/grammar/QUESTION.md)


## 几个熟识的概念

- Go由包（package）组成。类似于其他语言中的库（libraries）和模块（modules）。包是由一个或多个.go源文件组成的，它定义了包的主要功能（例如IO），每个.go源文件又是由包定义（package ddeclaration）开始的。例如，package main表示文件所属的包，后面是其导入的其他包的列表，然后是存储在该文件中的程序声明，只有package名为main的包可以包含main函数，一个可执行程序有且仅有一个main包。

- `可见性规则`:go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包含调用。根据约定，函数名首字母小写即为private，函数名首字母大写即为public。

- Go语言一般不使用分号（若在一行中有好几个表达式才用分号；分隔开，例如for循环）。所有控制结构的左括号不都能放在下一行，否则会报语法错误。

- 若程序中缺少导包（import）或者导入了该程序不需要的包，则该程序编译无法通过。Go也不允许出现未被使用的局部变量。但若出现未被使用的常量（const）等，编译器不会报错。

- 输出函数`fmt.Println()`，文本和变量一起打印时注意格式：`fmt.Println("i=", i)`而不是`fmt.Println("i="+ i)`，这里和*Java*语法有不同。若之名输出格式：`fmt.Println("i=%d", i)`就不容易混淆（这里跟C语言的输出一样）。

- `range`：Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。for 循环中第一个参数表示键，第二个表示值。

- 函数闭包的使用，闭包是一个匿名函数值，会引用到其外部的变量。

- 在 go 语言中，len()函数可以用来获取 map 中 key 的个数。 `make` 可以为这三种类型进行内存的分配，同时也只能用在这三种类型上。 make 初始化了内部的数据结构，填充适当的值，返回初始化后的非零值， `new` 可以用于各种类型的内存分配， `new(TYPE)` 分配了零值填充的 TYPE 类型的内存空间，并返回其地址。即它返回的是一个内存地址 \*TYPE 类型的值（指针）。来看看官方文档对它们的解释：
	```go
	// The make built-in function allocates and initializes an object of type
	// slice, map, or chan (only). Like new, the first argument is a type, not a
	// value. Unlike new, make's return type is the same as the type of its
	// argument, not a pointer to it. The specification of the result depends on
	// the type:
	//	Slice: The size specifies the length. The capacity of the slice is
	//	equal to its length. A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length, so make([]int, 0, 10) allocates a slice of length 0 and
	//	capacity 10.
	//	Map: An empty map is allocated with enough space to hold the
	//	specified number of elements. The size may be omitted, in which case
	//	a small starting size is allocated.
	//	Channel: The channel's buffer is initialized with the specified
	//	buffer capacity. If zero, or the size is omitted, the channel is
	//	unbuffered.
	func make(Type, size IntegerType) Type

	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly
	// allocated zero value of that type.
	func new(Type) *Type
	```

-  if 的使用： if...else... 语句属于同一个语句块，所以在 if 里面定义的变量可以在 else 语句里面使用，但是不能在 if...else... 语句块之外使用其定义的变量。错误示范：
	```go
	if f, err := os.Open(fname); err != nil {	// compile error: unused: f
		fmt.Println(err)			// NOTE: f and err can only be used in this block
	}
	f.Close()	// compile error: undefined f
	```
	修改代码：
	```go
	if f, err := os.Open(fname); err != nil {
		fmt.Println(err)
	} else {
		// f and err are visible here too
		f.Close()
	}
	```
	一种简洁的写法：
	```go
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
	```
	在这里我们打开了一个文件，它只有打开成功和失败两种可能，为了简单我们直接省略`else`，在`if`外部直接关闭打开的文件（第一种可能出现的情况），在`if`内部直接执行打开失败的处理（第二种可能出现的情况），当遇到下面的情况我们也可以这样处理：
	```go
	if condition {
		return x
	}
	return y
	```
	这种适用于我们只用判断两个条件的情况下，功能和`else`相同，但是看起来更加简洁明了。

-  Go 里每种类型都有属于自己的**空值**！例如：

    -  int 空值是0
    -  string 空值是” ”而不是 nil （这与其他语言确实有一点区别）
    -  Slice 空值是长度为0的 Slice 而不是 nil
    -  map 空值是 nil
    -  error 空值是 nil
    -  struct 空值是一个“所有成员都是空值”的空 Struct 而不是 nil

    **所以**:
    - 很多时候，你可以直接 append() 一个元素到空 Slice 里，但是你却不能对空 map 操作，因为 map 的空值是 nil
    - 你不能粗鲁地判断一个 struct 是不是 nil ，因为它永远不可能是 nil ，所以你可以通过返回一个 error 来判断是否为空，就像很多 Go 标准库里的做法那样：

        ```go
        if err != nil
        ```


## 程序结构

- 25个关键字(`if/default/chan/func`...)不可作为 Name ，37个 predeclared names (`true/nil/panic`...)可作为 Name 。 Name 以大写字母开头作用于全局（所有人都可访问），小写字母作用于局部（意思是仅对包内公开）。

- `declaration` 不同于 `assignment` ，`:=` 是 declaration ，`=` 是 assignment 。

- 变量声明的形式：

  ```go
  var s string                        // 定义变量s，类型为string，默认赋值""(若为int型则为0)
  var s = ""                          // 定义变量s，赋值为""(因此默认s变量为string型)
  var s string = ""                   // 定义变量s，类型为string，并给它赋值""
  s := ""                             // 用于局部变量(`short variable declaration`)
  var b, f, s = true, 2.3, "four"     // 定义多个变量，未事先指定变量类型，右边的值隐式地显示变量类型
  i, j := 0, "255"                    // 同上。这里是定义的局部变量
  ```

- 包下面的公共常量以大写字母开头，每个包下面需要一个包说明文档(package doc comment)，通常命名为 `doc.go`。


## 基本数据类型

- `rune` 是int32 的别名。用UTF-8进行编码。这个类型在什么时候使用呢？例如需要遍历字符串中的字符，可以循环每个字节（仅在使用US ASCII 编码字符串时与字符等价，而它们在Go中不存在！）。因此为了获得实际的字符，需要使用rune 类型。在UTF-8 世界的字符有时被称作runes。通常，当人们讨论字符时，多数是指8 位字符。UTF-8 字符可能会有32 位，称作rune。

- `const`中提供了6种*untyped*类型（boolean integer rune floating-point complex string），表示常量在定义时没有被显示指明数据类型。*注意：只有常量有untyped类型。*变量在未明确指定数据类型时也*会被隐式地指定数据类型*。在`const`中定义常量时如果不指明数据类型时要注意一些细节，比如：`iota`（自增的特殊常量）自定义KiB、MiB、GiB、TiB、PiB、EiB、ZiB、YiB时，YiB高达1208925819614629174706176字节……无法存放在int型中，但是`YiB/ZiB`是合法的（结果为1024），详见P78~P79。 iota 的使用如下：
	```go
	const(
		a = 'A'
		b
		c
		d = iota
		e
	)

	func main()  {
		fmt.Println(a,b,c,d,e)	// 65 65 65 3 4
	}
	```
	由输出可以得到以下结论：
	- 定义常量组：如果不提供初始值，则使用上行表达式（所以b、c的值均为 A 的 ASCII 码值65）
	- 每遇到一个const关键字，iota就重置为0（所以d为3，e为4，因为前面a、b、c占了0、1、2的位置）
	- iota是常量计数器，从0起，组中每定义一个常量自动递增1
	- 可以用初始化规则和iota实现枚举

- `map` ：格式为 `map[int]string ==> Key-Value`。 其中 **Value 可以是任意类型，即 interface{} ，但是 Key 必须是可用 == 比较的**，这样在 map 中才可以测试给定的一个 Key 是否已经存在于 map 中（尽管浮点型数据可以进行比较，但是通常不这么做，避免出现 NaN 的情况）。
	```go
	// declaration variable m1 as map
	var m1 map[string]string
	// use make() to create a non-nil map, cause nil map can not be assigned
	m1 = make(map[string]string)
	// assignment to the map
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "cc"
	m1["d"] = "dd"
	m1["e"] = "ee"
	m1["f"] = "ff"
	m1["g"] = "gg"
	m1["h"] = "hh"
	m1["i"] = "ii"
	m1["j"] = "jj"
	m1["k"] = "kk"

	// create a non-nil map variable
	m2 := make(map[string]string)

	// assignment
	m2["a"] = "aa"
	m2["b"] = "bb"

	// declaration and assignment
	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
	}

	// print the whole map
	fmt.Println(m3)			// map[a:aa b:bb]

	// modify the value
	m3["b"] = "ssr"
	fmt.Println(m3)			// map[a:aa b:ssr c:cc]

	// search the value which key is "a"
	// v is value, exists is a bool which represents
	// whether v exists or not.
	v, exists := m3["a"]
	fmt.Println(v, exists)		// aa true

	// find out whether the key value exists
	if v, ok := m1["d"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")	// Key Not Found
	}

	// traversal
	// NOTE: The output is random!
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	```
	最后一段代码演示了 map 可以使用 range 来遍历，但是请注意： map 的遍历顺序是随机的。这意味着输出结果的顺序无法预测。如果想要按照某个顺序遍历输出，则需要自己来定制。

- `array`和`slice`在形式上的区别：
  array初始化
  ```go
  month := [...]int{1, 2, 3, 4, 5, 6}       // ... 省略数组的长度指定，由{}里面的数据个数决定
  ```
  slice初始化
  ```go
  month := []int(1, 2, 3, 4, 5, 6)
  ```

- array可以用`==`和`!=`判断，但是slice不能直接用`==`和`!=`进行判断，可以选择用循环来比较slice里面的每一个元素，进而判断slice是否相等。
	```go
	arr1 := [...]int{1,2,3,4,5}
	arr2 := [...]int{1,2,3,4,5}
	fmt.Println(arr1 == arr2)	// true

	arr := [...]int{6,6,6,6,6,6}
	s1 := arr[1:2]
	s2 := arr[2:3]
	fmt.Println(s1 == s2)		// Error: operator not defined on []int
	```
	比较两个 slice 是否相等可以使用以下两个方法：
	- reflect 进行深度比较（DeepEqualD）：
		```go
		fmt.Println(reflect.DeepEqual(s1, s2))	// true
		```
	- 循环遍历比较：
		```go
		func IsEqual(s1 []int, s2 []int) bool {
			if len(s1) != len(s2) {
				return false
			}

			// make sure []int{} != []int(nil)
			if (s1 == nil) != (s2 == nil) {
				return false
			}

			for k, v := range s1 {
				if v != s2[k] {
					return false
				}
			}
			return true
		}
		```
		一般使用循环遍历的方法，因为反射的效率低下。

- 数组长度不可变，但是slice的长度可变，slice是依赖数组的。slice三要素：**指针**、**长度** `len()` 、**容量** `cap()` ，slice可以超出长度，但是不可超出容量。slice的容量为提取数组的开始位置到数组结束位置，如果没有超过slice的cap值，则会返回原来的slice引用，从而影响到引用同一个数组的其他slice，如果超过了cap值，则会新建一个slice，并返回新slice的引用，此种状态会跟原来的slice脱离引用关系。下面代码说明长度、容量的关系：
	```go
	s := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := s[2:4]
	fmt.Println(a, ",", len(a),",", cap(a))		// [3 4] , 2 , 7
	```

- slice包含一个底层数组（underlying array）的指针，所以向函数传递一个slice并在函数里面修改slice可以改变底层数组的值：
	```go
	// This is an example for explaining the use of
	// slice of modifying the underlying array elements.
	func main() {
		s := [...]int{2, 3, 4, 5, 6, 7, 8, 9}
		test1(s[:])
		fmt.Println(s)		// [9 8 7 6 5 4 3 2]

		word := "This is a test."
		test2(word)
		fmt.Println(word)	// This is a test.
	}

	// Slice contains a pointer to an element of an array,
	// passing a slice to a function permits the function
	// to modify the underlying array elements.
	func test1(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	// Passing a string to a function just pass the copy
	// of the string, so it can not affect original
	// value of the string.
	func test2(s string) {
		s = "Modified."
	}
	```


## 函数

- `defer`关键字：用于资源的释放（也可和 recover 关键字搭配使用，用于从 panic 中恢复），会在函数返回之前进行调用。

	- defer 是在 return 之前执行的（这里的 return 相当于汇编中的 `ret` 指令，并不是 function 中的 `return` ）。

 	-  ```return xxx``` 这一条语句并不是一条原子指令（①给返回值赋值，②用return返回该值）！使用 defer 关键字时函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。**defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致**。~~defer转换步骤：①返回值 = xxx②调用defer函数③return（直接返回“返回值”）~~

 	- 由此可见， defer 的三个基本原则（更详细全面的讲解看参考资料第一篇）：
		- 当defer调用函数的时候, 函数用到的每个参数和变量的值也会被计算。
		- Defer调用的函数将在当前函数返回的时候, 以后进先出的顺序执行。
		- Defer调用的函数可以在返回语句执行后读取或修改命名的返回值。

 	- 参考资料：
		- [Defer, Panic, and Recover[翻译]](https://my.oschina.net/chai2010/blog/119216)
		- [使用Defer的几个场景](https://my.oschina.net/chai2010/blog/140065)


## 接口

-  `interface` 类型定义了一组方法，如果某个对象实现了某个接口的**所有方法**，则此对象就实现了该接口。
	```go
	// defines an interface
	type Men interface {
		SayHi()
		Sing(lyrics string)
	}

	// defines a struct
	type Human struct {
		name  string
		// ...
	}

	// implementing an  interface
	// Human implements SayHi() and Sing(lyrics string),
	// so Human implements Men.
	func (h *Human) SayHi() {
		fmt.Printf("Hello, I'm %s.", h.name)
	}

	func (h *Human) Sing(lyrics string) {
		fmt.Println("Sing a song...", lyrics)
	}

	// print the result
	s := Human{"Aaron"}
	s.SayHi()
	s.Sing("Lost stars")
	```

- `接口值`:由两部分组成，一个具体的类型和那个类型的值。**注意：一个包含nil指针的接口不是nil接口。**

- **Go的接口类型是静态类型**而不是动态类型：某个接口类型的变量的类型始终不变，即使在运行时其内部存储的接口变量在变换值，他们始终是该接口类型。（这个结论待验证！）

- `Type Assertion`：golang中的所有程序都实现了interface{}的接口，这意味着，所有的类型如string,int,int64甚至是自定义的struct类型都就此拥有了interface{}的接口，所以type assertion可以用来推测未知的数据类型。这种做法和java中的Object类型比较类似。那么在一个数据通过func funcName(interface{})的方式传进来的时候，也就意味着这个参数被自动的转为interface{}的类型。有些则需要断言（Type Assertion），如以下代码编译器将返回：*cannot convert a (type interface{}) to type string: need type assertion*：

    ```go
    func funcName(a interface{}) string {
        return string(a)
    }
    ```
    使用type assertion并结合switch语句判断数据类型：
    ```go
    package main

    import "fmt"

    func main() {
	    var str interface{} = "abc"

	    switch v := str.(type) {
	    case string:
		    fmt.Println(v + "\t" + "string")
	    case int32, int64:
		    fmt.Println(v)
	    default:
		    fmt.Println("unknown")
	    }
    }
    ```


## Goroutines channels

- `Channels` 用来同步并发执行的函数并提供它们某种传值交流的机制。
	```go
	ch  <- x	// a send statement; send variable x to the channel ch
	x = <-ch	// a receive statement in an assignment statement
	<-ch		// a receive statement; result is discarded
	```

- `缓冲的 channel `： ```ch := make(chan int, 1)``` 保证往缓冲中 send 数据先于对应的 receive 数据，简单说就是在取的时候里面肯定有数据，否则就因取不到而阻塞；在存数据的时候保证里面有空间，否则就因 channel 满了无法存数据而阻塞（即超过缓冲数目，例如这里指超过了1）。

- `非缓冲的 channel `： ```ch := make(chan int)``` 保证取数据先于存数据，就是保证 send 的时候肯定有其他的 goroutine 在 receive ，否则就因放不进去而阻塞，**把 unbuffered channel 当成一个装满的 channel ，需要先 receive ，才能 send 数据给 channel 。**举个例子：
	```go
	package main
	
	import (
		"fmt"
	)
	
	func f1(in chan int) {
		fmt.Println(<-in)
	}
	
	func main() {
		out := make(chan int)
		out <- 2		//Panic: unbuffered channel is full, cannot send data to it
		go f1(out)
	}
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
	// 
	// goroutine 1 [chan send]:
	// ...
	```
	修改：
	```go
	package main
	
	import (
		"fmt"
	)
	
	func f1(in chan int) {
		in <- 2
	}
	
	func main() {
		out := make(chan int)
		go f1(out)
		fmt.Println(<-out)		//OK: receive the data from f1()
	}
	// Output:
	// 2
	```
	印证了非缓冲 channel 的特点。
	**注意理解：在创建 channel 的协程（这里指 main ）中， unbuffered channel 需要在 main 中 receive ，在 f1 中 send 。**

- `select`是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。


## 并发与共享变量

- `竞争条件检测`：在 `go build` ， `go run` 或者 `go test` 命令后面加上 `-race` 的flag（`go run -race xx.go`）。不过它只能检测到运行时的竞争条件，并不能证明之后程序不会发生数据竞争。

- `竞争条件`（race condition）：一般指在多个 goroutine 的交叉操作下导致程序不能给出正确的执行结果。原文：Race condition is a situation in which the program does not give the correct result for some interleavings of the operations of multiple goroutines.

- `数据争用`（data race）：是一种特定的竞争条件。只要至少两个 goroutine 并发地访问同一个变量并且至少一个进行了写（write）的操作（比如赋值操作），就会发生数据争用。原文：Data race is a particular kind of race condition. A data race occurs whenever two goroutines access the same variable concurrently and at least one of the accesses is a write(such as assignment).

- 避免data race的三种途径：
	- 避免在多个goroutine中对同一个variables进行write操作（通常我们都会在多个goroutine中对同一个数据进行更新，所以这个方案一般用于那些读取的操作）
 	- 避免在多个goroutine中访问同一个variable（通常用于多个goroutine更新同一个数据，使用channel发送信号），在Go语言中推崇**不要使用共享数据来通信，而是使用通信来共享数据**。
 	- 允许多个goroutine都能访问该variable，但是一次只能一个进行访问（使用sync.Mutex互斥锁以及多读单写sync.RWMutex读写锁）

- 允许多个goroutine都能访问该variable，但是一次只能一个进行访问：使用 buffered channel 作为信号，通知goroutine可以使用该variable。官方利用这点封装了一个`sync.Mutex` type（`sync`包），使用`xx.Lock()`和`xx.Unlock()`进行mutual exclusion（互斥，即单一访问shared variable）操作。注意lock和unlock操作不能嵌套（mutex locks are not re-entrant）

- `sync.RWMutex`：multiple readers, single writer lock. RLock只能在临界区共享变量没有任何写入操作时可用。RWMutex需要更复杂的内部记录，所以会让它比一般的无竞争锁的mutex更慢一些，如果不知道该怎么使用RWMutex则可以考虑使用Mutex。

- `sync.Once`：懒加载，适用于初始化成本高但是又不是必须在程序执行时需要使用的数据：
	```go
	var loadIconsOnce sync.Once
	var icons map[string]image.Image
	// Concurrency-safe.
	func Icon(name string) image.Image {
		loadIconsOnce.Do(loadIcons)
		return icons[name]
	}
	```
	每一次对 Do(loadIcons) 的调用都会锁定 mutex ，并会检查 boolean 变量。在第一次调用的时候，变量的值是 false ， Do 回调用 loadIcons 并将   boolean 设置成 true ，这样随后的调用什么也不会做，但是 mutex 同步保证 loadIcon 对内存（这里指 icons 变量）产生的效果对所有的 goroutine可见 。使用 sync.Once 我们能够避免在变量被构建完成之前和其他 goroutine 共享该变量。

- `sync.WaitGroup` ：线程同步，用于等待一组或一系列动作执行完成后才继续执行下面的操作：
	```go
	func main() {
		s := []int {2, 5, 9, 98, 72, 152}
		var wait sync.WaitGroup
		for _, num := range s {
			wait.Add(1)
			go func(num int) {
				fmt.Println(num)
				wait.Done()
			}(num)	// NOTE: The parameter must not be ignored! Or the func will return the wrong result!
		}

		wait.Wait()
		fmt.Println("Over")
	}
	```
	输出结果如下：
	```
	2
	152
	5
	9
	98
	72
	Over
	```
	实际输出顺序可能不同，但是一定是六个数据全部输出完了才会输出 `Over` 字符。下面对 `sync.WaitGroup` 中三个重要的方法做说明：
	- `Add()` ：即 WaitGroup counter 。当参数为正时，表示正在执行“一系列动作”；当参数为负时，会发生 panic ；当参数为0时，所有阻塞在 `Wait` 上的 goroutine 都将得到释放，也就是说此时的“一系列动作”已经执行完毕
	- `Done()` ：“消费”一个“一系列动作”， decrements the WaitGroup counter
	- `Wait()` ：将会一直阻塞到 WaitGroup counter 变为0，然后解除阻塞状态，执行下面的代码

- **使用mutex的两方面考虑**：
	- 规范代码（例如在一堆goroutine中）的执行顺序。为那些不是原子操作的操作而设计，避免并发操作的时候发生竞争（在执行xxx操作“中”执行其他操作）。
	- （更重要的）内存同步。

- 内存同步：现代计算机中可能会有一堆处理器，每一个都会有其本地缓存（local cache）。为了效率，对内存的写入一般会在每一个处理器中缓存，并在必要时一起flush到主存。这种情况下这些数据可能会以与当初goroutine写入顺序不同的顺序提交到主存。考虑一下下面代码片段的可能输出：
	```go
	var x, y int

	s := make(chan bool)
	h := make(chan bool)

	go func() {
		x = 1	// A1
		fmt.Println("y:", y, " ")	// A2
		s <- true
	}()

	go func() {
		y = 1	// B1
		fmt.Println("x:", x, " ")	// B2
		h <- true
	}()

	<-s
	<-h
	```
	期望它的输出结果是这样：
	```
	y:0 x:1
	x:0 y:1
	x:1 y:1
	y:1 x:1
	```
	然而事实上下面的两种情况也是可能出现的：
	```
	x:0 y:0
	y:0 x:0
	```
	这两种情况在所使用的编译器，CPU或者其他因素下也是会发生的。当不使用mutex来进行内存同步，这两个goroutine可能对对方都是不可见的（两个goroutine在不同CPU上执行，每一个核心都有自己的缓存，这样一个goroutine的写入对于另外一个goroutine的Println，在主存同步之前都是不可见的了）。

- Goroutine 和线程的区别 ：
	- `动态栈`：每一个 OS 线程都有一个固定大小的内存块（通常是2 M ）来作栈，这个栈会用来存储当前正在被调用或挂起的函数的内部变量，而 goroutine 通常只有2 KB ，它的作用跟 OS 线程类似，但是它的大小不是固定的，而且通常很小（所以它可以同时创建成百上千个），当然很大的也有1 GB。
	- `Goroutine 调度`： OS 线程会被操作系统内核调度，每几毫秒，一个硬件计时器会中断处理器调用内核函数（类似于名为 scheduler ）挂起当前执行的线程并保存内存中它的寄存器内容，检查线程列表并决定接下来哪个线程可以被执行，并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。这一系列操作涉及到上下文的切换和内存访问（寄存器的访问），并且增加 CPU 的运行周期，这些操作执行起来会很慢。 `Go` 运行时包含了自己的调度器，调度器使用了一些技术手段比如 `m:n` ，表示在 n 个操作系统上调度 m 个 goroutine， Go 调度器的工作和内核的调度是相似的，只不过它只关注 goroutine 。 Go 的调度不涉及内核上下文，所以调度一个 goroutine 比调度一个线程代价要小很多。
	- `GOMAXPROCS` ：决定会有多少个操作系统的线程同时执行 Go 代码，默认值是运行机器上的 CPU 核心数（GOMAXPROCS 就是前面所说的 `m:n` 调度中的n）。在休眠中的或者在通信中被阻塞的 goroutine 是不需要一个对应的线程来做调度的。
	- 为什么 Goroutine 没有 ID 号：防止滥用线程 id 导致代码出现莫名其妙的错误。


## 包

- `相对路径导入`：
	```go
	import   "./model"  // 当前文件同一目录的model目录
	```

- `绝对路径导入`：
	```go
	import   "shorturl/model"  // 加载GOPATH/src/shorturl/model模块
	```

- `点操作`：
	```go
	import . "fmt"

	func main() {
		Println("Hello world")
	}
	```
	这种方式导包之后，在调用包里的函数时可以省略包名。

- `别名操作`：
	```go
	import f "fmt"

	func main() {
		f.Println("Hello world")
	}
	```

- `_操作`：
	```go
	import (
		"image/jpeg"
		_ "image/png"
	)
	```
	`_操作`其实只是引入该包。当导入一个包时，它所有的 `init()` 函数就会被执行，但有些时候并非真的需要使用这些包，仅仅是希望它的 `init()` 函数被执行而已。这个时候就可以使用`_操作`引用该包了。使用`_操作`引用包是无法通过包名来调用包中的导出函数，而是只是为了简单的调用其 `init函数()` （用意：往往这些 init 函数里面是注册自己包里面的引擎，让外部可以方便的使用，就很多实现 `database/sql` 的引擎，在 init 函数里面都是调用了`sql.Register(name string, driver driver.Driver)` 注册自己，然后外部就可以使用了）。


## 测试

- 测试代码：以 `_test.go` 结尾的文件。当使用 `go build` 命令执行 go 程序时不会 build 这些文件，只有使用 `go test` 命令才会 build 。在 `_test.go` 文件中，有三种函数被特别对待：

	- `test function` ：以 Test 开头命名的函数。该函数用来测试程序逻辑的正确行为（原文： Exercise some program logic for correct behavior. 我的理解是这个函数用来测试代码逻辑的正确性），注意与下面的 `example function` 的区别。 `go test` 调用 test function 并报告结果（ PASS 或者 FAIL）。示例如下：

		```go
		import "testing"

		func TestDemo(t *testing.T) {	 // The function name must begin with Test; the optional suffix Name must begin with a capital letter!
			// Do some work.
		}
		```

	- `benchmark function` ：以 Benchmark 开头命名的函数。该函数是测量一个程序在固定工作负载下的性能。它强调测试代码的**性能**，它能提供一个参数 N ，用于指定操作执行的循环次数。示例如下：

		```go
		import "testing"

		func BenchmarkDemo(b *testing.B) {	// The function name must begin with Test; the optional suffix Name must begin with a capital letter!
			for i := 0; i < b.N; i++ {		// 循环N次
				Demo(...)	// Do some work.
			}
		}
		```

	- `example function` ：以 Example 开头命名的函数，它没有参数和返回值，也不导入 testing 包。示例如下：

		```go
		func ExampleDemo() {	// There is no import.
			// Do some work.
		}
		```

		- 主要用于文档：一个包的例子可以更加简洁直观地演示函数的用法，它还可以方便展示属于同一个接口的几种类型或函数直接的关系。同时注意，示例函数和注释不一样。示例函数需要接受编译器的编译时检查，这样可以保证示例代码不会腐烂成不能使用的旧代码。
		- 在使用 `go test` 执行测试的时候也会运行示例函数，可以与示例函数下面的输出样式进行比对（检查输出结果是否与输出样式匹配）。
		- 提供真实的演练场， 类似于在 golang 官网上直接运行代码， EXM ？
	
	> 在 Gogland 上进行测试的时候有时会出现 `Cannot find package "` 的提示，这是中文路径名造成的，需要更改路径。目前我使用的方法是强制 Run ，也能得到同样的测试结果。

- `TestMain` ： go test 的初始化，有些测试要在开始测试之前/后，打开/关闭链接，或者在测试开始前恢复数据库状态，则需要使用到该函数。
	```go
	func TestMain(m *testing.M) {
	    	fmt.Println("begin")	// 测试开始前的控制台输出
		m.Run()			// 执行测试函数，这里指具体的测试函数 Testxxx
		clearTables()		// 测试结束之前的操作，这里就简单地清除表记录
	}
	```
	输出看起来像这样：
	```
	begin 
	=== RUN Testxxx
	— PASS: Testxxx (0.00s) 
	PASS 
	```

## 反射

- `reflect 包`：定义了两个重要的类型——
	- `Type` ：表示一个 Go 类型，他是一个 interface 。 `reflect.TypeOf` 接收任意 interface{} 并且返回它的动态类型：

		```go
		t := reflect.TypeOf(3)		// a reflect.Type
		fmt.Println(t.String())	// "int"
		fmt.Println(t)			// "int"
		```

		注意： `reflect.TypeOf` 是返回 interface 值的动态类型（不返回 interface ？），所以下面的程序打印 `*os.File` 而不是 `io.Writer` ：

		```go
		var w io.Writer = os.Stdout
		fmt.Println(reflect.TypeOf(w))	// "*os.File"
		```

		然而 `reflect.Type` 可以返回 interface 类型。
	- `Value` ：持有任何一个类型的值。相应地 `reflect.ValueOf` 函数接收任意 interface{} 并返回一个包含 interface 的动态值的 `reflect.Value` ：

		```go
		v := reflect.ValueOf(3)	// a reflect.Value
		fmt.Println(v)			// "3"
		fmt.Println("%v\n", v)		// "3"
		fmt.Println(v.String())	// NOTE: "<int Value>"
		```

- 少使用反射：
	- 基于反射的代码是比较脆弱的。对于每一个会导致编译器报告类型错误的问题，在反射中都有与之相对应的问题，不同的是编译器会在构建时马上报告错误，而反射则是在真正运行到的时候才会抛出 panic 异常，可能是写完代码很久之后了，而且程序也可能运行了很长的时间。避免这种问题的方法是将所有的反射相关使用控制在包的内部，尽量避免在报的 API 中直接暴露 `reflect.Value` 类型，如果无法避免则在每个风险操作前进行额外的类型检查。
	- 反射降低了程序的安全性，还影响了自动化重构和分析工具的准确性，因为它们无法识别运行时才能确认的类型信息。
	- 反射的操作不能做静态类型检查（这很好理解，因为反射目的是在运行时动态获取类型信息的），大量反射的代码通常让人难以理解。
	- 基于反射的代码通常比正常的代码运行速度慢一到两个数量级（在实际中对此因素的考虑不多，因为大部分函数的性能和程序的整体性能关系不大，使用反射反而会使程序更加清晰，测试的时候每个数据集都很小所以也适合使用反射。但是对于性能关键的函数最好避免使用反射）。


## unsafe包

- 在优化内存空间时使用它们返回的结果对于理解原生的内存布局有帮助（这三个函数并不是不安全，只是它们涉及到了系统底层）：
	- `unsafe.Sizeof`：返回操作数在内存中的字节大小。参数可以是任意的表达式，但是它不会对表达式求值。
	- `unsafe.Aligmof`：返回对应参数的类型需要对齐的倍数，和 Sizeof 类似，它也返回一个常量表达式。
	- `unsafe.Offsetof`：参数必须是一个字段 `x.f` ，然后返回 f 字段相对于 x 起始地址的偏移量，包括可能的空洞（内存空洞：编译器自动添加的没有被使用的内存空间，用于保证后面每个字段或元素的地址相对于结构或数组的开始地址能够合理地对齐，对齐效率会更高）。

- `unsafe.Pointer`：指针类型，它可以包含任意类型变量的地址。
	- 因为我们不知道 unsafe.Pointer 类型变量 p 的具体类型，所以不能直接通过 `*p` 获取 unsafe.Pointer 指针指向的变量的值。
	- unsafe.Pointer 指针可以相互进行比较，并且支持和 nil 常量比较来判断是否为空指针。
	- 普通的 `*T` 指针可以和 unsafe.Pointer 指针相互转化：

		```go
		fun Float64bits(f float64) uint64 { return *(*uint)(unsafe.Pointer(&f))}

		fmt.Println("%#016x\n", Float64bits(1.0))	// "0x3ff0000000000000"
		```

		- 这种指针转换语法让我们可以在不破坏类型系统的前提下下内存写入任意的值。
		- 注意到，普通指针 \*float64 转换为 unsafe.Pointer 类型指针，然后 unsafe.Pointer 指针转换成 \*uint64 并输出结果，说明被转回普通的指针类型并不需要和原始的 \*T 类型相同。
		- 建议不使用它，因为使用它特别容易出错。
	- 一个 unsafe.Pointer 指针可以转换为 uintptr 类型（整型，它保存指针的值，当然这个值是唯一的。 uintptr is an integer type that is large enough to hold the bit pattern of any pointer.）。然后保存到指针型数值变量中（意思是把指针的值保存到变量中，而不是保存该指针）。该过程可逆，但是将 uintptr 转换为 unsafe.Pointer 指针可能会破坏类型系统，因为并不是所有的数字都是有效地址。下面出示一个例子：
		```go
		// This example proves that uintptr only stores the value of the pointer, not the pointer itself.
		s := 3
		t := &s
		ptr := unsafe.Pointer(t)
		t1 := uintptr(ptr)
		fmt.Println(ptr)		// 0xc0420361d0
		fmt.Println("t1:", t1)		// t1: 825741238736
		t1 = 666666666666		// modify the value of uintptr
		fmt.Println("Updated!", t1)	// Updated! 666666666666
		t2 := uintptr(ptr)
		fmt.Println("t3:", t2)		// t3: 825741238736
		```
		注意将 unsafe.Pointer 指针转为原生数字，然后再转为 unsafe.Pointer 类型指针的操作也是不安全的：
		```go
		// Subtly incorrect!
		var x struct{
			a bool
			b int16
			c []int
		}

		tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
		pb := (*int16)(unsafe.Pointer(tmp))
		*pb = 42
		fmt.Println(x.b)		// for the most part prints the right answer 42
		```
		有时候垃圾回收器会移动一些变量以降低内存碎片等问题（这种情况少见，所以上面的例子大多数情况下都会正确输出 `42` ），这种垃圾回收器被称为`移动GCs`。当一个变量被移动，所有保存该变量旧地址的指针都要被更新为变量移动后的新地址。从垃圾回收器的角度看， unsafe.Pointer 是一个指向变量的指针，并且随着变量的移动，指向该变量的 unsafe.Pointer 指针也要被更新。但是 uintptr 只是一个对 unsafe.Pointer 的**值的拷贝**的数据，它只是一个数字，即使因为移动GCs而改变了 unsafe.Pointer 的值， uintptr 仍然保存的是移动GCs**移动变量之前**对 unsafe.Pointer 的值的拷贝。正确的操作是这样的：
		```go
		// The example below takes the address of variable x, adds the offset of its
		// b field, converts the resulting address to *int16, and through that pointer
		// updates x.b
		var x struct{
			a bool
			b int16
			c []int
		}

		pb := (*int16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
		*pb = 42
		fmt.Println(x.b)
		```


## 忠告

- 如果想要更新变量的值或者该变量太大而要避免这种变量的拷贝，我们建议在实际情况中使用指针来操作。对于方法来说也是如此：方法的参数如果太大或者要修改参数本身的值我们也推荐使用指针的形式，就像这样：

    ```go
    fun (p *Point) ScaleBy(factor float64) {
        p.X *= factor
        p.Y *= factor
    }
    ```

    **注意：**
    调用时应该使用 ``` (*Point).ScaleBy ``` 。不加括号是错误的调用方式！

- 基于指针对象的方法需要注意：
    - 不管method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，go本身会帮助我们进行类型转换。
    - 在声明一个method的receiver该是指针还是非指针类型的时候，需要从两个方面考虑：
    	- 这个对象本身是不是特别大，如果声明为一个非指针类型，调用会产生一次拷贝，消耗空间；
    	- 如果使用指针类型，要注意这个指针类型始终指向同一块内存地址，即便对其进行了拷贝也是如此。

- 在编写并发代码的时候，要建立一种思想—— `对并发的直觉总是不能被信任的！` 。在完成代码的编写过后，记得使用 `go run -race xx.go` 来检测代码的 race condition 。

- 永远不要过早地优化代码。

- 尽量避免使用 `reflect` 和 `unsafe` 包。使用他们可能会带来实现上的便利，但是这牺牲了程序的健壮性及安全性，在后期维护中会带来莫名其妙的问题。

- 避免使用 cgo 调用 C 代码。除非对性能要求很高，越底层的代码越难控制程序的运行状态，越容易出现莫名其妙的问题。
