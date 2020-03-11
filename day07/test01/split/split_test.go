package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

//单元测试

/*
//如果字符串中包含分隔符 测试结果是否正确
func TestSplit(t *testing.T) {
	t.Log("如果字符串中包含分隔符 测试结果是否正确")
	got := Split("A:B:C", ":")      //调用函数得到的结果
	want := []string{"A", "B", "C"} //期望得到的结果

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v,实际得到：%v\n", want, got)
	}
}

//如果字符串中不包含分隔符 测试结果是否正确
func TestNoneSplit(t *testing.T) {
	t.Log("如果字符串中不包含分隔符 测试结果是否正确")
	got := Split("A:B:C", "*")
	want := []string{"A:B:C"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v,实际得到：%v\n", want, got)
	}
}

//分隔符是多个字符的
func TestMultiSepSplit(t *testing.T) {
	t.Log("如果分隔符是多个字符的 测试结果是否正确")
	got := Split("abcsbcabcd", "bc")
	want := []string{"a", "s", "a", "d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v,实际得到：%v\n", want, got)
	}
}

*/

//将多个测试用例放到一起组成测试组
func TestSplit(t *testing.T) {
	//定义一个存放测试数据的结构体
	type test struct {
		str  string
		sep  string
		want []string
	}
	//创建一个存放所有测试用例的map
	var tests = map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   test{"a:b:c", "*", []string{"a:b:c"}},
		"multi":  test{"abcsbcabcd", "bc", []string{"a", "s", "a", "d"}},
		"num":    test{"12321", "1", []string{"", "232", ""}},
	}

	//之前的方式
	// for name, item := range tests {
	// 	ret := Split(item.str, item.sep)
	// 	if !reflect.DeepEqual(ret, item.want) {
	// 		t.Fatalf("测试用例：%v失败，期望得到：%#v,实际得到：%#v\n", name, item.want, ret)
	// 	}
	// }

	//使用t.Run()执行子测试
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//测试之前做点什么
			t.Log("要开始测试了")
			defer func() {
				fmt.Println("沙河出太阳了 ")
			}()
			ret := Split(tc.str, tc.sep)
			t.Log("掼蛋 sb")
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("测试用例：%v失败，期望得到：%#v,实际得到：%#v\n", name, tc.want, ret)
			}
		})

	}
}

//基准测试 (性能基准)
func BenchmarkSplit(b *testing.B) {
	b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//并行测试
func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1)//设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河即有沙又有何", "沙")
		}
	})
}

//整个测试之前做的事和之后做的事 setup,teardown
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") //测试之前做的一些设置
	//如果TestMain 使用了flags,这里应该加上flag.Parse()
	retCode := m.Run()
	fmt.Println("write teardown code here... ") //测试之后做一些拆卸工作
	os.Exit(retCode)
}

//示例函数  (将标准输出结果跟OutPut代码作比较，看看是否一致)
func ExampleAdd() {
	fmt.Println(Add("官大吗", "dsb"))
	//OutPut:官大吗dsb
}
