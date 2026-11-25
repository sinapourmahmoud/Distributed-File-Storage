package main

import (
	"fmt"
	"log"

	"github.com/sinapourmahmoud/distributed-file-storage/p2p"
)

func onpeer(p p2p.Peer) error {
	p.Close()
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NopeHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        onpeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {

		for {
			msg := <-tr.Consume()

			fmt.Printf("%+v\n", msg)
		}

	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Distributed File Storage")

	select {}
}
