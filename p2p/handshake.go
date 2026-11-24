package p2p

type HandShakeFunc func(any) error

func NopeHandshakeFunc(any) error {return nil}
