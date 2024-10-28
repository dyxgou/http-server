package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	Conn  net.Conn
	MsgCh chan []byte
}

func CreatePeer(conn net.Conn, msgch chan []byte) *Peer {
	return &Peer{
		Conn:  conn,
		MsgCh: msgch,
	}
}

func (p *Peer) ReadConn(ctx context.Context) {
	buf := make([]byte, 1024)
	defer p.Conn.Close()
	defer ctx.Done()

	for {
		n, err := p.Conn.Read(buf)

		if err != nil {
			slog.Error("reading conn err", "err", err)
			return
		}

		slog.Info("message from", "addr", p.Conn.RemoteAddr())
		msgBuf := make([]byte, n)

		copy(msgBuf, buf[:n])
		fmt.Print("msg ", string(msgBuf), msgBuf)
		p.MsgCh <- msgBuf
	}
}
