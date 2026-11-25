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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddress string

	Decoder Decoder

	HandshakeFunc HandShakeFunc
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	rpcch chan RPC

	mu    sync.RWMutex
	peers map[net.Addr]string
}

func (t *TCPTransport) Consume() <-chan RPC {

	return t.rpcch

}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
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

	rpc := RPC{}

	for {

		if err := t.Decoder.Decode(conn, &rpc); err != nil {

			fmt.Printf("TCP error: %v", err)
			continue
		}

		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc

		fmt.Printf("Message : %+v\n", rpc)
	}

}
