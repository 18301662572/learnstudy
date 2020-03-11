package main

import (
	"net/http"
	"text/template"
)

//作业

func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./html.html")
	if err != nil {
		panic(err)
	}

	//判斷用戶是获取网页还是填完了要提交数据
	if r.Method == "POST" {
		//获取用户提交的数据
		r.ParseForm() //解析表单
		username := r.FormValue("username")
		pwd := r.FormValue("pwd")
		//校验用户名和密码是否正确
		//正常网站的登陆逻辑应该是连接数据库然后校验
		if username == "admin" && pwd == "123" {
			//登陆成功
			http.Redirect(w, r, "https://www.oldboyedu.com", 302)
		} else {
			//登陆失败
			//在login页面显示错误提示信息
			t.Execute(w, "用户名或密码错误")
		}
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8091", nil)
}
