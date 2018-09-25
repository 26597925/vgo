package main

import (
	"encoding/binary"
	"fmt"
	"vlock/vlp"
)

func testByteBuffer() {
	buffer := vlp.CreateByteBuffer(100)
	buffer.OrderSet(binary.LittleEndian)
	buffer.WriteByte(byte(1))
	buffer.WriteBytes([]byte{1,2})
	buffer.WriteShort(1)
	buffer.WriteInt(2)
	buffer.WriteLong(3)
	buffer.WriteString("1.1.0")
	buffer.WriteString("asdasdasdsad")

	fmt.Println(buffer.Order())
	fmt.Println(buffer.Slice())
	fmt.Println(len(buffer.Slice()))
	fmt.Printf("%p\n", buffer)

	getBuffer := vlp.WrapByteBuffer(buffer.Slice(), buffer.Length())
	getBuffer.OrderSet(binary.LittleEndian)

	fmt.Println(getBuffer.Order())
	fmt.Println(getBuffer.ReadBytes())
	fmt.Println(getBuffer.ReadByteSlice(2))
	fmt.Println(getBuffer.ReadShort())
	fmt.Println(getBuffer.ReadInt())
	fmt.Println(getBuffer.ReadLong())
	fmt.Println(getBuffer.ReadString())
	fmt.Println(getBuffer.ReadString())
	fmt.Printf("%p\n", getBuffer)
}

func testHeartbeat() {
	pack := vlp.CreatePacket(vlp.Heartbeat)
	buffer := vlp.CreateByteBuffer(100)
	pack.EncodePacket(buffer)
	fmt.Println(buffer.Slice())

	getbuffer := vlp.WrapByteBuffer(buffer.Slice(), buffer.Length())
	fmt.Println(getbuffer.ReadInt())
}

func testPacket() {
	buffer := vlp.CreateByteBuffer(100)
	pack := vlp.CreatePacket(vlp.Handshake)
	pack.AddFlags(vlp.FlagAutoAck)
	pack.SetSessionID(1)
	pack.SetBody([]byte("asdasd"))
	pack.SetCalcCheckCode()
	pack.SetCalcLrc()
	pack.EncodePacket(buffer)
	fmt.Println(buffer.Slice())

	getbuffer := vlp.WrapByteBuffer(buffer.Slice(), buffer.Length())
	getpack := vlp.CreatePacket(vlp.Unknown)
	getpack.DecodePacket(getbuffer)
	fmt.Println(getpack.Cmd())
	fmt.Println(getpack.SessionId())
	fmt.Println(getpack.Flags())
	fmt.Println(string(getpack.Body()))
	fmt.Println(getpack.CheckCode())
	fmt.Println(getpack.Lrc())
}

func main() {
	testByteBuffer()
	testHeartbeat()
	testPacket()
}