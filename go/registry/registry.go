package main

import (
	"fmt"
	"log"
	
	"golang.org/x/sys/windows/registry"
)

func main() {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Hello Go", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()
	
	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}
	
	// 写入32位整形值
	key.SetDWordValue("DWORD", 0xFFFFFFFF)
	
	// 写入64位整形值
	key.SetQWordValue("QDWORD", 0xFFFFFFFFFFFFFFFF)
	
	// 写入字符串
	key.SetStringValue("String", "hello")
	
	// 写入多行字符串
	key.SetStringsValue("Strings", []string{"hello", "world"})
	
	// 写入二进制
	key.SetBinaryValue("Binary", []byte{0x11, 0x22})
	
	// 读取字符串值
	s, _, _ := key.GetStringValue("String")
	fmt.Println(s)
	
	// 枚举所有值名称
	values, _ := key.ReadValueNames(0)
	fmt.Println(values)
	
	// 创建三个子键
	subkey1, _, _ := registry.CreateKey(key, "Sub1", registry.ALL_ACCESS)
	subkey2, _, _ := registry.CreateKey(key, "Sub2", registry.ALL_ACCESS)
	subkey3, _, _ := registry.CreateKey(subkey1, "Sub3", registry.ALL_ACCESS)
	defer subkey1.Close()
	defer subkey2.Close()
	defer subkey3.Close()
	
	// 枚举所有子键
	keys, _ := key.ReadSubKeyNames(0)
	fmt.Println(keys)
	
	// 该键有子项，所以会删除失败
	err = registry.DeleteKey(key, "Sub1")
	if err != nil {
		fmt.Println(err)
	}
	
	// 没有子项，删除成功
	registry.DeleteKey(key, "Sub2")
}
