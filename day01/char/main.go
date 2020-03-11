package main

import "fmt"

func main() {
	a1 := "Golang"
	a2 := 'G' //ASCII码下占一个字节（8位 8bit）
	fmt.Println(a1, a2)
	a3 := "中国"
	a4 := '中' //UTF-8编码下一个中文占3个字节
	fmt.Println(a3, a4)
	a5 := "hello傻瓜"
	fmt.Println(len(a5))

	//遍历字符串
	for i := 0; i < len(a5); i++ {
		fmt.Println(a5[i])        //默认使用ascii打印
		fmt.Printf("%c\n", a5[i]) //%c 8位字符，中文会出现乱码
	}

	fmt.Println()

	//for range 循环 是按照rune类型去遍历的
	for _, v := range a5 {
		fmt.Printf("%c\n", v)
	}

	for k, v := range a5 {
		fmt.Printf("%d%c\n", k, v)
	}

	//强制类型转换
	a6 := "big"
	//将字符串强制转换成字节数组类型
	byteArray := []byte(a6)
	fmt.Println(byteArray)
	byteArray[0] = 'i'
	fmt.Println(byteArray)
	//将字节数组类型强制转换成字符串类型
	a6 = string(byteArray)
	fmt.Println(a6)

	//hello==>olleh
	b1 := "hello"
	fmt.Println(b1)
	bArray := []byte(b1)
	bArray[0] = 'o'
	bArray[1] = 'l'
	bArray[2] = 'l'
	bArray[3] = 'e'
	bArray[4] = 'h'
	b1 = string(bArray)
	fmt.Println(b1)

	b2 := "hello"
	b2Array := []byte(b2)
	b3 := ""
	for i := len(b2Array) - 1; i >= 0; i-- {
		b3 += string(b2Array[i])
	}
	fmt.Println(b3)

	b4 := "hello"
	b4Array := []byte(b4)
	alength := len(b4Array)
	for i := 0; i < alength/2; i++ {
		b4Array[i], b4Array[alength-1-i] = b4Array[alength-1-i], b4Array[i]
	}
	b4 = string(b4Array)
	fmt.Println(b4)
}
