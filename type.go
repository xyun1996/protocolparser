package main

import (
	"bytes"
	"encoding/binary"
	"net"
)

type IP struct {
	Version        int    //4bit
	IHL            int    //4bit
	DSCP           int    //7bit
	ECN            int    //1bit
	TotalLen       uint16 //16bit
	Identification uint16 //16bit
	Flags          int    //2bit
	FragmentOffset int    //14bit
	TTL            uint8  //8bit
	Protocal       uint8  //8bit
	HeaderChecksum uint16 //16bit
	SourceAddress  string //32bit
	DestAddress    string //32bit
	Options        []byte
}

func (m *IP) decode(b []byte) error {
	buf := bytes.NewBuffer(b)

	info := buf.Next(1)
	m.Version = int(info[0] >> 4)
	m.IHL = int(info[0] << 4 >> 4)
	info = buf.Next(1)
	m.DSCP = int(info[0] >> 1)
	m.ECN = int(info[0] << 7 >> 7)
	bs := buf.Next(2)
	m.TotalLen = binary.BigEndian.Uint16(bs)
	bs = buf.Next(2)
	m.Identification = binary.BigEndian.Uint16(bs)
	bs = buf.Next(2)
	tmp := binary.BigEndian.Uint16(bs)
	m.Flags = int(tmp >> 14)
	m.FragmentOffset = int(tmp << 2 >> 2)
	bs = buf.Next(1)
	m.TTL = uint8(bs[0])
	bs = buf.Next(1)
	m.Protocal = uint8(bs[0])
	bs = buf.Next(2)
	m.HeaderChecksum = binary.BigEndian.Uint16(bs)
	bs = buf.Next(4)
	m.SourceAddress = Ipv4(binary.BigEndian.Uint32(bs))
	bs = buf.Next(4)
	m.DestAddress = Ipv4(binary.BigEndian.Uint32(bs))
	if m.IHL > 5 {
		bs = buf.Next((m.IHL - 5) * 4)
		m.Options = bs
	}
	return nil
}

func Ipv4(nn uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip.String()
}
