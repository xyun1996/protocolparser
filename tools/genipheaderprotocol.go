package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func genIpHeaderProtocol(filename string) error {
	fs, err := os.Open(filename)
	if err != nil {
		return err
	}
	br := bufio.NewReader(fs)
	protocol := map[int]string{}
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(a)
		params := strings.Split(line, "\t\t")
		if len(params) < 2 {
			continue
		}
		id, err := strconv.Atoi(params[0])
		if err != nil {
			continue
		}
		protocol[id] = params[1]
	}
	tpl, err := template.ParseFiles("ipheaderprotocol.tpl")
	if err != nil {
		return err
	}
	fs, err = os.OpenFile("../const/idheaderprotocol.go", os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != err {
		return err
	}
	return tpl.Execute(fs, protocol)
}

func main() {
	genIpHeaderProtocol("ipheaderprotocol.txt")
}
