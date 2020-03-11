package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//template示例

//User 自定义一个结构体
type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {

	//打开一个模板文件
	htmlByte, err := ioutil.ReadFile("./info.html")

	//添加自定义方法要在parse模板文件之前
	//1.自定义一个函数
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	//2.把自定义的函数告诉模板系统
	// template.New("info") //创建一个Template对象
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}) //给模板系统追加自定义函数
	//链式操作 原理：每一次执行完方法之后，返回操作的对象本身
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte)) //解析模板

	// template.New("info") //创建一个Template对象,名字叫 info
	// template.ParseFiles("./info.html") //把文件的名字当成模板的名字

	if err != nil {
		fmt.Println("open html failed,err:", err)
		return
	}
	//用数据渲染模板
	// data := "<li>三体</li>"
	// t.Execute(w, data)

	//用结构体数据渲染
	// user := User{
	// 	UserName: "豪杰",
	// 	Password: "123456",
	// 	Age:      12,
	// }
	// t.Execute(w, user)

	//用map渲染模板
	userMap := map[int]User{
		1: {"haojie", "1234", 18},
		2: {"hanxin", "1234", 24},
		3: {"guandama", "1234", 5},
	}
	t.Execute(w, userMap)

}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8090", nil)
}
