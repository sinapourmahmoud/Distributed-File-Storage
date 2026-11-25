package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

	tcpOpts := TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: NopeHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}

	listenAddr := ":3000"

	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tr.ListenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

}
