package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"strings"
	"time"
)

var Key registry.Key

var KeyArray []KeyMap

type KeyMap struct {
	id     int
	name   string
	binary []byte
}

func queryKeyByBinary(binary []byte) KeyMap {
	for _, keyMap := range KeyArray {
		if bytes.Equal(binary, keyMap.binary) {
			return keyMap
		}
	}
	return KeyMap{}
}
func queryKeyByName(name string) KeyMap {
	for _, keyMap := range KeyArray {
		if strings.Replace(name, " ", "", -1) == strings.Replace(keyMap.name, " ", "", -1) {
			return keyMap
		}
	}
	return KeyMap{}
}

func InitKeyArray() {
	KeyArray = append(KeyArray, KeyMap{
		name:   "L Ctrl",
		binary: []byte{29, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		name:   "L Win",
		binary: []byte{91, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		name:   "L Alt",
		binary: []byte{56, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		name:   "R Alt",
		binary: []byte{56, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		name:   "R Win",
		binary: []byte{92, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		name:   "R Ctrl",
		binary: []byte{29, 224},
	})

}

func showCurrectConfig() {
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("error:", a)
			time.Sleep(time.Second * 10)
		}
	}()
	var err error
	Key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Keyboard Layout`, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		fmt.Println("注册表打开失败!")
		return
	}
	value, _, err2 := Key.GetBinaryValue("Scancode Map")
	if err2 != nil || len(value) < 9 {
		fmt.Println("未查询当前按键修改")
	} else if value[8] == 1 {
		fmt.Println("未查询当前按键修改")
	} else {
		value = value[12:]
		fmt.Println("当前按键修改内容如下")
		for i := 0; i < len(value); i = i + 4 {
			key1 := value[i : i+2]
			key2 := value[i+2 : i+4]
			if bytes.Equal(key1, key2) {
				continue
			}
			fmt.Println("    ", queryKeyByBinary(key1).name, "  ->  ", queryKeyByBinary(key2).name)

		}
	}
}
func delConfig() bool {
	err2 := Key.DeleteValue("Scancode Map")
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("没找到相应键值")
	}
	return err2 == nil
}
func useConfig(config []byte) bool {
	err2 := Key.SetBinaryValue("Scancode Map", config)
	return err2 != nil
}
func genRegValue(strs []string) (result []byte) {
	result = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, int32(len(strs)+1))
	result = append(result, buffer.Bytes()...)
	for _, str := range strs {
		split := strings.Split(str, "->")
		config := append(queryKeyByName(split[0]).binary, queryKeyByName(split[1]).binary...)
		result = append(result, config...)
	}

	return
}

func createConfig() (newConfig []string) {
	fmt.Println("请根据提示修改配置：")
	mp := make(map[int]string)
	mp[1] = "L Ctrl"
	mp[2] = "L Win"
	mp[3] = "L Alt"
	mp[4] = "R Ctrl"
	mp[5] = "R Win"
	mp[6] = "R Alt"
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d = %s\n", i, mp[i])
	}

	for i := 1; i <= 6; i++ {
		fmt.Printf("请输入%s修改后的序号", mp[i])
		var n int
		fmt.Scan(&n)
		if i != n {
			newConfig = append(newConfig, "  "+mp[i]+"  ->  "+mp[n])
		}
	}
	fmt.Println("你要修改的键位有：")
	for _, s := range newConfig {
		fmt.Println(s)
	}
	return
}

func main() {
	InitKeyArray()
	showCurrectConfig()
	for {
		config := createConfig()
		fmt.Println("输入Y确认修改,输入D恢复系统默认配置,输入R重新配置,输入其他按键退出...")
		var n string
		fmt.Scan(&n)
		n = strings.ToUpper(n)
		if n == "Y" {
			useConfig(genRegValue(config))
			break
		} else if n == "D" {
			delConfig()
			break
		} else if n == "R" {
			continue
		} else {
			os.Exit(1)
		}

	}
	fmt.Println("修改完毕注销用户重新登录或重启电脑后生效,程序将5秒后退出！")
	time.Sleep(time.Second * 5)

}
