package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"os"
	"time"
)

func main() {
	args := os.Args
	if nil == args || len(args) < 3 {
		help("Usage: zu <add | get-all | get | remove | get-role | sync> [options...]")
	}

	command := args[1]
	zkUrls := args[2]

	switch command {
	case "get-all":
		getAll(zkUrls)
	case "add":
		add(zkUrls, args[3])
	case "remove":
		remove(zkUrls, args[4])
	}
}

func getAll(zkUrl string) {
	connect, _, err := zk.Connect([]string{zkUrl}, time.Second)
	if nil != err {
		os.Exit(1)
	}

	//connect.Reconfig()
	fmt.Println(connect.State())
}

func add(zkUrl string, nodeStr string) {
	fmt.Println(zkUrl, nodeStr)
	connect, _, err := zk.Connect([]string{zkUrl}, time.Second)
	if nil != err {
		os.Exit(1)
	}
	_, err = connect.IncrementalReconfig([]string{nodeStr}, []string{}, -1)
	if nil != err {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("ok")
		os.Exit(0)
	}
}

func remove(zkUrl string, nodeStr string) {
	connect, _, err := zk.Connect([]string{zkUrl}, time.Second)
	if nil != err {
		os.Exit(1)
	}
	_, err = connect.IncrementalReconfig([]string{}, []string{nodeStr}, -1)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("ok")
		os.Exit(0)
	}
}

func help(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
