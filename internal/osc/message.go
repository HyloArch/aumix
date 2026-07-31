package osc

import (
	"bytes"
	"encoding/binary"
	"math"
)

type Message struct {
	Address    string
	Parameters []Parameter
}

func (m Message) Encode(buffer []byte) int {
	copy(buffer, []byte(m.Address))

	outIndex := len(m.Address)
	outIndex = outIndex & ^0b11 + 4

	buffer[outIndex] = byte(',')
	outIndex++

	paramLen := len(m.Parameters)
	paramIndex := outIndex

	outIndex = (outIndex+paramLen) & ^0b11 + 4

	for i := range paramLen {
		switch p := m.Parameters[i].(type) {
		case IntParam:
			buffer[paramIndex] = byte('i')
			binary.BigEndian.PutUint32(buffer[outIndex:], uint32(p))
			outIndex += 4
		case FloatParam:
			buffer[paramIndex] = byte('f')
			bits := math.Float32bits(float32(p))
			binary.BigEndian.PutUint32(buffer[outIndex:], bits)
			outIndex += 4
		case StringParam:
			buffer[paramIndex] = byte('s')
			length := len(p)
			copy(buffer[outIndex:], []byte(p))
			outIndex = (outIndex+length) & ^0b11 + 4
		case ByteBlobParam:
			buffer[paramIndex] = byte('b')
			length := len(p)
			binary.BigEndian.PutUint32(buffer[outIndex:], uint32(length))
			outIndex += 4
			copy(buffer[outIndex:], p)
			outIndex = (outIndex+length-1) & ^0b11 + 4
		}
		paramIndex++
	}

	return outIndex
}

func Decode(input []byte) Message {
	var message Message

	inputIndex := bytes.IndexByte(input, 0x00)
	message.Address = string(input[:inputIndex])

	inputIndex = inputIndex & ^0b11 + 4

	if input[inputIndex] != byte(',') {
		return message
	}
	inputIndex++

	paramIndex := inputIndex
	paramLen := bytes.IndexByte(input[inputIndex:], 0x00)
	message.Parameters = make([]Parameter, paramLen)

	inputIndex = (inputIndex+paramLen) & ^0b11 + 4

	for i := range paramLen {
		switch input[paramIndex] {
		case byte('i'):
			message.Parameters[i] = IntParam(binary.BigEndian.Uint32(input[inputIndex:]))
			inputIndex += 4
		case byte('f'):
			bits := binary.BigEndian.Uint32(input[inputIndex:])
			message.Parameters[i] = FloatParam(math.Float32frombits(bits))
			inputIndex += 4
		case byte('s'):
			end := bytes.IndexByte(input[inputIndex:], 0x00)
			message.Parameters[i] = StringParam(input[inputIndex : inputIndex+end])
			inputIndex = (inputIndex+end) & ^0b11 + 4
		case byte('b'):
			length := int(binary.BigEndian.Uint32(input[inputIndex:]))
			inputIndex += 4
			blob := make(ByteBlobParam, length)
			copy(blob, input[inputIndex:inputIndex+length])
			message.Parameters[i] = blob
			inputIndex = (inputIndex+length-1) & ^0b11 + 4
		}
		paramIndex++
	}

	return message
}
