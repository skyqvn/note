package main

import (
	"log"
	"syscall"
)

func ListDrives() []string {
	var drives []string
	mask := driveMask()
	
	for i := 0; i < 26; i++ {
		if mask&1 == 1 {
			letter := string('A' + rune(i))
			drives = append(drives, letter+":")
		}
		mask >>= 1
	}
	
	return drives
}

func driveMask() uint32 {
	dll, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Println("Error loading kernel32.dll", err)
		return 0
	}
	handle, err := syscall.GetProcAddress(dll, "GetLogicalDrives")
	if err != nil {
		log.Println("Could not find GetLogicalDrives call", err)
		return 0
	}
	
	ret, _, err := syscall.SyscallN(handle)
	if err != syscall.Errno(0) {
		log.Println("Error calling GetLogicalDrives", err)
		return 0
	}
	
	return uint32(ret)
}
