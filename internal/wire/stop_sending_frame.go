package wire

import (
	"bytes"

	"github.com/gpfnw/quic-go/internal/protocol"
	"github.com/gpfnw/quic-go/internal/utils"
)

// A StopSendingFrame is a STOP_SENDING frame
type StopSendingFrame struct {
	StreamID  protocol.StreamID
	ErrorCode protocol.ApplicationErrorCode
}

// ParseStopSendingFrame parses a STOP_SENDING frame
func ParseStopSendingFrame(r *bytes.Reader, _ protocol.VersionNumber) (*StopSendingFrame, error) {
	if _, err := r.ReadByte(); err != nil { // read the TypeByte
		return nil, err
	}

	streamID, err := utils.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	errorCode, err := utils.BigEndian.ReadUint16(r)
	if err != nil {
		return nil, err
	}

	return &StopSendingFrame{
		StreamID:  protocol.StreamID(streamID),
		ErrorCode: protocol.ApplicationErrorCode(errorCode),
	}, nil
}

// MinLength of a written frame
func (f *StopSendingFrame) MinLength(_ protocol.VersionNumber) protocol.ByteCount {
	return 1 + utils.VarIntLen(uint64(f.StreamID)) + 2
}

func (f *StopSendingFrame) Write(b *bytes.Buffer, _ protocol.VersionNumber) error {
	b.WriteByte(0x0c)
	utils.WriteVarInt(b, uint64(f.StreamID))
	utils.BigEndian.WriteUint16(b, uint16(f.ErrorCode))
	return nil
}
