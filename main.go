package main

import (
	"KeyBoardResetForMac/keyMap"
	"KeyBoardResetForMac/reg"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func createConfig() map[int]int {
	fmt.Println("请根据提示修改配置：")
	fmt.Println("==============================")
	for _, key := range keyMap.KeyArray {
		fmt.Printf("   %d  =》  %s    \n\n", key.Id, key.Name)
	}
	fmt.Println("==============================")
	fmt.Println()
	newConfig := make(map[int]int)
	fmt.Println("请输入你想要设置的按键，先输入原始键位的编码,输入空格，再输入目标键位的编码，回车结束")
	fmt.Println("输入0表示设置完成")
	for {
		var a, b string
		fmt.Scan(&a)
		if a == "0" {
			break
		}
		fmt.Scan(&b)
		num1, err1 := strconv.Atoi(a)
		num2, err2 := strconv.Atoi(b)
		if err1 != nil || err2 != nil || num1 > len(keyMap.KeyArray) || num2 > len(keyMap.KeyArray) {
			fmt.Println("输入有误，请重新输入！")
			continue
		}
		newConfig[num1] = num2
		fmt.Println("你要修改配置如下：")
		for x, y := range newConfig {
			if x == y {
				continue
			}
			fmt.Printf("将  %-10s  修改为  %-10s\n", keyMap.QueryKeyById(x).Name, keyMap.QueryKeyById(y).Name)
		}
	}

	return newConfig
}

func main() {

	keyMap.InitKeyArray()
	reg.ShowCurrentConfig()
	var ok bool
	for {
		config := createConfig()
		fmt.Println()
		fmt.Println("请确认要修改的配置:")
		fmt.Println("======================")
		for x, y := range config {
			if x == y {
				continue
			}
			fmt.Printf("将  %-10s  修改为  %-10s\n", keyMap.QueryKeyById(x).Name, keyMap.QueryKeyById(y).Name)

		}
		fmt.Println("======================")
		fmt.Println()
		fmt.Println("输入Y确认修改,输入D恢复系统默认配置,输入R重新配置,输入其他按键退出...")
		var n string
		fmt.Scan(&n)
		n = strings.ToUpper(n)
		if n == "Y" {
			ok = reg.UseConfig(reg.GenRegValue(config))
			break
		} else if n == "D" {
			ok = reg.DelConfig()
			break
		} else if n == "R" {
			continue
		} else {
			os.Exit(1)
		}

	}
	if ok {
		fmt.Println("修改完毕,注销用户重新登录或重启电脑后生效,程序将在5秒后退出！")
	} else {
		fmt.Println("修改失败，请以管理员身份运行此程序！")
	}

	time.Sleep(time.Second * 5)
}
