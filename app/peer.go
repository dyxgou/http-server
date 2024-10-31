package main

import (
	"context"
	"log/slog"
	"net"
)

type Peer struct {
	Conn   net.Conn
	MsgCh  chan []byte
	Ctx    context.Context
	Cancel context.CancelFunc
}

func CreatePeer(ctx context.Context, conn net.Conn, msgch chan []byte) *Peer {
	ctx, cancel := context.WithCancel(ctx)

	return &Peer{
		Conn:   conn,
		MsgCh:  msgch,
		Ctx:    ctx,
		Cancel: cancel,
	}
}

func (p *Peer) ReadConn() {
	defer p.Cancel()
	buf := make([]byte, 1024)

	for {
		n, err := p.Conn.Read(buf)

		if err != nil {
			slog.Error("reading conn err", "err", err)
			return
		}

		slog.Info("message from", "addr", p.Conn.RemoteAddr())
		msgBuf := make([]byte, n)

		copy(msgBuf, buf[:n])
		p.MsgCh <- msgBuf
	}
}

func (p *Peer) WriteConn(msg *[]byte) {
	defer p.Cancel()

	const headers string = "HTTP/1.1 200 OK\r\n\r\n"
	headersLen := len(headers)
	ans := make([]byte, 0, headersLen+len(*msg))

	ans = append(ans, []byte(headers)...)
	ans = append(ans, *msg...)

	_, err := p.Conn.Write(ans)

	if err != nil {
		slog.Error("writing conn", "err", err)
		return
	}
}
