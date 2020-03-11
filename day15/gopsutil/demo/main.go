package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

//获取CPU信息
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

//cpu负载 （windows下：还没有实现）
func getLoad() {
	info, err := load.Avg() // windows下：not implemented yet
	if err != nil {
		fmt.Printf("load.avg() failed,err:%v", err)
		return
	}
	fmt.Println(info)
}

//Mem 内存信息
func getMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed,err:%v", err)
		return
	}
	fmt.Println(info)
}

// Host 主机信息
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

//Disk 磁盘信息
func getDiskInfo() {
	//获取所有的分区信息
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	fmt.Println(parts) //打印所有的分区信息

	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Println(diskInfo)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
	//磁盘IO
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

//网络IO
func getNetInfo() {
	info, _ := net.IOCounters(true)
	//fmt.Println(info)
	for _, netIO := range info {
		fmt.Println(netIO)
	}
	//for index, v := range info {
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	//}
}

func main() {
	//getCpuInfo()
	//getLoad() （not implemented yet: windows下还没有实现）
	//getMemInfo()
	//getHostInfo()
	//getDiskInfo()
	getNetInfo()
}
