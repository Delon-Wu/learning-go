//go:build windows
// +build windows

// 上面的是新版
// 下面是旧版

package main

import "fmt"

func PlatformSpecificFunction() {
	fmt.Println("This is Windows implementation")
}
