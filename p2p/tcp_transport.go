package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	// income connection => true
	// outcome connection => false
	incoming bool
}

func NewTCPPeer(conn net.Conn, incoming bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		incoming: incoming,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
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
			fmt.Printf("TCP accept error : %s\n", err)
		}

		t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {

	peer := NewTCPPeer(conn, true)
	fmt.Printf("New Connection : %+v\n", peer)
}
