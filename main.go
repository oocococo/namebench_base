package main

import (
	"fmt"
	"namebench/common"
	"namebench/config"
	"namebench/query"
	"sort"

	"github.com/miekg/dns"
)

func main() {
	var urls = [...]string{
		"www.qq.com",
		"www.baidu.com",
		"sinaimg.cn",
		"uestc.edu.cn",
		"agefans.cc",
		"weibo.com",
		"2miners.com",
		"bilibili.com",
	}
	var C = config.C
	if C.Iteration <= 0 {
		C.Iteration = 1
	}
	c := new(dns.Client)
	c.Net = "udp4"

	var result []common.NameserverTable

	for _, nameserver := range C.Nameservers {
		tmp := new(common.NameserverTable)
		tmp.Nameserver = nameserver
		tmp.Ch = make(chan int64, len(urls)*C.Iteration)
		result = append(result, *tmp)
	}

	pool := make(chan struct{}, C.Concurrency)
	for i := 0; i < C.Iteration; i++ {
		for _, url := range urls {
			for _, nameserver := range result {
				go query.RunWithLimit(c, url, nameserver, pool)
			}
		}
	}
	// count Rtt
	for i := range result {
		for count := 0; count < len(urls)*C.Iteration; count++ {
			result[i].Rtt += <-result[i].Ch
		}
	}
	// sort nameserver by rtt
	sort.Slice(result, func(i, j int) bool {
		return result[i].Rtt < result[j].Rtt
	})
	// print result
	for _, i := range result {
		fmt.Println(i.Nameserver, i.Rtt, "ms")
	}
}
