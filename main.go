package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Load the DLL
	dll, err := syscall.LoadDLL("mydll.dll")
	if err != nil {
		fmt.Println("Failed to load DLL:", err)
		return
	}
	defer dll.Release() // Release the DLL when we're done

	// Get a handle to a function in the DLL
	proc, err := dll.FindProc("MyFunction")
	if err != nil {
		fmt.Println("Failed to find function in DLL:", err)
		return
	}

	// Call the function
	result, _, _ := proc.Call()

	fmt.Println("Result:", result)
}
