package main

import (
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

// 将CPU信息写入到influxdb中
func writesCpuPoints(data *CpuInfo) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": data.CpuPercent,
	}
	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert cpu success")
}

// 将内存信息写入到influxdb中
func writesMemPoints(data *MemInfo) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{"mem": "mem"}
	fields := map[string]interface{}{
		"total": int64(data.Total),
		"available": int64(data.Available),
		"used": int64(data.Used),
		"used_percent": data.UsedPercent,
	}
	pt, err := client.NewPoint("memory", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert mem success")
}

// 将磁盘信息写入到influxdb中
func writesDiskPoints(data *DiskInfo) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	for k,v:=range data.PartitonUsageStat{
		tags:=map[string]string{"path":k}
		fields := map[string]interface{}{
			"total":int64(v.Total),
			"free":int64(v.Free),
			"used":int64(v.Used),
			"used_percent":v.UsedPercent,
			"inodes_total":int64(v.InodesTotal),
			"inodes_used":int64(v.InodesUsed),
			"inodes_free":int64(v.InodesFree),
			"inodes_used_percent":v.InodesUsedPercent,
		}
		pt, err := client.NewPoint("disk", tags, fields, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert disk success")
}

// 将网卡数据写入到influxdb中
func writesNetPoints(data *NetInfo) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	for k,v:=range data.NetIOCountersStat{
		tags:=map[string]string{"name":k}//把每个网卡存为tag
		fields := map[string]interface{}{
			"bytes_sent_rate":v.BytesSentRate,
			"bytes_recv_rate":v.BytesRecvRate,
			"packets_sent_rate":v.PacketsSentRate,
			"packets_recv_rate":v.PacketsRecvRate,
		}
		pt, err := client.NewPoint("net", tags, fields, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert net success")
}
