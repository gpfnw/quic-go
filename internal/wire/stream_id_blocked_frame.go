package wire

import (
	"bytes"

	"github.com/gpfnw/quic-go/internal/protocol"
	"github.com/gpfnw/quic-go/internal/utils"
)

// A StreamIDBlockedFrame is a STREAM_ID_BLOCKED frame
type StreamIDBlockedFrame struct {
	StreamID protocol.StreamID
}

// ParseStreamIDBlockedFrame parses a STREAM_ID_BLOCKED frame
func ParseStreamIDBlockedFrame(r *bytes.Reader, _ protocol.VersionNumber) (*StreamIDBlockedFrame, error) {
	if _, err := r.ReadByte(); err != nil {
		return nil, err
	}
	streamID, err := utils.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	return &StreamIDBlockedFrame{StreamID: protocol.StreamID(streamID)}, nil
}

func (f *StreamIDBlockedFrame) Write(b *bytes.Buffer, _ protocol.VersionNumber) error {
	typeByte := uint8(0x0a)
	b.WriteByte(typeByte)
	utils.WriteVarInt(b, uint64(f.StreamID))
	return nil
}

// MinLength of a written frame
func (f *StreamIDBlockedFrame) MinLength(_ protocol.VersionNumber) protocol.ByteCount {
	return 1 + utils.VarIntLen(uint64(f.StreamID))
}
