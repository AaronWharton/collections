# Golang 包使用解析
主要是对日常使用的包的某些函数或功能的理解与分析。


## http

- 发起请求：
	```go
	// http.Get
	resp, err := http.Get("https://github.com/AaronWharton/Go")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	
	// http.Post
	// contentType must be application/x-www-form-urlencoded
	resp, err := http.Post("https://www.baidu.com/",
		"",
		strings.NewReader("application/x-www-form-urlencoded"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	
	// http.PostForm
	resp, err := http.PostForm("https://www.baidu.com/",
	url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	```

- HandleFunc：
  ```go
  // HandleFunc registers the handler function for the given pattern
  // in the DefaultServeMux.
  // The documentation for ServeMux explains how patterns are matched.
  func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	  DefaultServeMux.HandleFunc(pattern, handler)
  }
  ```
  函数第一个参数是路由；第二个参数是一个函数：第一个参数可用于在客户端显示信息，而第二个参数用于在服务端输出信息。


## math

### math/rand

生成随机数：
```go
r := rand.Intn(5, 10)		// create a random Int number which is between [5,10)
```
但是每次运行程序发现都是输出相同的随机数。我们需要一个“种子”（这里我们使用时间）使得每次生成的随机数不同。
```go
r := rand.New(rand.NewSource(time.Now().UnixNano()))
fmt.Println(r.Intn(100))	// now it will create different value when run this code each time
```
方法
- `rand.Intn(n int)` ：返回一个在[0,n)之间的伪随机 Int 型数字，若 n <= 0 则发生 panic 。类似的有 `Int31n` 、 `Float32` （ float 类型不带参数，直接返回0~1之间的随机小数）等。
- `rand.New(src Source)` ：简而言之，这个函数就是使用 src 的随机数值去生成一个新的随机数并将新的随机数返回。当然这个 Source 是一个 interface ，所以使用 `rand.NewSource()` 来生成新的 Source （如上代码所示）。


## context

- [Go Concurrency Patterns: Context](https://blog.golang.org/context)
