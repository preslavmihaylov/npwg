package tftp

const (
	DatagramSize = 516
	BlockSize    = DatagramSize - 4 // minus header
)

type OpCode uint16

const (
	OpRRQ OpCode = iota + 1
	_            // No WRQ Support
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser

	ErrCodeCnt
)
