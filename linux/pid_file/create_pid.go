package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PIDFile stored the process id.
type PIDFile struct {
	path string
}

// Check if the process exists.
func ProcessExists(pid string) bool {
	if _, err := os.Stat(filepath.Join("/proc", pid)); err == nil {
		return true
	}
	return false
}

// Check if the PID file already exists.
func CheckPidFileAlreadyExists(path string) error {
	if pidByte, err := os.ReadFile(path); err == nil {
		pid := strings.TrimSpace(string(pidByte))
		if ProcessExists(pid) {
			return fmt.Errorf("ensure the process:%s is not running pid file:%s", pid, path)
		}
	}
	return nil
}

// Create a PID file in the specified path. The file content is the PID of this program.
func NewPIDFile(path string) (*PIDFile, error) {
	if err := CheckPidFileAlreadyExists(path); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(os.O_EXCL)); err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		return nil, err
	}
	return &PIDFile{path: path}, nil
}

// Remove the PID file
func (file PIDFile) Remove() error {
	return os.Remove(file.path)
}
