# Golang http 请求


### 使用原生 http 包发起 GET 请求

```go
resp, err := http.Get("https://www.github.com/")
if err != nil {
	fmt.Println(err)
}

defer resp.Body.Close()		// close connection after finish
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	fmt.Println(err)
}

fmt.Println(string(body))	// Converted body to string and print
```
