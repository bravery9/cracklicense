package main

import (
	"fmt"
	"os"
)

func main() {
	var bLicense = make([]byte, 0)
	var cs40 = []int8{27, -27, -66, 82, -58, 37, 92, 51, 85, -114, -118, 28, -74, 103, -53, 6}
	var cs41 = []int8{-128, -29, 42, 116, 32, 96, -72, -124, 65, -101, -96, -63, 113, -55, -86, 118}
	var cs42 = []int8{-78, 13, 72, 122, -35, -44, 113, 52, 24, -14, -43, -93, -82, 2, -89, -96}
	var cs43 = []int8{58, 68, 37, 73, 15, 56, -102, -18, -61, 18, -67, -41, 88, -83, 43, -103}
	var cs44 = []int8{94, -104, 25, 74, 1, -58, -76, -113, -91, -126, -90, -87, -4, -69, -110, -42}

	// header
	for _, b := range Int32ToByteArray(int32(-889274157)) {
		bLicense = append(bLicense, b)
	}

	// authlen
	for i := 0; i < 2; i++ {
		bLicense = append(bLicense, byte(0))
	}

	// validto
	for _, b := range Int32ToByteArray(int32(29999999)) {
		bLicense = append(bLicense, b)
	}

	// watermark
	for _, b := range Int32ToByteArray(int32(1)) {
		bLicense = append(bLicense, b)
	}

	// version
	bLicense = append(bLicense, byte(99))

	// cs40key
	bLicense = append(bLicense, byte(16))
	var cs40key = Int8ArrayToByteArray(cs40)
	for _, b := range cs40key {
		bLicense = append(bLicense, b)
	}

	// cs41key
	bLicense = append(bLicense, byte(16))
	var cs41key = Int8ArrayToByteArray(cs41)
	for _, b := range cs41key {
		bLicense = append(bLicense, b)
	}

	// cs42key
	bLicense = append(bLicense, byte(16))
	var cs42key = Int8ArrayToByteArray(cs42)
	for _, b := range cs42key {
		bLicense = append(bLicense, b)
	}

	// cs43key
	bLicense = append(bLicense, byte(16))
	var cs43key = Int8ArrayToByteArray(cs43)
	for _, b := range cs43key {
		bLicense = append(bLicense, b)
	}

	// cs44key
	bLicense = append(bLicense, byte(16))
	var cs44key = Int8ArrayToByteArray(cs44)
	for _, b := range cs44key {
		bLicense = append(bLicense, b)
	}

	authLen := int16(len(bLicense)) - 6
	bAuthLen := Int16ToByteArray(authLen)
	bLicense[4] = bAuthLen[0]
	bLicense[5] = bAuthLen[1]

	fmt.Println(bLicense)
	fmt.Println("length:", authLen)
	WriteBytesToFile("cobaltstrike.auth.unsign", bLicense)

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
