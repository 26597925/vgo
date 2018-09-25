/*
*@author wenford.li
*@email  26597925@qq.com
*/

package vlp

import (
	"encoding/binary"
)

const (
	MaxValue = 32767
)

type ByteBuffer struct {
	capacity	int
	length		int
	limit 		int
	position 	int
	buffer		[]byte
	order		binary.ByteOrder
}

func CreateByteBuffer(max int) *ByteBuffer{
	buf := &ByteBuffer{
		capacity:	max,
		length:		0,
		limit:		max,
		position:	0,
		order:		binary.BigEndian,
	}
	return buf
}

func WrapByteBuffer(buffer []byte, length int) *ByteBuffer{
	buf := &ByteBuffer{
		capacity:	length,
		length:		length,
		limit:		length,
		position:	0,
		buffer: 	buffer,
		order:		binary.BigEndian,
	}
	return buf
}

func (buf *ByteBuffer) Slice() []byte {
	return buf.buffer
}

func (buf *ByteBuffer) Length() int {
	return buf.length
}

func (buf *ByteBuffer) Capacity() int {
	return buf.capacity
}

func (buf *ByteBuffer) Order() string {
	return buf.order.String()
}

func (buf *ByteBuffer) OrderSet(order binary.ByteOrder) {
	buf.order = order
}

func (buf *ByteBuffer) ReadBytes() byte{
	pos := buf.position
	buf.position++
	return buf.buffer[pos]
}

func (buf *ByteBuffer) ReadShort() int16 {
	pos := buf.position
	if (pos + 2) > buf.limit {
		return 0
	}

	buffer := buf.buffer[pos:pos + 2]

	buf.position += 2
	return int16(buf.order.Uint16(buffer))
}

func (buf *ByteBuffer) ReadInt() int32 {
	pos := buf.position
	if (pos + 4) > buf.limit {
		return 0
	}

	buffer := buf.buffer[pos:pos + 4]
	buf.position += 4
	return int32(buf.order.Uint32(buffer))
}

func (buf *ByteBuffer) ReadLong() int64 {
	pos := buf.position
	if (pos + 8) > buf.limit {
		return 0
	}

	buffer := buf.buffer[pos:pos + 8]
	buf.position += 8
	return int64(buf.order.Uint64(buffer))
}

func (buf *ByteBuffer) ReadByteSlice(length int) []byte {
	pos := buf.position
	if (buf.limit - pos) < length {
		return nil
	}

	buffer := buf.buffer[pos:pos + length]
	buf.position += length
	return buffer
}

func (buf *ByteBuffer) ReadString() string {
	var length int32

	length = int32(buf.ReadShort())

	if length == 0{
		return ""
	}

	if length == MaxValue {
		length += buf.ReadInt()
	}

	buffer := buf.ReadByteSlice(int(length))
	return string(buffer)
}

func (buf *ByteBuffer) ReadableBytes() int {
	return buf.length - buf.position
}

func (buf *ByteBuffer) WriteByte(value byte) {
	if buf.position >= buf.limit {
		return
	}

	buf.buffer = append(buf.buffer, value)
	buf.position++
	buf.length = buf.position
}

func (buf *ByteBuffer) WriteShort(value int16) {
	pos := buf.position
	if(pos + 2) > buf.limit {
		return
	}
	buffer := make([]byte, 2)
	buf.order.PutUint16(buffer, uint16(value))
	buf.appendBuffer(buffer)

	buf.position += 2
	buf.length = buf.position
}

func (buf *ByteBuffer) WriteInt(value int32) {
	pos := buf.position
	if(pos + 4) > buf.limit {
		return
	}

	buffer := make([]byte, 4)
	buf.order.PutUint32(buffer, uint32(value))
	buf.appendBuffer(buffer)

	buf.position += 4
	buf.length = buf.position
}

func (buf *ByteBuffer) WriteLong(value int64) {
	pos := buf.position
	if(pos + 8) > buf.limit {
		return
	}

	buffer := make([]byte, 8)
	buf.order.PutUint64(buffer, uint64(value))
	buf.appendBuffer(buffer)

	buf.position += 8
	buf.length = buf.position
}

func (buf *ByteBuffer) WriteBytes(value []byte) {
	pos := buf.position
	length := len(value)
	if (pos + length) >= buf.limit {
		return
	}

	buf.appendBuffer(value)
	
	buf.position += length
	buf.length = buf.position
}

func (buf *ByteBuffer) WriteString(value string) {
	buffer := []byte(value)
	length := len(buffer)

	if (buf.position + length) >= buf.limit {
		return
	}

	if length == 0 {
		buf.WriteShort(0)
	}else if length < MaxValue {
		buf.WriteShort(int16(length))
		buf.WriteBytes(buffer)
	}else {
		buf.WriteShort(MaxValue)
		buf.WriteInt(int32(length - MaxValue))
		buf.WriteBytes(buffer)
	}
}

func (buf *ByteBuffer) appendBuffer(buffer []byte) {
	for _, x := range buffer{
		buf.buffer = append(buf.buffer, x)
	}
}