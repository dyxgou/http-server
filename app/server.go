package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

const defaultListenAddr = ":3000"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	Ln  net.Listener
	Ctx context.Context

	msgCh  chan []byte
	quitch chan<- struct{}
}

func CreateServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}

	return &Server{
		Config: cfg,
		Ctx:    context.Background(),
		msgCh:  make(chan []byte, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()
	s.Ln = ln
	slog.Info("The server has started on", "addr", s.ListenAddr)

	s.acceptConns()
	return nil
}

func (s *Server) acceptConns() {
	ctx, cancel := context.WithCancel(s.Ctx)
	defer cancel()

	for {
		conn, err := s.Ln.Accept()

		if err != nil {
			slog.Error("accepting conn err", "err", err)
			continue
		}

		slog.Info("conn acccepted", "addr", conn.RemoteAddr())

		go s.handleConn(ctx, conn)
	}
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn) {
	peer := CreatePeer(conn, s.msgCh)

	go peer.ReadConn(ctx)

	select {
	case <-ctx.Done():
		return
	case msg := <-s.msgCh:
		fmt.Println(string(msg))
	}
}
