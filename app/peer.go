package main

import (
	"context"
	"net"
)

type Peer struct {
	conn net.Conn
}

func (p *Peer) ReadConn(ctx context.Context) {
	buf := make([]byte, 1024)
	defer p.conn.Close()

	for {
		n, err := p.conn.Read()

		if err != nil {

		}
	}
}
