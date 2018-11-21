package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

// 哈希，又叫散列（hash），将任意长度的消息压缩到某一固定长度的消息摘要的函数。
// 特点是哈希操作不可逆，可用于文件完整性校验（或数字签名）、简单的密码存储（不够安全）。

// base64，常见的用于传输8Bit字节码的编码方式之一。特点是操作可逆（加解码）
// 用作HTTP表单和HTTP GET URL中的参数，采用Base64编码具有不可读性，需要解码后才能阅读。

func main() {
	// 入门级存储密码方式：hash
	//
	// 简单将密码进行hash操作，优点是：
	// 1.hash操作简单，消耗小
	// 2.得到的摘要唯一，碰撞的可能性小
	// 缺点是：攻击者容易得到一个常用密码的hash表（rainbow table），
	// 攻击者可以通过尝试rainbow table上的摘要暴力破解
	h := sha256.New()
	io.WriteString(h, "Hash string...")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	h = sha1.New()
	io.WriteString(h, "Hash string...")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	h = md5.New()
	io.WriteString(h, "Hash string...")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	// 进阶级存储密码方式：hash+盐
	//
	// 将明文密码hash后加上不公开的随机串再进行一次hash
	h = md5.New()
	io.WriteString(h, "This is password.")
	passwdmd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(passwdmd5) // 4ea8283560f34ac59e4115391bab11db

	salt := "aaron" // salt，不公开的随机字符串
	io.WriteString(h, salt)
	io.WriteString(h, passwdmd5)
	fmt.Printf("% x", h.Sum(nil))

	// 专家级存储密码方式：增加破译密码存储方式的成本例如增大计算量以致破译在有限时间内不可行
	fmt.Println("")

	// base64编码
	welcome := "你好，世界！ hello world"
	encode := base64.StdEncoding.EncodeToString([]byte(welcome))
	fmt.Println(encode)

	// base64解码
	decode, err := base64.StdEncoding.DecodeString(encode)
	// typically used in URLs and file names
	// decode, err = base64.URLEncoding.DecodeString(encode)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(decode))
}
