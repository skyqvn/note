package main

import (
	"fmt"
	"os"
	"syscall"
)

//LOCK_SH 建立共享锁定。多个进程可同时对同一个文件作共享锁定。
//LOCK_EX 建立互斥锁定。一个文件同时只有一个互斥锁定。
//LOCK_UN 解除文件锁定状态。
//LOCK_NB 无法建立锁定时，此操作可不被阻断，马上返回进程。通常与LOCK_SH或LOCK_EX 做OR(|)组合。

// 文件锁
type FileLock struct {
	path string
	f    *os.File
}

// 创建锁
func New(path string) *FileLock {
	return &FileLock{
		path: path,
	}
}

// 加锁
func (l *FileLock) Lock() error {
	f, err := os.Open(l.path)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.path, err)
	}
	return nil
}

// 释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
