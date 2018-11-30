package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件的创建和删除，All可以创建和删除子目录
	_ = os.Mkdir("aaron", 0777)
	_ = os.MkdirAll("aaron/test", 0777)
	_ = os.RemoveAll("aaron/test")
	// 当目录不为空的时候无法删除，所以先使用All删除当前目录下所有文件（夹），
	// 然后再使用Remove删除当前文件夹。由此也可想到可用Remove判断文件是否为空。
	if err := os.Remove("aaron"); err != nil {
		fmt.Println(err)
	}

	// 创建文件
	f, err := os.Create("test.txt")
	// 创建文件的第二种方式，具体见源码
	// f2 := os.NewFile()
	if err != nil {
		fmt.Println(err)
	}
	// 打开文件后记得使用defer关闭文件流
	defer f.Close()

	// 写入文件
	for i := 0; i < 10; i++ {
		_, _ = f.WriteString("This is a test file!\r\n")
		_, _ = f.Write([]byte("This is also for test!\r\n"))
	}

	// 读取文件
	fl, errl := os.Open("test.txt")
	if errl != nil {
		fmt.Println(errl)
	}
	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		// 当n == 0时，表示文件读完，跳出循环
		if n == 0 {
			break
		}
		// 将读取到的数据输出到控制台
		_, _ = os.Stdout.Write(buf[:n])
	}

	// 删除文件
	errr := os.Remove("test.txt")
	if errr != nil {
		fmt.Println(errr)
	}
}
