package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn   net.Conn
	reqch  chan *Request
	parser Parser

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPeer(parentCtx context.Context, conn net.Conn) *Peer {
	ctx, cancel := context.WithCancel(parentCtx)

	return &Peer{
		conn:   conn,
		reqch:  make(chan *Request, 10),
		parser: NewRequestParser(),

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

			req, err := p.parser.Parse(buf, n)

			if err != nil {
				slog.Error("parsing request err", "err", err)
				return
			}

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
