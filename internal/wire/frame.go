package wire

import (
	"bytes"

	"github.com/gpfnw/quic-go/internal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Write(b *bytes.Buffer, version protocol.VersionNumber) error
	MinLength(version protocol.VersionNumber) protocol.ByteCount
}
