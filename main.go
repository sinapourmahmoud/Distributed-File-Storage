package main

import (
	"fmt"
	"log"

	"github.com/sinapourmahmoud/distributed-file-storage/p2p"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NopeHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Distributed File Storage")

	select {}
}
