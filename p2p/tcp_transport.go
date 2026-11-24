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

type TCPTransportOpts struct {
	ListenAddress string

	Decoder Decoder

	HandshakeFunc HandShakeFunc
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]string
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)

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

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()

		fmt.Printf("TCP handshake error: %v", err)
		return

	}

	msg := &Message{}

	for {

		if err := t.Decoder.Decode(conn, msg); err != nil {

			fmt.Printf("TCP error: %v", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("Message : %+v\n",msg)
	}

}
