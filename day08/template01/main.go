package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

//template示例

func info(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	num := rand.Intn(10)
	dataStr := string(data)
	if num > 5 {
		dataStr = strings.Replace(dataStr, "{ooxx}", "<li>三体</li>", 1)
	} else {
		dataStr = strings.Replace(dataStr, "{ooxx}", "<li>对子哈特1枚</li>", 1)
	}
	w.Write([]byte(dataStr))
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8090", nil)
}
