package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

//json gob
type s struct {
	data map[string]interface{}
}

func jsonDemo(){
	var s1=s{
		data:make(map[string]interface{},8),
	}
	s1.data["count"]=1
	ret,err:=json.Marshal(s1.data)
	if err != nil {
		fmt.Println("maeshal failed",err)
	}
	fmt.Printf("%#v\n",string(ret))
	var s2=s{
		data:make(map[string]interface{},8),
	}
	err=json.Unmarshal(ret,&s2.data)
	 if err!=nil{
	 	fmt.Println("Unmarshal failed",err)
	 }
	fmt.Println(s2)
	for _,v:=range s2.data{
		fmt.Printf("value:%v,type:%T\n",v,v)
	}
}

func gobDemo(){
	var s1=s{
		data:make(map[string]interface{},8),
	}
	s1.data["count"]=1 //int类型

	//encode 编码
	buf:=new(bytes.Buffer)//返回的是对应类型的指针
	enc:=gob.NewEncoder(buf)//造一个编码器对象
	err:=enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed,err:",err)
		return
	}
	b:=buf.Bytes()//拿到编码之后的字节数据
	fmt.Println(b)
	var s2=s{
		data:make(map[string]interface{},8),
	}

	//decode 解码
	dec:=gob.NewDecoder(bytes.NewBuffer(b)) //造一个新的解码器对象
	err=dec.Decode(&s2.data) //把二进制数据解码到s2.data
	if err!=nil{
		fmt.Println("gob decode failed,err:",err)
		return
	}
	fmt.Println(s2.data)
	for _,v:=range s2.data{
		fmt.Printf("value:%v,type:%T\n",v,v)
	}
}

func main(){
	//jsonDemo()
	gobDemo()
}

