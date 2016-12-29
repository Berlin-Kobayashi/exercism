// Package circular implements a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
package circular

import (
	"errors"
)

const testVersion = 4

// Buffer represents a circular buffer.
type Buffer struct {
	content       []byte
	writePosition int
	readPosition  int
	length        int
}

// NewBuffer creates a new Buffer with the given size.
func NewBuffer(size int) *Buffer {
	return &Buffer{content: make([]byte, size, size)}
}

// WriteByte writes a new byte into the buffer. If the buffer is full an error is returned.
func (buffer *Buffer) WriteByte(input byte) error {
	if buffer.length == len(buffer.content) {
		return errors.New("Cannot write to full buffer without force")
	}

	buffer.content[buffer.writePosition] = input
	buffer.inrementWritePosition()
	buffer.length++

	return nil
}

// ReadByte reads the oldest byte in the buffer. If the buffer is empty an error is returned.
func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.length == 0 {
		return 0, errors.New("Cannot read from empty buffer")
	}

	result := buffer.content[buffer.readPosition]
	buffer.incrementReadPosition()
	buffer.length--

	return result, nil
}

// Overwrite writes a new byte into the buffer. If the buffer is full the oldest value in the buffer will be overwritten.
func (buffer *Buffer) Overwrite(input byte) {
	if buffer.length != len(buffer.content) {
		buffer.WriteByte(input)
	} else {
		buffer.content[buffer.writePosition] = input
		buffer.inrementWritePosition()
		buffer.incrementReadPosition()
	}
}

func (buffer *Buffer) inrementWritePosition() {
	buffer.writePosition = buffer.nextCirclePosition(buffer.writePosition)
}

func (buffer *Buffer) incrementReadPosition() {
	buffer.readPosition = buffer.nextCirclePosition(buffer.readPosition)
}

func (buffer *Buffer) nextCirclePosition(position int) int {
	return (position + 1) % len(buffer.content)
}

// Reset empites the buffer.
func (buffer *Buffer) Reset() {
	buffer.content = make([]byte, len(buffer.content), len(buffer.content))
	buffer.length = 0
	buffer.writePosition = 0
	buffer.readPosition = 0
}
