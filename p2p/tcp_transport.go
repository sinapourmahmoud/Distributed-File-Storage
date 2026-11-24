package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {

	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}

}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	handshakeFunc HandShakeFunc

	mu    sync.RWMutex
	peers map[net.Addr]string
}

func NewTCPTransport(listenAddress string, handshakeFunc HandShakeFunc) *TCPTransport {
	return &TCPTransport{
		handshakeFunc: handshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil

}

func (t *TCPTransport) startAcceptLoop() {

	for {

		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Printf("TCP access error: %s ", err)
		}

		go t.handleConn(conn)

	}

}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer := NewTCPPeer(conn, true)

	if err := t.handshakeFunc(conn); err != nil {

	}

	fmt.Printf("this is handle connection and this is the peer %v", peer)

}
