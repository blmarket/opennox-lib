package noxnet

import (
	"encoding/binary"
	"io"
)

func init() {
	RegisterMessage(&MsgTimestamp{})
	RegisterMessage(&MsgFullTimestamp{})
}

type MsgTimestamp struct {
	T uint16
}

func (*MsgTimestamp) NetOp() Op {
	return MSG_TIMESTAMP
}

func (*MsgTimestamp) EncodeSize() int {
	return 2
}

func (m *MsgTimestamp) Encode(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, io.ErrShortBuffer
	}
	binary.LittleEndian.PutUint16(data[:2], m.T)
	return 2, nil
}

func (m *MsgTimestamp) Decode(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	m.T = binary.LittleEndian.Uint16(data[:2])
	return 2, nil
}

type MsgFullTimestamp struct {
	T uint32
}

func (*MsgFullTimestamp) NetOp() Op {
	return MSG_FULL_TIMESTAMP
}

func (*MsgFullTimestamp) EncodeSize() int {
	return 4
}

func (m *MsgFullTimestamp) Encode(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, io.ErrShortBuffer
	}
	binary.LittleEndian.PutUint32(data[:4], m.T)
	return 4, nil
}

func (m *MsgFullTimestamp) Decode(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	m.T = binary.LittleEndian.Uint32(data[:4])
	return 4, nil
}
