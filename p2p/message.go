package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC here, holds any arbitrary data thats sent over the transport bw 2 nodes in network

type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
