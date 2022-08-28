package keyMap

import (
	"bytes"
	"strings"
)

var KeyArray []KeyMap

type KeyMap struct {
	Id     int
	Name   string
	Binary []byte
}

func QueryKeyByBinary(binary []byte) KeyMap {
	for _, keyMap := range KeyArray {
		if bytes.Equal(binary, keyMap.Binary) {
			return keyMap
		}
	}
	return KeyMap{}
}
func QueryKeyByName(name string) KeyMap {
	for _, keyMap := range KeyArray {
		if strings.Replace(name, " ", "", -1) == strings.Replace(keyMap.Name, " ", "", -1) {
			return keyMap
		}
	}
	return KeyMap{}
}

func QueryKeyById(id int) KeyMap {
	for _, keyMap := range KeyArray {
		if keyMap.Id == id {
			return keyMap
		}
	}
	return KeyMap{}
}

func InitKeyArray() {
	KeyArray = append(KeyArray, KeyMap{
		Id:     1,
		Name:   "L Ctrl",
		Binary: []byte{29, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     2,
		Name:   "L Win",
		Binary: []byte{91, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     3,
		Name:   "L Alt",
		Binary: []byte{56, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     4,
		Name:   "R Alt",
		Binary: []byte{56, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     5,
		Name:   "R Win",
		Binary: []byte{92, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     6,
		Name:   "R Ctrl",
		Binary: []byte{29, 224},
	})

}
