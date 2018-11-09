# 关于 Golang 的参数传递

这里我们着重讨论参数传递的方式以及在 Golang  中函数调用前后（当然包括参数传递）对实参的影响。先了解一些基本概念。


## 参数传递

###  定义

参数传递，是在程序运行中，实际参数就会将参数值传递给相应的形式参数，然后在函数中实现对数据处理和返回的过程。

- **实际参数**：简称实参，在调用函数/方法时，从主调过程传递给被调用过程的参数值。实参可以是变量名、数组名、常数或者表达式。
- **形式参数**：简称形参，指出现在函数/方法形参表中的变量名。函数/方法在被调用前没有为他们分配内存，其作用是说明自变量的类型和形态以及在过程中的作用。
- **实参与形参的关系**：
1. 形参只能是变量（要指明它的数据类型）；实参可以是变量、常量或者表达式。
2. 实参与形参的个数、位置以及它们对应的数据类型应当一致。
3. 调用函数时若出现实参是数组名，则传递给形参的是数组的首地址。
4. 实参传递给形参是单向传递。形参变量在未出现函数调用时并不占用内存，只在调用时才占用。调用结束后将释放内存。

### 方法

**按值传递参数**

按值传递参数时，是将实参变量的值复制到一个临时存储单元中。如果在调用过程中改变了形参的值，不会影响实参变量本身，即实参变量保持调用前的值不变。

**按地址传递参数**

按地址传递参数时，把实参变量的地址传送给被调用过程，实参和形参共用同一内存地址。在被调用过程中，形参的值一旦改变。相应实参的值也跟着改变。如果实参是一个常数或者表达式（不含变量的表达式，也可当作常数），则按传值方式处理。

**按数组传递参数**

按照按地址传递的方式传递数组。当数组作为实参传递给函数/方法，系统将实参数组的起始地址传给过程使形参数组也具有与实参数组相同的起始地址。


## Golang 中的参数传递

### 值传递

事实证明 Golang 的参数传递（常用的类型如： string 、 int 、 bool 、array 、 slice 、 map 、 chan ）都是值传递。

```go
func main() {
	b := false
	fmt.Println("b's address is:", &b)
	bo(b)
	fmt.Println(b)
}

func bo(b bool) {
	fmt.Println("this address is different from the original address:", &b)
	b = true
}

// Output:
// b's address is: 0xc0420361ac
// this address is different from the original address: 0xc0420361fa
// false
```

从上面代码可以看出在函数中修改值不会影响实参的原始值。其余的类型读者自行尝试输出查看结果。若要在函数中改变实参的值，则使用**指针传递**：

```go
var i int = 5

func main() {
	modify(&i)
	fmt.Println(i)
}

func modify(i *int) {
	*i = 6
}

// Output:
// 6
```


### 关于 slice 的参数传递

**数组的参数传递**

使用数组元素(array[x])或者数组(array)作为函数参数时，其使用方法和普通变量相同。即是`值传递`。

```go
func modifyElem(a int) {
	a += 100
}

func modifyArray(a [5]int) {
	a = [5]int{5,5,5,5,5}
}

func main() {
	var s = [5]int{1, 2, 3, 4, 5}
	modifyElem(s[0])
	fmt.Println(s[0])
	modifyArray(s)
	fmt.Println(s)
}

// Output:
// 1
// [1 2 3 4 5]
```

**slice 的参数传递**

slice 作为实参传入函数也是进行的`值传递`。我们可以通过代码来证明：

```go
// modify s
func modify(s []int) {
	fmt.Printf("%p \n", &s)
	s = []int{1,1,1,1}
	fmt.Println(s)
	fmt.Printf("%p \n", &s)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]
	fmt.Printf("%p \n", &s)
	modify(s)
	fmt.Println(s[3])
}

// Output:
// 0xc042002680 
// 0xc0420026c0 
// [1 1 1 1]
// 0xc0420026c0 
// 4
```

可以看到，实参传递之前的地址和在函数里面的地址是不同的，而且在函数里面修改实参的值也不会影响实参的实际值。当然，在函数里面对 slice 进行重新赋值不会改变它的地址（因为这里输出了两个相同的地址）。但是下面一段代码可能有点让人迷惑：

```go
// modify s[0] value
func modify(s1 []int) {
	s1[0] += 100
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]
	modify(s)
	fmt.Println(s[0])
}

// Output:
// 101
```

在 StackOverFlow 上有人做出了解释， [Are golang slices pass by value?](http://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value) 摘要如下：

> Everything in Go is passed by value. Slices too. But a slice value is a header, describing a contiguous section of a backing array, and a slice value only contains a pointer to the array where the elements are actually stored. The slice value does not include its elements (unlike arrays).
> 
> So when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the same backing array. Modifying the elements of the slice implies modifying the elements of the backing array, and so all slices which share the same backing array will "observe" the change.

简单地说： slice 作为参数传递给函数其实是传递 slice 的值，这个值被称作一个 `header` ，**它只包含了一个指向底层数组的指针**。当向函数传递一个 slice ，将复制一个 header 的副本，这个副本包含一个指向同一个底层数组的指针。修改 slice 的元素间接地修改底层数组的元素，也就是所有指向同一个底层数组的 slice 会响应这个变化，主函数的 slice 也就一同修改了 s[0] 的值。

为了验证这个问题，对上面的代码进一步修改：

```go
// modify s[0] value
func modify2(s []int) {
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	s[0] += 100
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]

	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	modify2(s)
	fmt.Println(s[0])
}

// Output:
// 0xc04203c400 
// 0xc042039f50 
// 0xc04203c440 
// 0xc042039f50 
// 0xc04203c440 
// 0xc042039f50 
// 101
```

在每个操作上面输出它的地址，显示地址不变，说明修改会对  `main` 函数里面的 slice 产生影响。结果证明，实参传递给函数的只是一个 slice 的副本，它们不是指向同一个内存地址的。在 `main` 函数和 `modify2` 函数里面我们打印了 s[0] 的内存地址，发现它们的内存地址是相同的，所以当我们在 `modify2` 函数里面修改 s[0] 会影响 s[0] 的原始值。

#### 进一步探究

```go
// modify s[0] value
func modify2(s []int) {
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	s[0] += 100
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:3]

	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	fmt.Println(cap(s), len(s))
	s1 := append(s, 1, 2)
	fmt.Println(cap(s1), len(s1))
	modify2(s)
	fmt.Println(s)
	fmt.Println(s1)
}

// Output:
//0xc042002440
//0xc04200a270
//5 3
//5 5
//0xc042002480
//0xc04200a270
//0xc042002480
//0xc04200a270
//[101 2 3]
//[101 2 3 1 2]
```
然后：
```go
// modify s[0] value
func modify2(s []int) {
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	s[0] += 100
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:3]

	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &s[0])
	fmt.Println(cap(s), len(s))
	s1 := append(s, 1, 2, 3)
	fmt.Println(cap(s1), len(s1))
	modify2(s)
	fmt.Println(s)
	fmt.Println(s1)
}

//Output:
//0xc04204c3e0
//0xc042070030
//5 3
//10 6
//0xc04204c420
//0xc042070030
//0xc04204c420
//0xc042070030
//[101 2 3]
//[1 2 3 1 2 3]
```
这个结果正好验证了[StackOverFlow](http://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value)上的结论，当slice指向的底层数组不一致时，对数组元素的改变便不会影响到slice的改变，毕竟，两个slice代表的底层数组都不同了。

### 结论

事实证明： Golang 中所有基本数据类型的参数传递都是通过值传递完成的。


## 对非基本类型数据（ interface 等）的参数传递及其变量复制的一些讨论
