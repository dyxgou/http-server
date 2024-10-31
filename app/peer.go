package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn  net.Conn
	reqch chan *Request

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPeer(parentCtx context.Context, conn net.Conn) *Peer {
	ctx, cancel := context.WithCancel(parentCtx)

	return &Peer{
		conn:  conn,
		reqch: make(chan *Request, 10),

		ctx:    ctx,
		cancel: cancel,
	}
}

func (p *Peer) ReadConn() {
	defer p.cancel()
	defer p.conn.Close()

	buf := make([]byte, 2048)

	for {
		select {
		case <-p.ctx.Done():
			return
		default:
			n, err := p.conn.Read(buf)

			if err != nil {
				slog.Error("reading conn err", "err", err)
				return
			}

			m, err := GetMethod(buf[0])

			if err != nil {
				slog.Error("reading err err", "err", err)
				return
			}

			msgBuf := make([]byte, n)
			copy(msgBuf, buf[m.GetLen():])
			req := NewRequest(m, msgBuf)

			p.reqch <- req
		}
	}
}

func (p *Peer) WriteConn(req *Request) {
	defer p.cancel()

	select {
	case <-p.ctx.Done():
		return
	default:

	}
}

func (p *Peer) HandleMsg() {
	defer p.cancel()

	for {
		select {
		case <-p.ctx.Done():
			return
		case msg := <-p.reqch:
			fmt.Println(msg)
		}
	}
}
