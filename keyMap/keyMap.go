package keyMap

import (
	"bytes"
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
		Name:   "左侧 Ctrl",
		Binary: []byte{29, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     2,
		Name:   "左侧 Win",
		Binary: []byte{91, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     3,
		Name:   "左侧 Alt",
		Binary: []byte{56, 0},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     4,
		Name:   "右侧 Alt",
		Binary: []byte{56, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     5,
		Name:   "右侧 Win",
		Binary: []byte{92, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     6,
		Name:   "右侧 Ctrl",
		Binary: []byte{29, 224},
	})
	KeyArray = append(KeyArray, KeyMap{
		Id:     7,
		Name:   "右侧 Menu",
		Binary: []byte{93, 224},
	})

}
