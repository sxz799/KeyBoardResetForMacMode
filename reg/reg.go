package reg

import (
	"KeyBoardResetForMac/keyMap"
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"time"
)

var Key registry.Key

func ShowCurrentConfig() {
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
		time.Sleep(time.Second * 10)
		return
	}
	value, _, err2 := Key.GetBinaryValue("Scancode Map")
	if err2 != nil || len(value) < 9 {
		fmt.Println("未查询到按键修改配置")
	} else if value[8] == 1 {
		fmt.Println("未查询到按键修改配置")
	} else {
		value = value[12:]
		fmt.Println("当前按键修改内容如下:")
		for i := 0; i < len(value); i = i + 4 {
			key1 := value[i : i+2]
			key2 := value[i+2 : i+4]
			if bytes.Equal(key1, key2) {
				continue
			}
			fmt.Println("    ", keyMap.QueryKeyByBinary(key1).Name, "  ->  ", keyMap.QueryKeyByBinary(key2).Name)
		}
		fmt.Println("===================================")
		fmt.Println()
	}
}
func DelConfig() bool {
	err2 := Key.DeleteValue("Scancode Map")
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("没找到相应键值")
	}
	return err2 == nil
}
func UseConfig(config []byte) bool {
	err2 := Key.SetBinaryValue("Scancode Map", config)
	return err2 == nil
}
func GenRegValue(mp map[int]int) (result []byte) {
	result = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, int32(len(mp)+1))
	result = append(result, buffer.Bytes()...)
	for a, b := range mp {
		if a == b {
			continue
		}
		config := append(keyMap.QueryKeyById(a).Binary, keyMap.QueryKeyById(b).Binary...)
		result = append(result, config...)
	}
	result = append(result, []byte{0, 0, 0, 0}...)
	return
}
