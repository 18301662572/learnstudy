// runtime_pprof/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			//time.Sleep(time.Second)
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		defer file.Close()
		pprof.StartCPUProfile(file)//开始收集CPU性能测试数据
		defer pprof.StopCPUProfile()//停止CPU性能分析
		// 应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据。
		// 得到采样数据之后，使用go tool pprof 工具进行CPU性能分析。
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}

	//cmd: 11pprof_demo.exe -cpu=true 11pprof_demo.exe -mem=true
	//工具行命令：go tool pprof ./cpu.pprof
	//top 3
}