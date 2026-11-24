package p2p



type HandShakeFunc func(Peer) error

func NopeHandshakeFunc(Peer) error { return nil }
