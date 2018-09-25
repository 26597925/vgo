/*
*@author wenford.li
*@email  26597925@qq.com
*/

package vlp

/*
length	int	4	表示body的长度
cmd	byte	1	表示消息协议类型
checkcode	short	2	是根据body生成的一个校验码
flags	byte	1	表示当前包启用的特性，比如是否启用加密，是否启用压缩
sessionId	int	4	消息会话标识用于消息响应
lrc	byte	1	纵向冗余校验，用于校验header
*/
const (
	HeaderLen = 13
	HbPacketByte = -33

	FlagCrypto = 1
	FlagCompress = 2
	FlagBizAck = 4
	FlagAutoAck = 8
	FlagJSONBody = 16
)

type Packet struct {
	cmd 		byte
	checkCode	int16
	flags		byte
	sessionID	int
	lrc			byte
	body		[]byte
	bodyLength	int
}

func CreatePacket(cmd byte) *Packet {
	pack := &Packet{
		cmd: 		cmd,
		checkCode:	0,
		flags:		0,
		sessionID:	0,
		lrc:		0,
		bodyLength: 0,
	}
	return pack
}

func (pack *Packet) Cmd() byte{
	return pack.cmd
}

func (pack *Packet) SessionId() int {
	return pack.sessionID
}

func (pack *Packet) SetSessionID(sessionID int) {
	pack.sessionID = sessionID
}

func (pack *Packet) Flags() byte{
	return pack.flags
}

func (pack *Packet) AddFlags(flag byte) {
	pack.flags |= flag
}

func (pack *Packet) HasFlags(flag byte) bool {
	return (pack.flags & flag) != 0
}

func (pack *Packet) Body() []byte {
	return pack.body
}

func (pack *Packet) BodyLength() int {
	return pack.bodyLength
}

func (pack *Packet) SetBody(body []byte) {
	pack.bodyLength = len(body)
	pack.body = body
}

func (pack *Packet) CheckCode() int16{
	return pack.checkCode
}

func (pack *Packet) SetCalcCheckCode() {
	pack.checkCode = pack.calcCheckCode()
}

func (pack *Packet) Lrc() byte{
	return pack.lrc
}

func (pack *Packet) SetCalcLrc() {
	pack.lrc = pack.calcLrc()
}

func (pack *Packet) ValidCheckCode() bool{
	return pack.calcCheckCode() == pack.checkCode
}

func (pack *Packet) ValidLrc() bool {
	return (pack.lrc ^ pack.calcLrc()) == 0
}

func (pack *Packet) DecodePacket(buffer *ByteBuffer) {
	bodyLength := buffer.ReadInt()
	cmd := buffer.ReadBytes()

	pack.cmd = cmd
	pack.bodyLength = int(bodyLength)
	pack.checkCode = buffer.ReadShort()
	pack.flags = buffer.ReadBytes()
	pack.sessionID = int(buffer.ReadInt())
	pack.lrc = buffer.ReadBytes()

	if bodyLength > 0 {
		pack.body = buffer.ReadByteSlice(pack.bodyLength)
	}
}

func (pack *Packet) EncodePacket(buffer *ByteBuffer) {
	if pack.cmd == Heartbeat {
		buffer.WriteInt(HbPacketByte)
	} else {
		buffer.WriteInt(int32(pack.bodyLength))
		buffer.WriteByte(pack.cmd)
		buffer.WriteShort(pack.checkCode)
		buffer.WriteByte(pack.flags)
		buffer.WriteInt(int32(pack.sessionID))
		buffer.WriteByte(pack.lrc)
		if pack.bodyLength > 0 {
			buffer.WriteBytes(pack.body)
		}
	}
}

func (pack *Packet) calcLrc() byte{
	header := CreateByteBuffer(HeaderLen - 1)
	header.WriteInt(int32(pack.bodyLength))
	header.WriteByte(pack.cmd)
	header.WriteShort(pack.checkCode)
	header.WriteByte(pack.flags)
	header.WriteInt(int32(pack.sessionID))

	var lrc byte
	for i := 0; i < header.length; i++ {
		lrc ^= header.buffer[i];
	}
	return lrc
}

func (pack *Packet) calcCheckCode() int16{
	var checkCode int16
	for  i:= 0; i < pack.bodyLength; i++ {
		checkCode += (int16(pack.body[i]) & 0x0ff);
	}
	return checkCode
}