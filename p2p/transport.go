package p2p

import "net"

// Peer is an interface for representation of remote node

type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Now, lets make Transport -> handling connection bw nodes in the network. Can be TCP, UDP, WebSockets, yada yada

type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
