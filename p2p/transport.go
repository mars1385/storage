package p2p

// represent remote node
type Peer interface{}

// represent commination between nodes . can be (tcp,upd,websocket,......)
type Transport interface {
	ListenAndAccept() error
}
