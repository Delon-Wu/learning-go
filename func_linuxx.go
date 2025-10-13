//go:build linuxx
// +build linuxx

package main

import "fmt"

func PlatformSpecificFunction() {
	fmt.Println("This is Linux implementation")
}
