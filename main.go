package main

import (
	"fmt"
	"syscall"
)

func main() {
	dll, err := syscall.LoadDLL("exportgo.dll")
	if err != nil {
		fmt.Println("Failed to load DLL:", err)
		return
	}
	defer dll.Release()

	proc, err := dll.FindProc("Sum")
	if err != nil {
		fmt.Println("Failed to find function in DLL:", err)
		return
	}

	result, _, _ := proc.Call(991, 3)

	fmt.Println("Result:", result)
}
