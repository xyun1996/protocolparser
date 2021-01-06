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
	fmt.Printf("bytes len %d\n", len(b))
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
	b := make([]byte, len(hex)/2)
	hex = strings.ToUpper(hex)
	buff := bytes.NewBufferString(hex)
	for i := 0; i < len(hex); i += 2 {
		info := buff.Next(2)
		b[i/2] = (info[1] - 48) | (info[0]-48)<<4
	}
	fmt.Println(b)
	return b, nil
}
