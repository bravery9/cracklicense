package main

import (
	"fmt"
	"os"
)

func main() {
	var blicense = make([]byte, 0)
	var cs40 = []int8{27, -27, -66, 82, -58, 37, 92, 51, 85, -114, -118, 28, -74, 103, -53, 6}
	var cs41 = []int8{-128, -29, 42, 116, 32, 96, -72, -124, 65, -101, -96, -63, 113, -55, -86, 118}
	var cs42 = []int8{-78, 13, 72, 122, -35, -44, 113, 52, 24, -14, -43, -93, -82, 2, -89, -96}
	var cs43 = []int8{58, 68, 37, 73, 15, 56, -102, -18, -61, 18, -67, -41, 88, -83, 43, -103}

	// header
	for _, b := range Int32ToByteArray(int32(-889274157)) {
		blicense = append(blicense, b)
	}

	// authlen
	for i := 0; i < 2; i++ {
		blicense = append(blicense, byte(0))
	}

	// validto
	for _, b := range Int32ToByteArray(int32(29999999)) {
		blicense = append(blicense, b)
	}

	// watermark
	for _, b := range Int32ToByteArray(int32(1)) {
		blicense = append(blicense, b)
	}

	// version
	blicense = append(blicense, byte(99))

	// cs40key
	blicense = append(blicense, byte(16))
	var cs40key = Int8ArrayToByteArray(cs40)
	for _, b := range cs40key {
		blicense = append(blicense, b)
	}

	// cs41key
	blicense = append(blicense, byte(16))
	var cs41key = Int8ArrayToByteArray(cs41)
	for _, b := range cs41key {
		blicense = append(blicense, b)
	}

	// cs42key
	blicense = append(blicense, byte(16))
	var cs42key = Int8ArrayToByteArray(cs42)
	for _, b := range cs42key {
		blicense = append(blicense, b)
	}

	// cs43key
	blicense = append(blicense, byte(16))
	var cs43key = Int8ArrayToByteArray(cs43)
	for _, b := range cs43key {
		blicense = append(blicense, b)
	}

	authlen := int16(len(blicense)) - 6
	bauthlen := Int16ToByteArray(authlen)
	blicense[4] = bauthlen[0]
	blicense[5] = bauthlen[1]

	fmt.Println(blicense)
	fmt.Println("length:", authlen)
	WriteBytesToFile("cobaltstrike.auth.unsign", blicense)

}

func Int32ToByteArray(i int32) []byte {
	arr := make([]byte, 4)
	arr[0] = byte(i>>24) & 0xFF
	arr[1] = byte(i>>16) & 0xFF
	arr[2] = byte(i>>8) & 0xFF
	arr[3] = byte(i) & 0xFF
	return arr
}

func Int16ToByteArray(i int16) []byte {
	arr := make([]byte, 2)
	arr[0] = byte(i>>8) & 0xFF
	arr[1] = byte(i) & 0xFF
	return arr
}

func WriteBytesToFile(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(data)
}

func Int8ArrayToByteArray(data []int8) []byte {
	arr := make([]byte, len(data))
	for i, v := range data {
		arr[i] = byte(v)
	}
	return arr
}
