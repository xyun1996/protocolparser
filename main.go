package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func main() {
	var ip = &IP{}
	b, err := hexStream2bytes("450000282538400080069aa0c0a80645681a0bf0")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ip.decode(b)
	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("%+v\n", ip)
	}
}

//450000282538400080069aa0c0a80645681a0bf0
func hexStream2bytes(hex string) ([]byte, error) {
	if len(hex)%2 != 0 {
		return nil, errors.New("hex 长度 无法被2整除")
	}
	hex = strings.ToUpper(hex)
	b := make([]byte, len(hex)/2)
	buff := bytes.NewBufferString(hex)
	for i := 0; i < len(hex); i += 2 {
		info := buff.Next(2)
		b[i/2] = hex2dec(info[1]) | hex2dec(info[0])<<4
		//fmt.Printf("%c %c -> %d\n", info[0], info[1], b[i/2])
	}
	fmt.Println(b)
	return b, nil
}

func hex2dec(a byte) byte {
	if a >= '0' && a <= '9' {
		return a - '0'
	}
	if a >= 'A' && a <= 'F' {
		return a - 'A' + 10
	}
	return 0
}
