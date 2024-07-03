package p2p

// Handshake function

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
